package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/theme"
)

func Scrapbook(renderer *lipgloss.Renderer) Page {
	var scrapbook = Page{
		title:       "Scrapbook",
		description: "Notes, blogs, etc…",
		renderer:    renderer,
		content: func(w int, markdown mdrenderer) string {
			tr, err := glamour.NewTermRenderer(
				glamour.WithStylePath("dracula"),
				glamour.WithWordWrap(w/2-2),
			)
			if err != nil {
				return fmt.Sprintf("Failed to create term renderer: %s", err)
			}

			left := renderer.NewStyle().
				Width(w / 2).
				PaddingRight(1)

			right := renderer.NewStyle().
				Width(w / 2).
				PaddingLeft(1)

			container := renderer.NewStyle()
			padded := renderer.NewStyle().
				PaddingLeft(2).
				PaddingRight(2)

			g, err := tr.Render(`## Recent gists

- [name of gist](https://gist.com/b6222)
- [name of gist](https://gist.com/b6222)
- [name of gist](https://gist.com/b6222)
- [name of gist](https://gist.com/b6222)
- [name of gist](https://gist.com/b6222)
			`)
			if err != nil {
				return fmt.Sprintf("failed to render recent commits: %s", err)
			}
			gists := left.Render(g)

			b, err := tr.Render("## Recent blogs")
			if err != nil {
				return fmt.Sprintf("failed to render recent blogs: %s", err)
			}
			blogs := right.Render(b)

			row := container.Render(
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					gists,
					blogs,
				),
			)

			return strings.Join([]string{
				markdown("# Scrapbook"),
				padded.Foreground(lipgloss.Color(theme.Dracula.Foreground)).Render("Just some stuff...\n"),
				row,
			}, "")
		},
	}

	return scrapbook
}
