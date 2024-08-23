package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/commands"
)

type homeModel struct {
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
	return homeModel{
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
	return strings.Join(
		[]string{
			h.md(`
# Home

I’m a software engineer from the UK, currently working as a _Senior Technical Lead_.

I live in the countryside with my beautiful partner, cats and dog.

**Fun facts**
- My day starts at 03:00am every morning.
- I collect Toy Story Alien memorabilia.
			`, h.contentWidth),
			"\n",
		}, "")
}

func (h homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.SectionSizeMsg:
		h.contentWidth = msg.Width
		return h, nil
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
