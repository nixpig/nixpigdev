package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Home(renderer *lipgloss.Renderer) Page {
	var home = Page{
		title:       "Home",
		description: "Where the ♥ is",
		renderer:    renderer,
		content: strings.Join([]string{
			"# Home",
			"I’m a software engineer from the UK, currently working as a _Senior Technical Lead_.",
			"I live in the countryside with my beautiful partner, cats and dog.",
			"My day job consists mostly of _TypeScript_ and _Java_ on _Azure_.",
			"In my free time, I'm currently enjoying learning **Go** and dabbling in **Rust**.",
		}, "\n\n"),
	}

	return home
}
