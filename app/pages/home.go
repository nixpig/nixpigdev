package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/commands"
	"github.com/nixpig/nixpigdev/app/theme"
)

type homeModel struct {
	style        lipgloss.Style
	title        string
	description  string
	renderer     *lipgloss.Renderer
	md           mdrenderer
	contentWidth int
}

func NewHome(
	renderer *lipgloss.Renderer,
	md mdrenderer,
) homeModel {
	homeStyle := renderer.NewStyle()

	return homeModel{
		style:       homeStyle,
		title:       "Home",
		description: "Where the ♥ is",
		renderer:    renderer,
		md:          md,
	}

}

func (h homeModel) Init() tea.Cmd {
	return nil
}

func (h homeModel) View() string {
	work := h.renderer.NewStyle().
		Width(h.contentWidth / 2).
		PaddingLeft(1).
		PaddingRight(1).
		Render("My day job consists mostly of TypeScript and Java on Azure.")

	bar := h.renderer.NewStyle().
		Width(h.contentWidth / 2).
		PaddingLeft(1).
		PaddingRight(1).
		Render("In my free time, I'm currently enjoying learning Go and dabbling in Rust.")

	baz := lipgloss.JoinHorizontal(lipgloss.Top, work, bar)

	qux := h.renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Foreground)).PaddingLeft(1).PaddingRight(1).Render(baz)

	return strings.Join(
		[]string{
			h.md(`
# Home

I’m a software engineer from the UK, currently working as a _Senior Technical Lead_.
			`, h.contentWidth),

			qux,
			"\n",
			h.md(`
I live in the countryside with my beautiful partner, cats and dog, and enjoy going to the gym and binge-watching Netflix.

**Fun facts**
- My day starts at 03:00am every morning.
- I collect Toy Story Alien memorabilia.
			`, h.contentWidth),
			"\n",
		}, "")
}

func (h homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h.contentWidth = msg.Width
	case commands.SetContentWidth:
		h.contentWidth = int(msg)
	}

	return h, nil
}

// 🗓

func (h homeModel) Title() string {
	return h.title
}

func (h homeModel) Description() string {
	return h.description
}

func (h homeModel) FilterValue() string {
	return fmt.Sprintf("%s %s", h.title, h.description)
}
