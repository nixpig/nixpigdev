package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
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

	inactiveTabStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 2, 0, 1).Foreground(inactive).BorderForeground(inactive)
	activeTabStyle   = inactiveTabStyle.Bold(true).Foreground(active)
	contentStyle     = lipgloss.NewStyle().Padding(0, 1, 1, 0).Border(lipgloss.NormalBorder()).BorderForeground(inactive)
	footerStyle      = lipgloss.NewStyle().Foreground(faded).AlignHorizontal(lipgloss.Center)

	containerStyle = lipgloss.NewStyle().Align(lipgloss.Center).Padding(2)
)

type page struct {
	title string
	key   string
}

type model struct {
	ready      bool
	pages      []page
	activePage int
	viewport   viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "j", "down":
			m.viewport.LineDown(1)

		case "k", "up":
			m.viewport.LineUp(1)

		default:
			selection := slices.IndexFunc(m.pages, func(p page) bool {
				return p.key == msg.String()
			})

			if selection != -1 {
				m.activePage = selection
			}

			content, err := os.ReadFile(fmt.Sprintf("pages/%s.md", m.pages[m.activePage].title))
			if err != nil {
				m.viewport.SetContent(fmt.Sprintf("Error loading content: %s", err))
			} else {
				md, err := glamour.Render(string(content), "dracula")
				if err != nil {
					m.viewport.SetContent(fmt.Sprintf("Error rendering markdown: %s", err))
				}
				m.viewport.SetContent(md)
			}

			return m, nil
		}

	case tea.WindowSizeMsg:
		tabsHeight := lipgloss.Height(m.tabsView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := tabsHeight + footerHeight

		if !m.ready {
			// -2 for the line breaks between sections
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = tabsHeight + 3
			content, err := os.ReadFile(fmt.Sprintf("pages/%s.md", m.pages[m.activePage].title))
			if err != nil {
				m.viewport.SetContent(fmt.Sprintf("Error loading content: %s", err))
			} else {
				md, err := glamour.Render(string(content), "dracula")
				if err != nil {
					m.viewport.SetContent(fmt.Sprintf("Error rendering markdown: %s", err))
				}
				m.viewport.SetContent(md)
			}
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) tabsView() string {
	renderedTabs := make([]string, len(m.pages))

	for i, page := range m.pages {
		tabContent := fmt.Sprintf("(%s) %s", page.key, page.title)
		if i == m.activePage {
			renderedTabs[i] = activeTabStyle.Render(tabContent)
		} else {
			renderedTabs[i] = inactiveTabStyle.Render(tabContent)
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
}

func (m model) footerView() string {
	helpMsg := "Press ↑↓ or j/k to scroll."
	quitMsg := "Press 'q' or 'esc' to quit."
	return footerStyle.Render(fmt.Sprintf("%s %s", helpMsg, quitMsg))
}

func (m model) View() string {
	if !m.ready {
		return "\nLoading..."
	}

	v := strings.Builder{}

	v.WriteString(m.tabsView())
	v.WriteString("\n")
	v.WriteString(m.viewport.View())
	v.WriteString("\n")
	v.WriteString(m.footerView())

	return v.String()
}

func main() {
	pages := []page{
		{
			key:   "h",
			title: "home",
		},
		{

			key:   "a",
			title: "about",
		},
		{
			key:   "p",
			title: "projects",
		},
		{
			key:   "c",
			title: "contact",
		},
	}

	m := model{
		activePage: 0,
		pages:      pages,
	}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to start: %s\n", err)
		os.Exit(1)
	}
}
