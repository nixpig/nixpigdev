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
	pages []string
}

func newContent(pages []string) *content {
	contentStyle := lipgloss.NewStyle().MarginLeft(1)
	initialModel := viewport.New(0, 0)

	page, err := os.ReadFile(pages[0])
	if err != nil {
		initialModel.SetContent(fmt.Sprintf("Error loading '%s': %s", pages[0], err))
	} else {
		rendered, err := glamour.Render(string(page), "dracula")
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
