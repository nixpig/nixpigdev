package app

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func md(content string, wrap int) string {
	tr, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(wrap),
		glamour.WithStylePath("dracula"),
	)
	if err != nil {
		return fmt.Sprintf("Failed to create term renderer: %s", err)
	}
	rendered, err := tr.Render(content)
	if err != nil {
		return fmt.Sprintf("Failed to render '%s': %s", content, err)
	}

	return rendered
}
