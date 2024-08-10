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
		Footer:  sections.NewFooter(renderer, keys.InputKeys),
	}
}

type model struct {
	Term     string
	Nav      *sections.Nav
	Content  *sections.Content
	Footer   *sections.Footer
	Renderer *lipgloss.Renderer

	ready      bool
	activePage int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
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
			_, cmd = m.Nav.Update(sections.SelectIndex(m.activePage))
			cmds = append(cmds, cmd)
			_, cmd = m.Content.Update(pages.ActivePage(m.activePage))
			cmds = append(cmds, cmd)

		case key.Matches(msg, keys.InputKeys.Prev):
			if m.activePage == 0 {
				m.activePage = m.Nav.Length - 1
			} else {
				m.activePage--
			}
			_, cmd = m.Nav.Update(sections.SelectIndex(m.activePage))
			cmds = append(cmds, cmd)
			_, cmd = m.Content.Update(pages.ActivePage(m.activePage))
			cmds = append(cmds, cmd)

		}

		_, cmd = m.Content.Update(msg)
		cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		m.ready = false

		viewportHeight := msg.Height - m.Footer.Height() - 2

		_, cmd = m.Nav.Update(tea.WindowSizeMsg{
			Width:  23,
			Height: viewportHeight,
		})
		cmds = append(cmds, cmd)

		_, cmd = m.Content.Update(tea.WindowSizeMsg{
			Width:  msg.Width - m.Nav.Width(),
			Height: viewportHeight,
		})
		cmds = append(cmds, cmd)

		// explicitly call update so that wordwrap is applied
		_, cmd = m.Content.Update(pages.ActivePage(m.activePage))
		cmds = append(cmds, cmd)

		m.ready = true
	}

	return m, tea.Batch(cmds...)
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
