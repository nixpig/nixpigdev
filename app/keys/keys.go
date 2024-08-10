package keys

import "github.com/charmbracelet/bubbles/key"

type keys struct {
	Quit key.Binding
	Up   key.Binding
	Down key.Binding
	Next key.Binding
	Prev key.Binding
}

func (k keys) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
		k.Next,
		k.Prev,
		k.Quit,
	}
}

func (k keys) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

var InputKeys = keys{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q/esc", "quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("⬆/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("⬇/j", "down"),
	),
	Next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "prev"),
	),
}
