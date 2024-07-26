package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type content struct {
	style lipgloss.Style
	model viewport.Model
	pages []page
}

func newContent(pages []page) *content {
	contentStyle := lipgloss.NewStyle().MarginLeft(1)
	initialModel := viewport.New(0, 0)

	f, err := os.ReadFile(pages[0].filepath)
	if err != nil {
		initialModel.SetContent(fmt.Sprintf("Error loading '%s': %s", pages[0], err))
	} else {
		rendered, err := glamour.Render(string(f), "dracula")
		if err != nil {
			initialModel.SetContent(fmt.Sprintf("Error rendering '%s': %s", pages[0], err))
		} else {

			initialModel.SetContent(rendered)
		}
	}

	return &content{
		style: contentStyle,
		model: initialModel,
		pages: pages,
	}
}

func (c *content) view() string {
	return c.style.Render(c.model.View())
}

func (c *content) update(pageNum int) string {
	f, err := os.ReadFile(c.pages[pageNum].filepath)
	if err != nil {
		c.model.SetContent(fmt.Sprintf("Error loading '%s': %s", c.pages[pageNum], err))
	} else {
		rendered, err := glamour.Render(string(f), "dracula")
		if err != nil {
			c.model.SetContent(fmt.Sprintf("Error rendering '%s': %s", c.pages[pageNum], err))
		} else {

			c.model.SetContent(rendered)
		}
	}

	return c.view()
}
