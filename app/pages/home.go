package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/theme"
)

func Home(
	renderer *lipgloss.Renderer,
) Page {
	var home = Page{
		title:       "Home",
		description: "Where the ♥ is",
		renderer:    renderer,
		content: func(w int, markdown mdrenderer) string {
			foo := renderer.NewStyle().
				Width(w / 2).
				PaddingLeft(1).
				PaddingRight(1).
				Render("☀ Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.")

			bar := renderer.NewStyle().
				Width(w / 2).
				PaddingLeft(1).
				PaddingRight(1).
				Render("☽ Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.")

			baz := lipgloss.JoinHorizontal(lipgloss.Top, foo, bar)

			qux := renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Foreground)).PaddingLeft(1).PaddingRight(1).Render(baz)

			return strings.Join(
				[]string{
					markdown(`
# Home

I’m a software engineer from the UK, currently working as a _Senior Technical Lead_.

I live in the countryside with my beautiful partner, cats and dog.
					`),
					qux,
					"\n",
					markdown(`
My day job consists mostly of _TypeScript_ and _Java_ on _Azure_.

In my free time, I'm currently enjoying learning **Go** and dabbling in **Rust**.
					`),
				}, "")
		},
	}

	return home
}
