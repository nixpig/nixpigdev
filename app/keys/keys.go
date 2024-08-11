package keys

import "github.com/charmbracelet/bubbles/key"

type globalKeys struct {
	Quit key.Binding
	Up   key.Binding
	Down key.Binding
	Next key.Binding
	Prev key.Binding
}

func (gk globalKeys) ShortHelp() []key.Binding {
	return []key.Binding{
		gk.Up,
		gk.Down,
		gk.Next,
		gk.Prev,
		gk.Quit,
	}
}

func (gk globalKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{gk.ShortHelp()}
}

var GlobalKeys = globalKeys{
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

type formKeys struct {
	Enter key.Binding
	Next  key.Binding
	Prev  key.Binding
}

func (fk formKeys) ShortHelp() []key.Binding {
	return []key.Binding{
		fk.Enter,
		fk.Next,
		fk.Prev,
	}
}

func (fk formKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

var FormKeys = formKeys{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
	Next: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("⬇", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("⬆", "prev"),
	),
}
