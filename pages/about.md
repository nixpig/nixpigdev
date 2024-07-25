# About

Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.

```go
package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var (
	background      = lipgloss.Color("#11111b")
	backgroundAlt   = lipgloss.Color("#181825")
	backgroundTrans = lipgloss.Color("#aa181825")
	foreground      = lipgloss.Color("#cdd6f4")
	primary         = lipgloss.Color("#a6e3a1")
	secondary       = lipgloss.Color("#eba0ac")
	alert           = lipgloss.Color("#f38ba8")
	active          = lipgloss.Color("#cdd6f4")
	disabled        = lipgloss.Color("#a6adc8")
	inactive        = lipgloss.Color("#a6adc8")
	faded           = lipgloss.Color("#a6adc8")

	inactiveTabStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 2, 0, 1).Foreground(inactive).BorderForeground(inactive)
	activeTabStyle   = inactiveTabStyle.Bold(true).Foreground(active)
	contentStyle     = lipgloss.NewStyle().Padding(0, 1, 1, 0).Border(lipgloss.NormalBorder()).BorderForeground(inactive)
	quitStyle        = lipgloss.NewStyle().Foreground(faded)

	containerStyle = lipgloss.NewStyle().Align(lipgloss.Center).Padding(2)
)

type page struct {
	title string
	key   string
}
```
