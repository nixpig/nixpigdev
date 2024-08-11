package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ContentSize struct {
	Width  int
	Height int
}

type contentCallback func(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string

type ActivePage int

type mdrenderer func(md string) string

type Page interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string
	Title() string
	Description() string
	FilterValue() string
}

// func (p Page) FilterValue() string {
// }
//
// func (p Page) Init() tea.Cmd {
// 	return nil
// }
//
// func (p Page) Update(msg tea.Msg) (tea.Msg, tea.Cmd) {
// 	// TODO: somehow need to propagate Update to child bubbles, e.g. Contact form
// 	// Maybe just turn Page into an interface that Contact, Home, etc, need to implement??
//
// 	return nil, nil
// }

// func (p Page) View(
// 	s ContentSize,
// 	md mdrenderer,
// 	renderer *lipgloss.Renderer,
// ) string {
// 	return p.content(s, md, renderer)
// }
//
// func (p Page) Title() string {
// 	return p.title
// }
//
// func (p Page) Description() string {
// 	return p.description
// }
//
// func (p Page) FilterValue() string {
// 	return fmt.Sprintf("%s %s", p.title, p.description)
// }
