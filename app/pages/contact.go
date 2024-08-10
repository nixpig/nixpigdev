package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/keys"
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

func (c *contact) Update(msg tea.Msg) (tea.Msg, tea.Cmd) {
	m, cmd := c.form.Update(msg)
	return m, cmd
}

func (c *contact) View(s ContentSize, md mdrenderer, renderer *lipgloss.Renderer) string {
	return strings.Join([]string{
		md(`
# Contact

Feel free to reach out and say "Hi!"

**✉ Email:** [hi@nixpig.dev](mailto:hi@nixpig.dev)`),

		c.form.View(),
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

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type form struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func NewForm() form {
	m := form{
		inputs: make([]textinput.Model, 3),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Email"
			t.CharLimit = 64
		case 2:
			t.Placeholder = "Message"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
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
		case msg.String() == "tab":
			m.focusIndex++

		case key.Matches(msg, keys.InputKeys.Next):
			m.focusIndex++

		case key.Matches(msg, keys.InputKeys.Prev):
			m.focusIndex--

		}

	default:
		cmds := make([]tea.Cmd, len(m.inputs))
		for i := 0; i <= len(m.inputs)-1; i++ {
			if i == m.focusIndex {
				fmt.Println("focus this one: ", i)
				// Set focused state
				cmds[i] = m.inputs[i].Focus()
				m.inputs[i].PromptStyle = focusedStyle
				m.inputs[i].TextStyle = focusedStyle
				continue
			}
			// Remove focused state
			m.inputs[i].Blur()
			m.inputs[i].PromptStyle = noStyle
			m.inputs[i].TextStyle = noStyle
		}

		return m, tea.Batch(cmds...)

	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *form) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *form) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
