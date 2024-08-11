package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
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
	helpKeyMap help.KeyMap
	helpModel  help.Model
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
		case key.Matches(msg, keys.FormKeys.Enter):
			if m.focusIndex == len(m.inputs) {
				fmt.Println("submit form")
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

	return m, nil
}

func (m *form) updateInputs(msg tea.Msg) tea.Cmd {
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == m.focusIndex {
			m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
			continue
		}

		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = noStyle
		m.inputs[i].TextStyle = noStyle
	}

	for i := range m.inputs {
		m.inputs[i], _ = m.inputs[i].Update(msg)
	}

	return nil
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
	fmt.Fprintf(&b, "\n\n%s\n\n%s", *button, m.helpModel.View(m.helpKeyMap))

	return b.String()
}
