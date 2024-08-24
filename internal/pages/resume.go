package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/internal/commands"
	"github.com/nixpig/nixpigdev/pkg/markdown"
)

type resumeModel struct {
	title        string
	description  string
	termRenderer *lipgloss.Renderer
	mdRenderer   markdown.Renderer
	contentWidth int
}

func NewResume(
	termRenderer *lipgloss.Renderer,
	mdRenderer markdown.Renderer,
) resumeModel {
	return resumeModel{
		title:        "Resume",
		description:  "Skills + experience",
		termRenderer: termRenderer,
		mdRenderer:   mdRenderer,
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
		r.mdRenderer(`
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
