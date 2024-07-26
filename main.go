package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	ready      bool
	activePage int
	nav        *nav
	content    *content
	footer     *footer
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, inputKeys.quit):
			return m, tea.Quit

		case key.Matches(msg, inputKeys.down):
			m.content.model.LineDown(1)

		case key.Matches(msg, inputKeys.up):
			m.content.model.LineUp(1)

		case key.Matches(msg, inputKeys.next):
			m.nav.model.CursorDown()
			if m.nav.model.Index() != m.activePage {
				m.content.update(m.nav.model.Index())
				m.activePage = m.nav.model.Index()
			}

		case key.Matches(msg, inputKeys.prev):
			m.nav.model.CursorUp()
			if m.nav.model.Index() != m.activePage {
				m.content.update(m.nav.model.Index())
				m.activePage = m.nav.model.Index()
			}
		}

	case tea.WindowSizeMsg:
		// layout on window size
		m.ready = false

		viewportHeight := msg.Height - m.footer.style.GetHeight() - 2

		m.nav.model.SetHeight(viewportHeight)
		m.nav.model.SetWidth(25)

		m.content.model.Width = msg.Width - m.nav.model.Width()
		m.content.model.Height = viewportHeight

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
				m.nav.view(),
				m.content.view(),
			),
			m.footer.view(),
		),
	)

	return v.String()
}

func main() {
	pages := []page{
		{
			title:    "🏡 Home",
			desc:     "The home page",
			filepath: "pages/home.md",
		},
		{
			title:    "✨ About",
			desc:     "The about section",
			filepath: "pages/about.md",
		},
		{
			title:    "🏗️ Projects",
			desc:     "Open source stuff",
			filepath: "pages/projects.md",
		},
		{
			title:    "📬️ Contact",
			desc:     "Socials and stuff",
			filepath: "pages/contact.md",
		},
	}

	m := model{
		activePage: 0,
		content:    newContent(pages),
		nav:        newNav(pages),
		footer:     newFooter(inputKeys),
	}

	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to start: %s\n", err)
		os.Exit(1)
	}
}
