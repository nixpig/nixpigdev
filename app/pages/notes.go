package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Notes(renderer *lipgloss.Renderer) Page {
	notes := Page{
		title:       "Notes",
		description: "Blogs and stuff",
		renderer:    renderer,
		content: strings.Join([]string{
			"# Notes",
		}, "\n\n"),
	}

	return notes
}
