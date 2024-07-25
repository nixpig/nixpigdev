package main

import "github.com/charmbracelet/bubbles/key"

type keys struct {
	quit key.Binding
}

var inputKeys = keys{
	quit: key.NewBinding(
	// key.WithKeys("")
	),
}
