package app

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/nixpig/nixpigdev/app/commands"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/pages"
	"github.com/nixpig/nixpigdev/app/sections"
)

const (
	heightOffset = 2
	navWidth     = 23
)

type appModel struct {
	renderer *lipgloss.Renderer

	term   string
	width  int
	height int

	navModel      tea.Model
	viewportModel viewport.Model
	footerModel   tea.Model

	pages []tea.Model

	ready      bool
	activePage int
}

func New(pty ssh.Pty, renderer *lipgloss.Renderer) appModel {
	pageModels := []tea.Model{
		pages.NewHome(renderer, md),
		pages.NewScrapbook(renderer, md),
		pages.NewProjects(renderer, md),
		pages.NewResume(renderer, md),
		pages.NewUses(renderer, md),
		pages.NewContact(renderer, md),
	}

	viewportModel := viewport.New(0, 0)

	return appModel{
		term:          pty.Term,
		width:         pty.Window.Width,
		height:        pty.Window.Height,
		navModel:      sections.NewNav(renderer, pageModels),
		footerModel:   sections.NewFooter(renderer, keys.GlobalKeys),
		viewportModel: viewportModel,
		pages:         pageModels,
	}
}

func (m appModel) Init() tea.Cmd {
	m.pages[m.activePage].Init()
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

		case key.Matches(msg, keys.GlobalKeys.Down):
			m.viewportModel.LineDown(1)

		case key.Matches(msg, keys.GlobalKeys.Up):
			m.viewportModel.LineUp(1)
		}

	case tea.WindowSizeMsg:
		m.ready = false

		m.navModel, cmd = m.navModel.Update(commands.SectionSizeMsg{
			Width:  navWidth,
			Height: msg.Height - heightOffset,
		})
		cmds = append(cmds, cmd)

		m.pages[m.activePage], cmd = m.pages[m.activePage].Update(commands.SectionSizeMsg{
			Width:  msg.Width - navWidth,
			Height: msg.Height - heightOffset,
		})
		cmds = append(cmds, cmd)

		m.viewportModel.Width = msg.Width - navWidth
		m.viewportModel.Height = msg.Height - heightOffset
		m.viewportModel.SetContent(m.pages[m.activePage].View())

		m.ready = true

		return m, tea.Batch(cmds...)

	case commands.PageNavigationMsg:
		if int(msg) >= 0 || int(msg) < len(m.pages) {
			m.activePage = int(msg)
			m.viewportModel.Width = m.width - navWidth
			m.viewportModel.Height = m.height - heightOffset

			m.pages[m.activePage], cmd = m.pages[m.activePage].Update(commands.SectionSizeMsg{
				Width:  m.width - navWidth,
				Height: m.height - heightOffset,
			})

			cmd = m.pages[m.activePage].Init()
			cmds = append(cmds, cmd)

			m.viewportModel.SetContent(m.pages[m.activePage].View())
			m.viewportModel.GotoTop()
		}

		return m, tea.Batch(cmds...)
	}

	m.navModel, cmd = m.navModel.Update(msg)
	cmds = append(cmds, cmd)

	m.footerModel, cmd = m.footerModel.Update(msg)
	cmds = append(cmds, cmd)

	m.pages[m.activePage], cmd = m.pages[m.activePage].Update(msg)
	cmds = append(cmds, cmd)

	m.viewportModel.Width = m.width - navWidth
	m.viewportModel.Height = m.height - heightOffset
	m.viewportModel.SetContent(m.pages[m.activePage].View())

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
				m.viewportModel.View(),
			),
			m.footerModel.View(),
		),
	)

	return v.String()
}
