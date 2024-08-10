package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/pages"
	"github.com/nixpig/nixpigdev/app/sections"
)

func New(pty ssh.Pty, renderer *lipgloss.Renderer) model {
	ch := make(chan string)
	var pages = []pages.Page{
		pages.Home(renderer),
		pages.Scrapbook(renderer),
		pages.Projects(renderer),
		pages.Uses(renderer),
		pages.Resume(renderer),
		pages.Contact(renderer, ch),
	}

	return model{
		Term:    pty.Term,
		Width:   pty.Window.Width,
		Height:  pty.Window.Height,
		Content: sections.NewContent(renderer, pages),
		Nav:     sections.NewNav(renderer, pages),
		Footer:  sections.NewFooter(renderer, keys.InputKeys),
		Ch:      ch,
	}
}

type model struct {
	Term     string
	Width    int
	Height   int
	Nav      *sections.Nav
	Content  *sections.Content
	Footer   *sections.Footer
	Renderer *lipgloss.Renderer
	Ch       chan string

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
		case key.Matches(msg, keys.InputKeys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keys.InputKeys.Next):
			if m.activePage == m.Nav.Length-1 {
				m.activePage = 0
			} else {
				m.activePage++
			}
			m.Nav.Update(sections.SelectIndex(m.activePage))
			m.Content.Update(pages.ActivePage(m.activePage))

		case key.Matches(msg, keys.InputKeys.Prev):
			if m.activePage == 0 {
				m.activePage = m.Nav.Length - 1
			} else {
				m.activePage--
			}
			m.Nav.Update(sections.SelectIndex(m.activePage))
			m.Content.Update(pages.ActivePage(m.activePage))
		}

		m.Content.Update(msg)

	case tea.WindowSizeMsg:
		m.ready = false

		viewportHeight := msg.Height - m.Footer.GetHeight() - 2

		m.Nav.Update(tea.WindowSizeMsg{
			Width:  23,
			Height: viewportHeight,
		})

		m.Content.Update(tea.WindowSizeMsg{
			Width:  msg.Width - m.Nav.Width(),
			Height: viewportHeight,
		})

		// explicitly call update so that wordwrap is applied
		m.Content.Update(pages.ActivePage(m.activePage))

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
				m.Nav.View(),
				m.Content.View(),
			),
			m.Footer.View(),
		),
	)

	return v.String()
}
