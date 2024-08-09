package pages

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Page struct {
	title       string
	description string
	content     func(w int, md func(p string) string) string
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

func (p Page) Content(w int, md func(p string) string) string {
	return p.renderer.NewStyle().Render(p.content(w, md))
}
