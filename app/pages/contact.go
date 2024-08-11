package pages

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/theme"
)

type contact struct {
	title       string
	description string
	form        form
}

var Contact = contact{
	title:       "Contact",
	description: "Come say hi!",
	form:        NewForm(),
}

func (c *contact) Init() tea.Cmd {
	return nil
}

func (c *contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := c.form.Update(msg)
	return nil, cmd
}

func (c *contact) View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string {
	return strings.Join([]string{
		md(`
# Contact

Feel free to reach out and say "Hi!"

**✉ Email:** [hi@nixpig.dev](mailto:hi@nixpig.dev)`),

		c.form.View(renderer),
	}, "")
}

func (c *contact) Title() string {
	return c.title
}

func (c *contact) Description() string {
	return c.description
}

func (c *contact) FilterValue() string {
	return fmt.Sprintf("%s %s", c.title, c.description)
}

var ()

type form struct {
	focusIndex       int
	inputs           []textinput.Model
	cursorMode       cursor.Mode
	helpKeyMap       help.KeyMap
	helpModel        help.Model
	validationErrors []string
}

func NewForm() form {
	helpModel := help.New()
	helpModel.ShortSeparator = " • "
	helpModel.Styles.ShortKey = lipgloss.NewStyle().Bold(true)
	helpModel.Styles.ShortDesc = lipgloss.NewStyle().Faint(true)

	m := form{
		inputs:     make([]textinput.Model, 3),
		helpKeyMap: keys.FormKeys,
		helpModel:  helpModel,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CharLimit = 32

		switch i {
		case 0:
			t.Prompt = "Name:    "
			t.Placeholder = "janedoe"
			t.Focus()
		case 1:
			t.Prompt = "Email:   "
			t.Placeholder = "jane@example.com"
			t.CharLimit = 64
		case 2:
			t.Prompt = "Message: "
			t.Placeholder = "Hi, there!"
			t.CharLimit = 1024
		}

		m.inputs[i] = t
	}

	return m
}

func (m *form) Init() tea.Cmd {
	return textinput.Blink
}

func (m *form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.FormKeys.Enter):
			if m.focusIndex == len(m.inputs) {
				m.validationErrors = []string{}
				if len(m.inputs[0].Value()) == 0 {
					m.validationErrors = append(m.validationErrors, "name: no name provided")
				}
				if _, err := mail.ParseAddress(m.inputs[1].Value()); err != nil {
					m.validationErrors = append(m.validationErrors, err.Error())
				}
				if len(m.inputs[2].Value()) == 0 {
					m.validationErrors = append(m.validationErrors, "message: no message provided")
				}

				fmt.Println("validate and submit form")
			} else {
				m.focusIndex++
			}

		case key.Matches(msg, keys.FormKeys.Next):
			if m.focusIndex < len(m.inputs) {
				m.focusIndex++
			}

		case key.Matches(msg, keys.FormKeys.Prev):
			if m.focusIndex > 0 {
				m.focusIndex--
			}
		}
	}

	m.updateInputs(msg)

	return nil, nil
}

func (m *form) updateInputs(msg tea.Msg) tea.Cmd {
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == m.focusIndex {
			m.inputs[i].Focus()
			continue
		}

		m.inputs[i].Blur()
	}

	for i := range m.inputs {
		m.inputs[i], _ = m.inputs[i].Update(msg)
	}

	return nil
}

func (m *form) View(renderer *lipgloss.Renderer) string {
	focusedStyle := renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Pink))
	blurredStyle := renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Faint))
	noStyle := renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Foreground))

	focusedButton := focusedStyle.Render("[ Submit ]")
	blurredButton := fmt.Sprintf("%s %s %s", noStyle.Render("["), blurredStyle.Render("Submit"), noStyle.Render("]"))

	var b strings.Builder

	for i := range m.inputs {
		if i == m.focusIndex {
			m.inputs[i].Cursor.Style = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
			m.inputs[i].PromptStyle = focusedStyle
		} else {
			m.inputs[i].TextStyle = noStyle
			m.inputs[i].PromptStyle = noStyle
		}
		m.inputs[i].PlaceholderStyle = blurredStyle
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	if len(m.validationErrors) > 0 {
		for _, e := range m.validationErrors {
			b.WriteString(renderer.NewStyle().Foreground(lipgloss.Color(theme.Dracula.Red)).Render(fmt.Sprintf("\n⚠ %s", e)))
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n%s", *button, m.helpModel.View(m.helpKeyMap))

	return renderer.NewStyle().PaddingLeft(2).PaddingRight(2).Render(b.String())
}
