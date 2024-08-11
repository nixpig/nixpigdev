package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type resume struct {
	title       string
	description string
}

var Resume = resume{
	title:       "Resume",
	description: "Skills + experience",
}

func (r *resume) Init() tea.Cmd {
	return nil
}

func (r *resume) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (r *resume) View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string {

	// tr, err := glamour.NewTermRenderer(
	// 	glamour.WithStylePath("dracula"),
	// 	glamour.WithWordWrap(w/2-2),
	// )
	// if err != nil {
	// 	return fmt.Sprintf("failed to create term renderer: %s", err)
	// }

	return strings.Join([]string{
		md(`
# Résumé

Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.

## Skills
- Architectural design and hands-on coding
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
				`),
	}, "")
}

func (r *resume) Title() string {
	return r.title
}

func (r *resume) Description() string {
	return r.description
}

func (r *resume) FilterValue() string {
	return fmt.Sprintf("%s %s", r.title, r.description)
}
