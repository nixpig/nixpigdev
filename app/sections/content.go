package sections

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/keys"
)

type contentModel struct {
	style         lipgloss.Style
	viewportModel viewport.Model
	contents      []tea.Model
	renderer      *lipgloss.Renderer
	activePage    int
}

func NewContent(
	renderer *lipgloss.Renderer,
	contents []tea.Model,
) *contentModel {
	contentStyle := renderer.NewStyle()
	viewportModel := viewport.New(0, 0)

	c := &contentModel{
		style:         contentStyle,
		viewportModel: viewportModel,
		contents:      contents,
		renderer:      renderer,
	}

	return c
}

func (c contentModel) Init() tea.Cmd {
	c.contents[c.activePage].Init()
	return nil
}

func (c contentModel) View() string {
	c.viewportModel.SetContent(
		c.contents[c.activePage].View(),
	)

	return c.viewportModel.View()
}

func (c contentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.GlobalKeys.Down):
			fmt.Println("content -> scroll down")
			c.viewportModel.LineDown(5)
			fmt.Println("down percent: ", c.viewportModel.ScrollPercent()*100)
			fmt.Println("total lines: ", c.viewportModel.TotalLineCount())
			fmt.Println("visible lines: ", c.viewportModel.VisibleLineCount())
		case key.Matches(msg, keys.GlobalKeys.Up):
			fmt.Println("content -> scroll up")
			c.viewportModel.LineUp(5)
		}

	// case commands.NavigateToPage:
	// 	fmt.Println("content -> select index")
	// 	fmt.Println((c.viewportModel.Width))
	// 	c.viewportModel.GotoTop()
	// 	c.activePage = int(msg)
	// 	c.contents[c.activePage].Init()
	// 	c.contents[c.activePage], cmd = c.contents[c.activePage].Update(commands.SetContentWidth(c.viewportModel.Width))
	// 	cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		c.viewportModel.Width = msg.Width
		c.viewportModel.Height = msg.Height
		c.contents[c.activePage], cmd = c.contents[c.activePage].Update(msg)
		return c, cmd
	}

	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}
