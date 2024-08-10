package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/nixpig/nixpigdev/app/pages"
)

func New(pty ssh.Pty, renderer *lipgloss.Renderer) model {
	var pages = []pages.Page{
		pages.Home(renderer),
		pages.Scrapbook(renderer),
		pages.Projects(renderer),
		pages.Uses(renderer),
		pages.Resume(renderer),
		pages.Contact(renderer),
	}

	return model{
		Term:    pty.Term,
		Width:   pty.Window.Width,
		Height:  pty.Window.Height,
		Content: newContent(renderer, pages),
		Nav:     newNav(renderer, pages),
		Footer:  newFooter(renderer, InputKeys),
	}
}

type model struct {
	Term     string
	Width    int
	Height   int
	Nav      *nav
	Content  *content
	Footer   *footer
	Renderer *lipgloss.Renderer

	ready      bool
	activePage int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, InputKeys.quit):
			return m, tea.Quit

		case key.Matches(msg, InputKeys.down):
			m.Content.model.LineDown(1)

		case key.Matches(msg, InputKeys.up):
			m.Content.model.LineUp(1)

		case key.Matches(msg, InputKeys.next):
			m.Nav.model.CursorDown()
			if m.Nav.model.Index() != m.activePage {
				m.Content.update(m.Nav.model.Index())
				m.activePage = m.Nav.model.Index()
			}

		case key.Matches(msg, InputKeys.prev):
			m.Nav.model.CursorUp()
			if m.Nav.model.Index() != m.activePage {
				m.Content.update(m.Nav.model.Index())
				m.activePage = m.Nav.model.Index()
			}
		}

	case tea.WindowSizeMsg:
		m.ready = false

		viewportHeight := msg.Height - m.Footer.style.GetHeight() - 2

		m.Nav.model.SetWidth(23)
		m.Nav.model.SetHeight(viewportHeight)

		m.Content.model.Width = msg.Width - m.Nav.model.Width()
		m.Content.model.Height = viewportHeight

		// explicitly call update so that wordwrap is applied
		m.Content.update(m.activePage)

		m.ready = true
	}

	return m, nil
}

func (m model) View() string {
	if !m.ready {
		return "\nLoading..."
	}

	v := strings.Builder{}

	v.WriteString(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				m.Nav.view(),
				m.Content.view(),
			),
			m.Footer.view(),
		),
	)

	return v.String()
}
