package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

type footer struct {
	style      lipgloss.Style
	model      help.Model
	helpKeyMap help.KeyMap
}

func newFooter(renderer *lipgloss.Renderer, helpKeyMap help.KeyMap) *footer {
	footerStyle := renderer.
		NewStyle().
		AlignHorizontal(lipgloss.Center).
		Padding(0, 1)

	initialModel := help.New()

	initialModel.ShortSeparator = " • "
	initialModel.Styles.ShortKey = renderer.NewStyle().Bold(true)
	initialModel.Styles.ShortDesc = renderer.NewStyle().Faint(true)

	return &footer{
		style:      footerStyle,
		model:      initialModel,
		helpKeyMap: helpKeyMap,
	}
}

func (f *footer) view() string {
	return f.style.Render(f.model.View(f.helpKeyMap))
}
