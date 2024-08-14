package pages

import (
	"fmt"

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
	}
}

func (p projectsModel) Init() tea.Cmd {
	return nil
}

func (p projectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.SectionSizeMsg:
		p.contentWidth = msg.Width
		return p, nil
	}

	return p, nil
}

func (p projectsModel) View() string {
	return p.md(`
# Projects

[syringe.sh](https://github.com/nixpig/syringe.sh) • _Go_

Self-hostable distributed database-per-user encrypted secrets management over SSH.

[joubini](https://github.com/nixpig/joubini) • _Rust_

Super-simple to configure HTTP/S reverse proxy for local dev; supports HTTP/1.1, HTTP/2, SSL (+ web sockets coming soon).

[corkscrew](https://github.com/nixpig/corkscrew) • _Rust_

Batch executor for HTTP requests configured in a simple YAML schema.
			`, p.contentWidth)
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
