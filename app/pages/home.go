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
	work := renderer.NewStyle().
		Width(s.Width / 2).
		PaddingLeft(1).
		PaddingRight(1).
		Render("My day job consists mostly of TypeScript and Java on Azure.")

	bar := renderer.NewStyle().
		Width(s.Width / 2).
		PaddingLeft(1).
		PaddingRight(1).
		Render("In my free time, I'm currently enjoying learning Go and dabbling in Rust.")

	baz := lipgloss.JoinHorizontal(lipgloss.Top, work, bar)

	qux := renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Foreground)).PaddingLeft(1).PaddingRight(1).Render(baz)

	return strings.Join(
		[]string{
			md(`
# Home

I’m a software engineer from the UK, currently working as a _Senior Technical Lead_.
			`),

			qux,
			"\n",
			md(`
I live in the countryside with my beautiful partner, cats and dog, and enjoy going to the gym and binge-watching Netflix.

**Fun facts**
- 🕰 My day starts at 03:00am every morning.
- I collect Toy Story Alien memorabilia.
			`),
			"\n",
		}, "")
}

// 🗓

func (h *home) Title() string {
	return h.title
}

func (h *home) Description() string {
	return h.description
}

func (h *home) FilterValue() string {
	return fmt.Sprintf("%s %s", h.title, h.description)
}
