package pages

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func NewPage(
	title,
	description,
	content string,
	renderer *lipgloss.Renderer,
) Page {
	return Page{
		title:       title,
		description: description,
		content:     content,
		renderer:    renderer,
	}
}

type Page struct {
	title       string
	description string
	content     string
	renderer    *lipgloss.Renderer
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

func (p Page) Content() string {
	return p.renderer.NewStyle().Render(p.content)
}
