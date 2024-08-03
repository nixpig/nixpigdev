package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type content struct {
	style lipgloss.Style
	model viewport.Model
	pages []Page
}

func NewContent(renderer *lipgloss.Renderer, pages []Page) *content {
	contentStyle := renderer.NewStyle()
	initialModel := viewport.New(0, 0)

	c := &content{
		style: contentStyle,
		model: initialModel,
		pages: pages,
	}

	c.model.SetContent(c.style.Render(c.md(pages[0].Content)))

	return c
}

func (c *content) view() string {
	return c.style.Render(c.model.View())
}

func (c *content) update(pageNum int) string {
	c.model.GotoTop()
	c.model.SetContent(c.style.Render(c.md(c.pages[pageNum].Content)))

	return c.view()
}

func (c *content) md(plain string) string {
	tr, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(c.model.Width),
		glamour.WithStylePath("dracula"),
	)
	if err != nil {
		return fmt.Sprintf("Failed to create term renderer: %s", err)
	}
	rendered, err := tr.Render(plain)
	if err != nil {
		return fmt.Sprintf("Failed to render '%s': %s", plain, err)
	}

	return rendered
}
