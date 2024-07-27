package app

import "github.com/charmbracelet/bubbles/key"

type keys struct {
	quit key.Binding
	up   key.Binding
	down key.Binding
	next key.Binding
	prev key.Binding
}

func (k keys) ShortHelp() []key.Binding {
	return []key.Binding{
		k.up,
		k.down,
		k.next,
		k.prev,
		k.quit,
	}
}

func (k keys) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

var InputKeys = keys{
	quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q/esc", "quit"),
	),
	up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "up"),
	),
	down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "down"),
	),
	next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next"),
	),
	prev: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "prev"),
	),
}
