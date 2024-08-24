package sections

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type footerModel struct {
	style      lipgloss.Style
	helpModel  help.Model
	helpKeyMap help.KeyMap
}

func NewFooter(
	renderer *lipgloss.Renderer,
	helpKeyMap help.KeyMap,
) *footerModel {
	footerStyle := renderer.
		NewStyle().
		AlignHorizontal(lipgloss.Center).
		Padding(0, 1)

	helpModel := help.New()

	helpModel.ShortSeparator = " • "
	helpModel.Styles.ShortKey = renderer.NewStyle().Bold(true)
	helpModel.Styles.ShortDesc = renderer.NewStyle().Faint(true)

	return &footerModel{
		style:      footerStyle,
		helpModel:  helpModel,
		helpKeyMap: helpKeyMap,
	}
}

func (f footerModel) Init() tea.Cmd {
	return nil
}

func (f footerModel) View() string {
	return f.style.Render(f.helpModel.View(f.helpKeyMap))
}

func (f footerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return f, nil
}
