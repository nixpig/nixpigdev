package pages

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ContentSize struct {
	Width  int
	Height int
}

type contentCallback func(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string

type mdrenderer func(content string, wrap int) string

type Page interface {
	list.Item
	tea.Model
}
