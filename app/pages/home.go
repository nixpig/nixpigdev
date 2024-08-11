package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/theme"
)

type home struct {
	title       string
	description string
}

var Home = home{
	title:       "Home",
	description: "Where the ♥ is",
}

func (h *home) Init() tea.Cmd {
	return nil
}

func (h *home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (h *home) View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string {
	foo := renderer.NewStyle().
		Width(s.Width / 2).
		PaddingLeft(1).
		PaddingRight(1).
		Render("☀ Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.")

	bar := renderer.NewStyle().
		Width(s.Width / 2).
		PaddingLeft(1).
		PaddingRight(1).
		Render("☽ Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.")

	baz := lipgloss.JoinHorizontal(lipgloss.Top, foo, bar)

	qux := renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Foreground)).PaddingLeft(1).PaddingRight(1).Render(baz)

	return strings.Join(
		[]string{
			md(`
# Home

I’m a software engineer from the UK, currently working as a _Senior Technical Lead_.

I live in the countryside with my beautiful partner, cats and dog.
					`),
			qux,
			"\n",
			md(`
My day job consists mostly of _TypeScript_ and _Java_ on _Azure_.

In my free time, I'm currently enjoying learning **Go** and dabbling in **Rust**.
					`),
		}, "")
}

func (h *home) Title() string {
	return h.title
}

func (h *home) Description() string {
	return h.description
}

func (h *home) FilterValue() string {
	return fmt.Sprintf("%s %s", h.title, h.description)
}
