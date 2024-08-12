package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/nixpig/nixpigdev/app/commands"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/pages"
	"github.com/nixpig/nixpigdev/app/sections"
)

type state int

const (
	homeView state = iota
	scrapbookView
	projectsView
	resumeView
	usesView
	contactView
)

const (
	heightOffset = 2
	navWidth     = 23
)

type appModel struct {
	renderer *lipgloss.Renderer
	state    state

	term   string
	width  int
	height int

	navModel     tea.Model
	footerModel  tea.Model
	contentModel tea.Model

	pages []tea.Model

	ready      bool
	activePage int
}

func New(pty ssh.Pty, renderer *lipgloss.Renderer) appModel {
	var pages = []tea.Model{
		pages.NewHome(renderer, md),
		// pages.NewScrapbook(renderer, md),
		// pages.NewProjects(renderer, md),
		// pages.NewResume(renderer, md),
		// pages.NewUses(renderer, md),
		// pages.NewContact(renderer, md),
	}

	return appModel{
		term:         pty.Term,
		width:        pty.Window.Width,
		height:       pty.Window.Height,
		contentModel: sections.NewContent(renderer, pages),
		navModel:     sections.NewNav(renderer, pages),
		footerModel:  sections.NewFooter(renderer, keys.GlobalKeys),
		pages:        pages,
	}
}

func (m appModel) Init() tea.Cmd {
	m.navModel.Init()
	m.contentModel.Init()
	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.GlobalKeys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keys.GlobalKeys.Next):
			if m.activePage == len(m.pages)-1 {
				m.activePage = 0
			} else {
				m.activePage++
			}

			updatedNav, navCmd := m.navModel.Update(commands.SelectIndex(m.activePage))
			m.navModel = updatedNav

			updatedContent, contentCmd := m.contentModel.Update(commands.SelectIndex(m.activePage))
			m.contentModel = updatedContent

			return m, tea.Batch(navCmd, contentCmd)

		case key.Matches(msg, keys.GlobalKeys.Prev):
			if m.activePage == 0 {
				m.activePage = len(m.pages) - 1
			} else {
				m.activePage--
			}

			updatedNav, navCmd := m.navModel.Update(commands.SelectIndex(m.activePage))
			m.navModel = updatedNav

			updatedContent, contentCmd := m.contentModel.Update(commands.SelectIndex(m.activePage))
			m.contentModel = updatedContent

			return m, tea.Batch(navCmd, contentCmd)
		}

		m.contentModel.Update(msg)
		return m, nil

	case tea.WindowSizeMsg:
		m.ready = false

		updatedNav, navCmd := m.navModel.Update(tea.WindowSizeMsg{
			Width:  navWidth,
			Height: msg.Height - heightOffset,
		})
		m.navModel = updatedNav

		updatedContentSize, contentSizeCmd := m.contentModel.Update(tea.WindowSizeMsg{
			Width:  msg.Width - navWidth,
			Height: msg.Height - heightOffset,
		})
		m.contentModel = updatedContentSize

		// explicitly call update so that wordwrap is applied
		updatedContentActive, contentActiveCmd := m.contentModel.Update(commands.SelectIndex(m.activePage))
		m.contentModel = updatedContentActive

		m.ready = true

		return m, tea.Sequence(navCmd, contentSizeCmd, contentActiveCmd)
	}

	updatedNav, navCmd := m.navModel.Update(msg)
	m.navModel = updatedNav

	updatedContent, contentCmd := m.contentModel.Update(msg)
	m.contentModel = updatedContent

	return m, tea.Batch(navCmd, contentCmd)
}

func (m appModel) View() string {
	if !m.ready {
		return "\nLoading..."
	}

	v := strings.Builder{}

	v.WriteString(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				m.navModel.View(),
				m.contentModel.View(),
			),
			m.footerModel.View(),
		),
	)

	return v.String()
}
