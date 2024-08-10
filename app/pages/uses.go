package pages

import "github.com/charmbracelet/lipgloss"

func Uses(renderer *lipgloss.Renderer) Page {
	var uses = Page{
		title:       "Uses",
		description: "Tools of the trade",
		renderer:    renderer,
		content: func(w int, markdown func(p string) string) string {
			return markdown(`
# Uses

I'm a simple person, with simple needs. I spend most of my time in the terminal, and my setup is built around being able to work in that environment (somewhat) efficiently.

## Hardware

- **Keyboard:** [Feker Alice 80](https://fekertech.com/products/feker-alice)
- **Mouse:** The original [Mad Catz RAT 1](https://uk.webuy.com/product-detail?id=0728658050467C)
- **Headset:** [JBL Quantum 810](https://uk.jbl.com/gaming-headsets/JBLQ810WLBLK.html)

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
- **Browser:** [Chromium](https://www.chromium.org)
  - with [Vimium](https://vimium.github.io)

## Dotfiles

- [nixpig/dotfiles](https://github.com/nixpig/dotfiles)
			`)
		},
	}

	return uses
}
