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
	style      lipgloss.Style
	model      viewport.Model
	contents   []pages.Page
	renderer   *lipgloss.Renderer
	activePage int
}

func NewContent(renderer *lipgloss.Renderer, contents []pages.Page) *Content {
	contentStyle := renderer.NewStyle()
	initialModel := viewport.New(0, 0)

	c := &Content{
		style:    contentStyle,
		model:    initialModel,
		contents: contents,
	}

	return c
}

func (c *Content) Init() tea.Cmd {
	return nil
}

func (c *Content) View() string {
	c.model.SetContent(
		c.contents[c.activePage].View(
			pages.ContentSize{
				Width:  c.model.Width,
				Height: c.model.Height,
			},
			c.md,
			c.renderer,
		),
	)
	return c.model.View()
}

func (c *Content) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.GlobalKeys.Down):
			c.model.LineDown(1)
		case key.Matches(msg, keys.GlobalKeys.Up):
			c.model.LineUp(1)
		case key.Matches(msg, keys.GlobalKeys.Next), key.Matches(msg, keys.GlobalKeys.Prev):
			c.model.GotoTop()
		}

		_, cmd := c.contents[c.activePage].Update(msg)
		return c, cmd

	case pages.ActivePage:
		c.activePage = int(msg)

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
