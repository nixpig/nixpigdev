package sections

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Footer struct {
	style      lipgloss.Style
	model      help.Model
	helpKeyMap help.KeyMap
}

func (f *Footer) Init() tea.Cmd {
	return nil
}

func NewFooter(renderer *lipgloss.Renderer, helpKeyMap help.KeyMap) *Footer {
	footerStyle := renderer.
		NewStyle().
		AlignHorizontal(lipgloss.Center).
		Padding(0, 1)

	initialModel := help.New()

	initialModel.ShortSeparator = " • "
	initialModel.Styles.ShortKey = renderer.NewStyle().Bold(true)
	initialModel.Styles.ShortDesc = renderer.NewStyle().Faint(true)

	return &Footer{
		style:      footerStyle,
		model:      initialModel,
		helpKeyMap: helpKeyMap,
	}
}

func (f *Footer) View() string {
	return f.style.Render(f.model.View(f.helpKeyMap))
}

func (f *Footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return f, nil
}

func (f *Footer) GetHeight() int {
	return f.style.GetHeight()
}
