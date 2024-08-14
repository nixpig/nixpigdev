package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/commands"
)

type usesModel struct {
	title        string
	description  string
	renderer     *lipgloss.Renderer
	md           mdrenderer
	contentWidth int
}

func NewUses(
	renderer *lipgloss.Renderer,
	md mdrenderer,
) usesModel {
	return usesModel{
		title:       "Uses",
		description: "Tools of the trade",
		renderer:    renderer,
		md:          md,
	}
}

func (u usesModel) Init() tea.Cmd {
	return nil
}

func (u usesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.SectionSizeMsg:
		u.contentWidth = msg.Width
		return u, nil
	}

	return u, nil
}

func (u usesModel) View() string {
	return u.md(`
# Uses

I'm a simple person, with simple needs. I spend most of my time in the terminal, and my setup is built around being able to work in that environment (somewhat) efficiently.

## Software

- **OS:** [Arch Linux](https://archlinux.org)
- **Editor:** [Neovim](https://neovim.io)
- **Terminal:** [Terminator](https://gnome-terminator.org)
- **Shell:** [Bash](https://www.gnu.org/software/bash)
- **Prompt:** [Starship](https://starship.rs)
- **Multiplexer:** [Tmux](https://github.com/tmux/tmux)
- **Window manager:** [i3](https://i3wm.org)
- **Dotfiles manager:** [yadm](https://yadm.io)
- **Font:** [Operator Mono](https://typography.com/fonts/operator/overview)
- **Browser:** [Chromium](https://www.chromium.org) with [Vimium](https://vimium.github.io)

## Dotfiles

- [nixpig/dotfiles](https://github.com/nixpig/dotfiles)

## Peripherals

- **Keyboard:** [Feker Alice 80](https://fekertech.com/products/feker-alice)
- **Mouse:** The original [Mad Catz RAT 1](https://uk.webuy.com/product-detail?id=0728658050467C)
- **Headset:** [JBL Quantum 810](https://uk.jbl.com/gaming-headsets/JBLQ810WLBLK.html)

			`, u.contentWidth)
}

func (u usesModel) Title() string {
	return u.title
}

func (u usesModel) Description() string {
	return u.description
}

func (u usesModel) FilterValue() string {
	return fmt.Sprintf("%s %s", u.title, u.description)
}
