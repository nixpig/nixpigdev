package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/internal/commands"
)

type resumeModel struct {
	title        string
	description  string
	renderer     *lipgloss.Renderer
	md           mdrenderer
	contentWidth int
}

func NewResume(
	renderer *lipgloss.Renderer,
	md mdrenderer,
) resumeModel {
	return resumeModel{
		title:       "Resume",
		description: "Skills + experience",
		renderer:    renderer,
		md:          md,
	}

}

func (r resumeModel) Init() tea.Cmd {
	return nil
}

func (r resumeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.SectionSizeMsg:
		r.contentWidth = msg.Width
		return r, nil
	}

	return r, nil
}

func (r resumeModel) View() string {
	return strings.Join([]string{
		r.md(`
# Résumé

## Skills
- Architectural design 
- Hands-on coding
- Leading technical intiatives

## Experience

				`, r.contentWidth),
	}, "")
}

func (r resumeModel) Title() string {
	return r.title
}

func (r resumeModel) Description() string {
	return r.description
}

func (r resumeModel) FilterValue() string {
	return fmt.Sprintf("%s %s", r.title, r.description)
}
