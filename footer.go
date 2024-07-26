package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

type footer struct {
	style      lipgloss.Style
	model      help.Model
	helpKeyMap help.KeyMap
}

func newFooter(helpKeyMap help.KeyMap) *footer {
	footerStyle := lipgloss.
		NewStyle().
		AlignHorizontal(lipgloss.Center).
		Padding(0, 1)

	initialModel := help.New()

	initialModel.ShortSeparator = " • "
	initialModel.Styles.ShortKey = lipgloss.NewStyle().Bold(true)
	initialModel.Styles.ShortDesc = lipgloss.NewStyle().Faint(true)

	return &footer{
		style:      footerStyle,
		model:      initialModel,
		helpKeyMap: helpKeyMap,
	}
}

func (f *footer) view() string {
	return f.style.Render(f.model.View(f.helpKeyMap))
}
