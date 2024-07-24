package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
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

	inactiveTabStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 2, 0, 1).Foreground(inactive).BorderForeground(inactive)
	activeTabStyle   = inactiveTabStyle.Bold(true).Foreground(active)
	contentStyle     = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.NormalBorder()).BorderForeground(inactive)
	quitStyle        = lipgloss.NewStyle().Foreground(faded)

	containerStyle = lipgloss.NewStyle().Align(lipgloss.Center).Padding(2)
)

type page struct {
	title   string
	key     string
	content string
}

type model struct {
	pages      []page
	activePage int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		default:
			selection := slices.IndexFunc(m.pages, func(p page) bool {
				return p.key == msg.String()
			})

			if selection != -1 {
				m.activePage = selection
			}

			return m, nil
		}

	}

	return m, nil
}

func (m model) View() string {
	termWidth, _, err := term.GetSize(os.Stdin.Fd())
	if err != nil {
		return ""
	}
	v := strings.Builder{}

	var renderedTabs []string

	for i, page := range m.pages {
		if i == m.activePage {
			renderedTabs = append(renderedTabs, activeTabStyle.Render(fmt.Sprintf("(%s) %s", page.key, page.title)))
		} else {
			renderedTabs = append(renderedTabs, inactiveTabStyle.Render(fmt.Sprintf("(%s) %s", page.key, page.title)))
		}
	}

	joinedTabs := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	v.WriteString(joinedTabs)

	v.WriteString("\n")
	v.WriteString(contentStyle.Width(lipgloss.Width(joinedTabs) - 2).Render(m.pages[m.activePage].content))

	v.WriteString("\n")
	v.WriteString(quitStyle.Render("Press 'q' or 'esc' to quit."))

	return containerStyle.MarginLeft((termWidth / 2) - (lipgloss.Width(joinedTabs) / 2)).Render(v.String())
}

func main() {
	pages := []page{
		{
			key:     "h",
			title:   "home",
			content: "This is the home page",
		},
		{

			key:     "a",
			title:   "about",
			content: "This is the about page",
		},
		{
			key:     "p",
			title:   "projects",
			content: "This is the projects page",
		},
		{
			key:     "c",
			title:   "contact",
			content: "This is the contact page",
		},
	}

	m := model{
		activePage: 0,
		pages:      pages,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to start: %s\n", err)
		os.Exit(1)
	}
}
