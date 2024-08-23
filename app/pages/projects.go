package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/commands"
)

type projectsModel struct {
	title        string
	description  string
	renderer     *lipgloss.Renderer
	md           mdrenderer
	contentWidth int
	projects     commands.FetchProjectsSuccessMsg
}

func NewProjects(
	renderer *lipgloss.Renderer,
	md mdrenderer,
) projectsModel {
	return projectsModel{
		title:       "Projects",
		description: "OSS + personal projects",
		renderer:    renderer,
		md:          md,
		projects:    commands.FetchProjectsSuccessMsg{},
	}
}

func (p projectsModel) Init() tea.Cmd {
	return commands.FetchProjects()
}

func (p projectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.SectionSizeMsg:
		p.contentWidth = msg.Width
		return p, nil

	case commands.FetchProjectsSuccessMsg:
		p.projects = msg
	}

	return p, nil
}

func (p projectsModel) View() string {
	c := strings.Builder{}
	c.WriteString(`
# Projects

## Personal projects`)

	for _, project := range p.projects {
		c.WriteString(fmt.Sprintf("\n\n[%s](%s)\n\n%s\n\n---", project.Name, project.HTMLURL, project.Description))
	}

	return p.md(c.String(), p.contentWidth)
}

func (p projectsModel) Title() string {
	return p.title
}

func (p projectsModel) Description() string {
	return p.description
}

func (p projectsModel) FilterValue() string {
	return fmt.Sprintf("%s %s", p.title, p.description)
}
