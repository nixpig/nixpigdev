package sections

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/pages"
)

type Content struct {
	style    lipgloss.Style
	model    viewport.Model
	contents []pages.Page
	renderer *lipgloss.Renderer
}

func NewContent(renderer *lipgloss.Renderer, contents []pages.Page) *Content {
	contentStyle := renderer.NewStyle()
	initialModel := viewport.New(0, 0)

	c := &Content{
		style:    contentStyle,
		model:    initialModel,
		contents: contents,
	}

	c.model.SetContent(
		c.style.Render(
			contents[0].View(
				pages.ContentSize{
					Width:  c.model.Width,
					Height: c.model.Height,
				},
				c.md,
				renderer,
			),
		),
	)

	return c
}

func (c *Content) Init() tea.Cmd {
	return nil
}

func (c *Content) View() string {
	return c.style.Render(c.model.View())
}

func (c *Content) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	fmt.Println(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.InputKeys.Down):
			c.model.LineDown(1)
		case key.Matches(msg, keys.InputKeys.Up):
			c.model.LineUp(1)
		case key.Matches(msg, keys.InputKeys.Next), key.Matches(msg, keys.InputKeys.Prev):
			c.model.GotoTop()
		}

	case pages.ActivePage:
		c.model.SetContent(
			c.style.Render(
				c.contents[msg].View(
					pages.ContentSize{
						Width:  c.model.Width,
						Height: c.model.Height,
					},
					c.md,
					c.renderer,
				),
			),
		)

	case tea.WindowSizeMsg:
		c.model.Width = msg.Width
		c.model.Height = msg.Height
	}

	return c, nil
}

func (c *Content) md(plain string) string {
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
