package sections

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/commands"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/theme"
)

type navModel struct {
	style     lipgloss.Style
	listModel list.Model
}

func NewNav(
	renderer *lipgloss.Renderer,
	contents []tea.Model,
) *navModel {
	navStyle := renderer.NewStyle().
		MarginTop(1).
		PaddingRight(0)

	var listItems = make([]list.Item, len(contents))
	for i, page := range contents {
		p, ok := page.(list.Item)
		if !ok {
			fmt.Println("cannot type assert page to list item")
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
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		n.listModel.SetWidth(msg.Width)
		n.listModel.SetHeight(msg.Height)

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, keys.GlobalKeys.Next):
			if n.listModel.Index() < len(n.listModel.Items())-1 {
				n.listModel.Select(n.listModel.Index() + 1)
			}
			// TODO: send view enum nav command
			cmd = func() tea.Msg { return commands.NavigateToPage(n.listModel.Index() + 1) }
			cmds = append(cmds, cmd)
			fmt.Println("nav -> send down command: ", n.listModel.SelectedItem())

		case key.Matches(msg, keys.GlobalKeys.Prev):
			if n.listModel.Index() > 0 {
				n.listModel.Select(n.listModel.Index() - 1)
			}
			// TODO: send view enum nav command
			cmd = func() tea.Msg { return commands.NavigateToPage(n.listModel.Index() - 1) }
			cmds = append(cmds, cmd)
			fmt.Println("nav -> send up command: ", n.listModel.SelectedItem())
		}
	}

	return n, tea.Batch(cmds...)
}
