package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ready      bool
	activePage int
	Nav        *nav
	Content    *content
	Footer     *footer
	Renderer   *lipgloss.Renderer
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.Content.model.GotoTop()
				m.Content.update(m.Nav.model.Index())
				m.activePage = m.Nav.model.Index()
			}

		case key.Matches(msg, InputKeys.prev):
			m.Nav.model.CursorUp()
			if m.Nav.model.Index() != m.activePage {
				m.Content.model.GotoTop()
				m.Content.update(m.Nav.model.Index())
				m.activePage = m.Nav.model.Index()
			}
		}

	case tea.WindowSizeMsg:
		// layout on window size
		m.ready = false

		viewportHeight := msg.Height - m.Footer.style.GetHeight() - 2

		m.Nav.model.SetHeight(viewportHeight)
		m.Nav.model.SetWidth(25)

		m.Content.model.Width = msg.Width - m.Nav.model.Width()
		m.Content.model.Height = viewportHeight

		m.ready = true
	}

	return m, nil
}

func (m Model) View() string {
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

func Run() {
}

func LoadFileContent(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Sprintf("Failed to load '%s': %s", filepath, err)
	}

	return string(content)
}
