package pages

// import (
// 	"fmt"
//
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )
//
// type projects struct {
// 	title       string
// 	description string
// }
//
// var Projects = projects{
// 	title:       "Projects",
// 	description: "OSS + personal projects",
// }
//
// func (p *projects) Init() tea.Cmd {
// 	return nil
// }
//
// func (p *projects) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	return nil, nil
// }
//
// func (p *projects) View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string {
// 	return md(`
// # Projects
//
// [syringe.sh](https://github.com/nixpig/syringe.sh) • _Go_
//
// Self-hostable distributed database-per-user encrypted secrets management over SSH.
//
// [joubini](https://github.com/nixpig/joubini) • _Rust_
//
// Super-simple to configure HTTP/S reverse proxy for local dev; supports HTTP/1.1, HTTP/2, SSL (+ web sockets coming soon).
//
// [corkscrew](https://github.com/nixpig/corkscrew) • _Rust_
//
// Batch executor for HTTP requests configured in a simple YAML schema.
// 			`)
// }
//
// func (p *projects) Title() string {
// 	return p.title
// }
//
// func (p *projects) Description() string {
// 	return p.description
// }
//
// func (p *projects) FilterValue() string {
// 	return fmt.Sprintf("%s %s", p.title, p.description)
// }
