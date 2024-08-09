package pages

import "github.com/charmbracelet/lipgloss"

func Contact(renderer *lipgloss.Renderer) Page {
	contact := Page{
		title:       "Contact",
		description: "Come say hi!",
		renderer:    renderer,
		content: `
# Contact

Feel free to reach out and say "Hi!"

**✉ Email:** [hi@nixpig.dev](mailto:hi@nixpig.dev)
		`,
	}

	return contact
}
