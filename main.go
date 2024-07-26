package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	background      = lipgloss.Color("#11111b")
	backgroundAlt   = lipgloss.Color("#181825")
	backgroundTrans = lipgloss.Color("#aa181825")
	foreground      = lipgloss.Color("#cdd6f4")
	primary         = lipgloss.Color("#a6e3a1")
	secondary       = lipgloss.Color("#eba0ac")
	alert           = lipgloss.Color("#f38ba8")
	active          = lipgloss.Color("#cdd6f4")
	disabled        = lipgloss.Color("#a6adc8")
	inactive        = lipgloss.Color("#a6adc8")
	faded           = lipgloss.Color("#a6adc8")
)

type page struct {
	title string
	key   string
}

type model struct {
	ready      bool
	activePage int
	nav        *nav
	content    *content
	footer     footer
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

		case key.Matches(msg, inputKeys.prev):
			m.nav.model.CursorUp()

		}

	case tea.WindowSizeMsg:
		// layout on window size
		if !m.ready {
			termWidth := msg.Width
			termHeight := msg.Height

			footerHeight := m.footer.style.GetHeight()
			navWidth := m.nav.style.GetWidth()

			m.nav.model.SetHeight(termHeight - footerHeight - 2)
			m.nav.model.SetWidth(25)
			m.content.model.Width = termWidth - navWidth
			m.content.model.Height = termHeight - footerHeight - 2

			m.ready = true
		}

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
	m := model{
		activePage: 0,
		content:    newContent([]string{"pages/home.md"}),
		nav: newNav([]list.Item{
			item{title: "🏡 Home", desc: "The home page"},
			item{title: "✨ About", desc: "The about section"},
			item{title: "🏗️ Projects", desc: "Open source stuff"},
			item{title: "📬️ Contact", desc: "Socials and stuff"},
		}),
		footer: newFooter(inputKeys),
	}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to start: %s\n", err)
		os.Exit(1)
	}
}
