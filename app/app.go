package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/pages"
	"github.com/nixpig/nixpigdev/app/sections"
)

func New(pty ssh.Pty, renderer *lipgloss.Renderer) model {
	var pages = []pages.Page{
		&pages.Home,
		&pages.Scrapbook,
		&pages.Projects,
		&pages.Resume,
		&pages.Uses,
		&pages.Contact,
	}

	return model{
		Term:         pty.Term,
		Width:        pty.Window.Width,
		Height:       pty.Window.Height,
		ContentModel: sections.NewContent(renderer, pages),
		NavModel:     sections.NewNav(renderer, pages),
		FooterModel:  sections.NewFooter(renderer, keys.GlobalKeys),
		pages:        pages,
	}
}

type model struct {
	Term         string
	Width        int
	Height       int
	NavModel     *sections.Nav
	ContentModel *sections.Content
	FooterModel  *sections.Footer
	Renderer     *lipgloss.Renderer
	pages        []pages.Page

	ready      bool
	activePage int
}

func (m model) Init() tea.Cmd {
	m.ContentModel.Init()
	m.NavModel.Init()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.GlobalKeys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keys.GlobalKeys.Next):
			if m.activePage == m.NavModel.Length-1 {
				m.activePage = 0
			} else {
				m.activePage++
			}

			updatedNav, navCmd := m.NavModel.Update(sections.SelectIndex(m.activePage))
			navModel, ok := updatedNav.(*sections.Nav)
			if ok {
				m.NavModel = navModel
			}

			updatedContent, contentCmd := m.ContentModel.Update(pages.ActivePage(m.activePage))
			contentModel, ok := updatedContent.(*sections.Content)
			if ok {
				m.ContentModel = contentModel
			}

			return m, tea.Batch(navCmd, contentCmd)

		case key.Matches(msg, keys.GlobalKeys.Prev):
			if m.activePage == 0 {
				m.activePage = m.NavModel.Length - 1
			} else {
				m.activePage--
			}

			updatedNav, navCmd := m.NavModel.Update(sections.SelectIndex(m.activePage))
			navModel, ok := updatedNav.(*sections.Nav)
			if ok {
				m.NavModel = navModel
			}

			updatedContent, contentCmd := m.ContentModel.Update(pages.ActivePage(m.activePage))
			contentModel, ok := updatedContent.(*sections.Content)
			if ok {
				m.ContentModel = contentModel
			}

			return m, tea.Batch(navCmd, contentCmd)
		}

		m.ContentModel.Update(msg)
		return m, nil

	case tea.WindowSizeMsg:
		m.ready = false

		viewportHeight := msg.Height - m.FooterModel.Height() - 2

		updatedNav, navCmd := m.NavModel.Update(tea.WindowSizeMsg{
			Width:  23,
			Height: viewportHeight,
		})

		navModel, ok := updatedNav.(*sections.Nav)
		if ok {
			m.NavModel = navModel
		}

		updatedContentSize, contentSizeCmd := m.ContentModel.Update(tea.WindowSizeMsg{
			Width:  msg.Width - m.NavModel.Width(),
			Height: viewportHeight,
		})
		contentSizeModel, ok := updatedContentSize.(*sections.Content)
		if ok {
			m.ContentModel = contentSizeModel
		}

		// explicitly call update so that wordwrap is applied
		updatedContentActive, contentActiveCmd := m.ContentModel.Update(pages.ActivePage(m.activePage))
		contentActiveModel, ok := updatedContentActive.(*sections.Content)
		if ok {
			m.ContentModel = contentActiveModel
		}

		m.ready = true

		return m, tea.Sequence(navCmd, contentSizeCmd, contentActiveCmd)
	}

	updatedNav, navCmd := m.NavModel.Update(msg)
	navModel, ok := updatedNav.(*sections.Nav)
	if ok {
		m.NavModel = navModel
	}

	updatedContent, contentCmd := m.ContentModel.Update(msg)
	contentModel, ok := updatedContent.(*sections.Content)
	if ok {
		m.ContentModel = contentModel
	}

	return m, tea.Batch(navCmd, contentCmd)
}

func (m model) View() string {
	if !m.ready {
		return "\nLoading..."
	}

	v := strings.Builder{}

	v.WriteString(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				m.NavModel.View(),
				m.ContentModel.View(),
			),
			m.FooterModel.View(),
		),
	)

	return v.String()
}
