package sections

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/pages"
	"github.com/nixpig/nixpigdev/app/theme"
)

type SelectIndex int

type Nav struct {
	style  lipgloss.Style
	model  list.Model
	Length int
}

func NewNav(renderer *lipgloss.Renderer, contents []pages.Page) *Nav {
	navStyle := renderer.NewStyle().
		MarginTop(1).
		PaddingRight(0)

	var listItems = make([]list.Item, len(contents))
	for i, page := range contents {
		listItems[i] = page
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

	return &Nav{
		style:  navStyle,
		model:  initialModel,
		Length: len(contents),
	}
}

func (n *Nav) View() string {
	return n.style.Render(n.model.View())
}

func (n *Nav) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		n.model.SetWidth(msg.Width)
		n.model.SetHeight(msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.GlobalKeys.Next):
			n.model.CursorDown()
		case key.Matches(msg, keys.GlobalKeys.Prev):
			n.model.CursorUp()
		}

	case SelectIndex:
		n.model.Select(int(msg))
	}

	return nil, nil
}

func (n *Nav) Init() tea.Cmd {
	return nil
}

func (n *Nav) Width() int {
	return n.model.Width()
}
