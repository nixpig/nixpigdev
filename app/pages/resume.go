package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Resume(renderer *lipgloss.Renderer) Page {
	var resume = Page{
		title:       "Resume",
		description: "Skills + experience",
		renderer:    renderer,
		content: func(w int, markdown func(p string) string) string {
			// tr, err := glamour.NewTermRenderer(
			// 	glamour.WithStylePath("dracula"),
			// 	glamour.WithWordWrap(w/2-2),
			// )
			// if err != nil {
			// 	return fmt.Sprintf("failed to create term renderer: %s", err)
			// }

			return strings.Join([]string{
				markdown(`
# Résumé

Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.

## Skills
- Leading technical intiatives
- Architectural design and hands-on coding

## Experience

**HR + Payroll SaaS (2018 - Present)**

- _Senior Technical Lead (2024 - Present)_
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
		},
	}

	return resume
}
