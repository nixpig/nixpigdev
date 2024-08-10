package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ActivePage int

type mdrenderer func(md string) string

type Page struct {
	title       string
	description string
	content     func(w int, md mdrenderer) string
	renderer    *lipgloss.Renderer
}

func (p Page) Init() tea.Cmd {
	return nil
}

func (p Page) Update(msg tea.Msg) (tea.Msg, tea.Cmd) {
	return nil, nil
}

func (p Page) View(w int, md mdrenderer) string {
	return p.renderer.NewStyle().Render(p.content(w, md))
}

func (p Page) Title() string {
	return p.title
}

func (p Page) Description() string {
	return p.description
}

func (p Page) FilterValue() string {
	return fmt.Sprintf("%s %s", p.title, p.description)
}
