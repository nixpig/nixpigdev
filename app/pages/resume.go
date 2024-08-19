package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/commands"
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

Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.

## Skills
- Architectural design 
- Hands-on coding
- Leading technical intiatives

## Experience

**HR + Payroll SaaS (2018 - Present)**

- _Senior Technical Lead, IC (2024 - Present)_
- _Senior Engineering Lead (2021 - 2024)_
- _Senior Engineer (2018 - 2021)_

**Contractor (2016 - 2018)** - _Front End Developer_

**Design agency (2015 - 2016)** - _Front End Developer_

**Travel booking website (2014 - 2015)** - _Front End Developer_

**Freelance (2011 - 2014)** - _Web Developer_

## Education

**University (2012 - 2014)** - _Engineering, Foundation Deg._

**College (2011 - 2012)** - _ICT_
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
