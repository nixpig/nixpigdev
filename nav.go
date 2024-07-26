package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type nav struct {
	style lipgloss.Style
	model list.Model
}

func newNav(pages []page) *nav {
	navStyle := lipgloss.NewStyle().
		MarginTop(1).
		PaddingRight(1)

	var listItems = make([]list.Item, len(pages))
	for i, page := range pages {
		listItems[i] = page
	}

	initialModel := list.New(
		listItems,
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
