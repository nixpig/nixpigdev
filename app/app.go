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
	var pages = []pages.Page{
		&pages.Uses,
		&pages.Contact,
	}

	return model{
		Content: sections.NewContent(renderer, pages),
		Nav:     sections.NewNav(renderer, pages),
		Footer:  sections.NewFooter(renderer, keys.GlobalKeys),
		pages:   pages,
	}
}

type model struct {
	Term     string
	Nav      *sections.Nav
	Content  *sections.Content
	Footer   *sections.Footer
	Renderer *lipgloss.Renderer
	pages    []pages.Page

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
		case key.Matches(msg, keys.GlobalKeys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keys.GlobalKeys.Next):
			if m.activePage == m.Nav.Length-1 {
				m.activePage = 0
			} else {
				m.activePage++
			}

			_, navCmd := m.Nav.Update(sections.SelectIndex(m.activePage))
			_, contentCmd := m.Content.Update(pages.ActivePage(m.activePage))
			return m, tea.Batch(navCmd, contentCmd)

		case key.Matches(msg, keys.GlobalKeys.Prev):
			if m.activePage == 0 {
				m.activePage = m.Nav.Length - 1
			} else {
				m.activePage--
			}

			_, navCmd := m.Nav.Update(sections.SelectIndex(m.activePage))
			_, contentCmd := m.Content.Update(pages.ActivePage(m.activePage))
			return m, tea.Batch(navCmd, contentCmd)
		}

		m.Content.Update(msg)

	case tea.WindowSizeMsg:
		m.ready = false

		viewportHeight := msg.Height - m.Footer.Height() - 2

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
