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
		pages.NewHome(int(homeView), renderer, md),
		// pages.NewScrapbook(renderer, md),
		// pages.NewProjects(renderer, md),
		// pages.NewResume(renderer, md),
		pages.NewUses(int(usesView), renderer, md),
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
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.GlobalKeys.Quit):
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.ready = false

		m.navModel, cmd = m.navModel.Update(tea.WindowSizeMsg{
			Width:  navWidth,
			Height: msg.Height - heightOffset,
		})
		cmds = append(cmds, cmd)

		m.contentModel, cmd = m.contentModel.Update(tea.WindowSizeMsg{
			Width:  msg.Width - navWidth,
			Height: msg.Height - heightOffset,
		})
		cmds = append(cmds, cmd)

		m.ready = true

		return m, tea.Batch(cmds...)
	}

	m.navModel, cmd = m.navModel.Update(msg)
	cmds = append(cmds, cmd)

	m.contentModel, cmd = m.contentModel.Update(msg)
	cmds = append(cmds, cmd)

	m.footerModel, cmd = m.footerModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
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
