package sections

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/internal/commands"
	"github.com/nixpig/nixpigdev/internal/keys"
	"github.com/nixpig/nixpigdev/pkg/theme"
)

type navModel struct {
	style     lipgloss.Style
	listModel list.Model
}

func NewNav(
	renderer *lipgloss.Renderer,
	pageModels []tea.Model,
) *navModel {
	navStyle := renderer.NewStyle().
		MarginTop(1).
		PaddingRight(0)

	var listItems = make([]list.Item, len(pageModels))
	for i, page := range pageModels {
		p, ok := page.(list.Item)
		if !ok {
			fmt.Println("cannot assert page as list item")
			continue
		}
		listItems[i] = p
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles = list.NewDefaultItemStyles()

	delegate.Styles.NormalTitle = renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Dracula.Foreground)).
		PaddingLeft(2)

	delegate.Styles.SelectedTitle = renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Dracula.Pink)).
		PaddingLeft(1).
		BorderLeft(true).
		BorderForeground(lipgloss.Color(theme.Dracula.Pink)).
		BorderStyle(lipgloss.ThickBorder())

	delegate.Styles.NormalDesc = renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Dracula.Faint)).
		PaddingLeft(2)

	delegate.Styles.SelectedDesc = renderer.NewStyle().
		Foreground(lipgloss.Color(theme.Dracula.Faint)).
		PaddingLeft(1).
		BorderLeft(true).
		BorderForeground(lipgloss.Color(theme.Dracula.Pink)).
		BorderStyle(lipgloss.ThickBorder())

	initialModel := list.New(
		listItems,
		delegate,
		0, 0,
	)

	promptStyle := renderer.NewStyle().
		Background(lipgloss.Color(theme.Dracula.Purple)).
		Foreground(lipgloss.Color(theme.Dracula.Prompt))

	titleStyle := renderer.NewStyle().
		Background(lipgloss.Color(theme.Dracula.Purple)).
		Foreground(lipgloss.Color(theme.Dracula.Foreground))

	title := fmt.Sprintf(
		" %s%s ",
		promptStyle.Render("$ "),
		titleStyle.Render("ssh nixpig.dev "),
	)

	initialModel.Styles.Title = renderer.NewStyle().
		Background(lipgloss.Color(theme.Dracula.Purple)).
		Foreground(lipgloss.Color(theme.Dracula.Foreground))

	initialModel.Title = title
	initialModel.SetShowPagination(false)
	initialModel.SetShowHelp(false)
	initialModel.SetFilteringEnabled(false)
	initialModel.SetShowStatusBar(false)

	return &navModel{
		style:     navStyle,
		listModel: initialModel,
	}
}

func (n navModel) Init() tea.Cmd {
	return nil
}

func (n navModel) View() string {
	return n.style.Render(n.listModel.View())
}

func (n navModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.SectionSizeMsg:
		n.listModel.SetWidth(msg.Width)
		n.listModel.SetHeight(msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.GlobalKeys.Next):
			if n.listModel.Index() < len(n.listModel.Items())-1 {
				n.listModel.Select(n.listModel.Index() + 1)
			}

			return n, commands.NavigatePageCmd(n.listModel.Index())

		case key.Matches(msg, keys.GlobalKeys.Prev):
			if n.listModel.Index() > 0 {
				n.listModel.Select(n.listModel.Index() - 1)
			}

			return n, commands.NavigatePageCmd(n.listModel.Index())
		}
	}

	return n, nil
}
