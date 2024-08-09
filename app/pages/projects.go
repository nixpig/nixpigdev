package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Projects(renderer *lipgloss.Renderer) Page {
	var projects = Page{
		title:       "Projects",
		description: "OSS + personal work",
		renderer:    renderer,
		content: strings.Join([]string{
			"# Projects",
			"[syringe.sh](https://github.com/nixpig/syringe.sh) • _Go_",
			"Self-hostable distributed database-per-user encrypted secrets management over SSH.",
			"[joubini](https://github.com/nixpig/joubini) • _Rust_",
			"Super-simple to configure HTTP/S reverse proxy for local dev; supports HTTP/1.1, HTTP/2, SSL (+ web sockets coming soon).",
			"[corkscrew](https://github.com/nixpig/corkscrew) • _Rust_",
			"Batch executor for HTTP requests configured in a simple YAML schema.",
		}, "\n\n",
		),
	}

	return projects
}
