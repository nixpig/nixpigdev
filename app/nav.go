package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type nav struct {
	style lipgloss.Style
	model list.Model
}

func NewNav(renderer *lipgloss.Renderer, pages []Page) *nav {
	navStyle := renderer.NewStyle().
		MarginTop(1).
		PaddingRight(1)

	var listItems = make([]list.Item, len(pages))
	for i, page := range pages {
		listItems[i] = page
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles = list.NewDefaultItemStyles()

	delegate.Styles.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F8F8F2")).
		PaddingLeft(2)

	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF79C6")).
		PaddingLeft(1).
		BorderLeft(true).
		BorderForeground(lipgloss.Color("#FF79C6")).
		BorderStyle(lipgloss.NormalBorder())

	delegate.Styles.NormalDesc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a5a6a7")).
		PaddingLeft(2)

	delegate.Styles.SelectedDesc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#e58ac0")).
		PaddingLeft(1).
		BorderLeft(true).
		BorderForeground(lipgloss.Color("#FF79C6")).
		BorderStyle(lipgloss.NormalBorder())

	initialModel := list.New(
		listItems,
		delegate,
		0, 0,
	)

	promptStyle := renderer.NewStyle().
		Background(lipgloss.Color("#BD93F9")).
		Foreground(lipgloss.Color("#8e50e6"))

	titleStyle := renderer.NewStyle().
		Background(lipgloss.Color("#BD93F9")).
		Foreground(lipgloss.Color("#F8F8F2"))

	t := strings.Join([]string{
		" ",
		promptStyle.Render("$ "),
		titleStyle.Render("ssh nixpig.dev "),
	}, "")

	initialModel.SetShowPagination(false)
	initialModel.SetShowHelp(false)
	initialModel.Title = t
	initialModel.SetFilteringEnabled(false)
	initialModel.SetShowStatusBar(false)

	// initialModel.Styles.Title = renderer.NewStyle().Foreground(lipgloss.Color("#BD93F9"))

	initialModel.Styles.Title = renderer.NewStyle().
		Background(lipgloss.Color("#BD93F9")).
		Foreground(lipgloss.Color("#F8F8F2"))

	return &nav{
		style: navStyle,
		model: initialModel,
	}
}

func (n *nav) view() string {
	return n.style.Render(n.model.View())
}
