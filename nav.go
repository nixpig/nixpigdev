package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type nav struct {
	style lipgloss.Style
	model list.Model
}

func newNav(items []list.Item) *nav {
	navStyle := lipgloss.NewStyle().
		MarginTop(1).
		PaddingRight(1)

	initialModel := list.New(
		items,
		list.NewDefaultDelegate(),
		0, 0,
	)

	initialModel.SetShowPagination(false)
	initialModel.SetShowHelp(false)
	initialModel.Title = "@nixpig"
	initialModel.SetFilteringEnabled(false)
	initialModel.SetShowStatusBar(false)

	return &nav{
		style: navStyle,
		model: initialModel,
	}
}

func (n *nav) view() string {
	return n.style.Render(n.model.View())
}

type item struct {
	title string
	desc  string
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.desc
}

func (i item) FilterValue() string {
	return fmt.Sprintf("%s %s", i.title, i.desc)
}
