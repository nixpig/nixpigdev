package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ContentSize struct {
	Width  int
	Height int
}

type contentCallback func(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string

type ActivePage int

type mdrenderer func(md string) string

type Page struct {
	title       string
	description string
	content     contentCallback
}

func (p Page) Init() tea.Cmd {
	return nil
}

func (p Page) Update(msg tea.Msg) (tea.Msg, tea.Cmd) {
	return nil, nil
}

func (p Page) View(
	s ContentSize,
	md mdrenderer,
	renderer *lipgloss.Renderer,
) string {
	return p.content(s, md, renderer)
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
