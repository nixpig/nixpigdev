package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/theme"
)

type scrapbook struct {
	title       string
	description string
}

var Scrapbook = scrapbook{
	title:       "Scrapbook",
	description: "Notes, blogs, gists…",
}

func (sb *scrapbook) Init() tea.Cmd {
	return nil
}

func (sb *scrapbook) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (sb *scrapbook) View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string {
	tr, err := glamour.NewTermRenderer(
		glamour.WithStylePath("dracula"),
		glamour.WithWordWrap(s.Width/2-2),
	)
	if err != nil {
		return fmt.Sprintf("Failed to create term renderer: %s", err)
	}

	left := renderer.NewStyle().
		Width(s.Width / 2).
		PaddingRight(1)

	right := renderer.NewStyle().
		Width(s.Width / 2).
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
		md("# Scrapbook"),
		padded.Foreground(lipgloss.Color(theme.Dracula.Foreground)).Render("Just some stuff...\n"),
		row,
	}, "")
}

func (sb *scrapbook) Title() string {
	return sb.title
}

func (sb *scrapbook) Description() string {
	return sb.description
}

func (sb *scrapbook) FilterValue() string {
	return fmt.Sprintf("%s %s", sb.title, sb.description)
}
