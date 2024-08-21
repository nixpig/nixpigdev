package pages

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/nixpigdev/app/commands"
	"github.com/nixpig/nixpigdev/app/keys"
	"github.com/nixpig/nixpigdev/app/theme"
)

type contactModel struct {
	title        string
	description  string
	formModel    tea.Model
	renderer     *lipgloss.Renderer
	md           mdrenderer
	contentWidth int
}

func NewContact(
	renderer *lipgloss.Renderer,
	md mdrenderer,
) contactModel {
	return contactModel{
		title:       "Contact",
		description: "Come say hi!",
		formModel:   NewForm(renderer),
		renderer:    renderer,
		md:          md,
	}
}

func (c contactModel) Init() tea.Cmd {
	return nil
}

func (c contactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case commands.SectionSizeMsg:
		c.contentWidth = msg.Width
		return c, nil
	}

	c.formModel, cmd = c.formModel.Update(msg)
	cmds = append(cmds, cmd)
	return c, tea.Batch(cmds...)
}

func (c contactModel) View() string {
	return strings.Join([]string{
		c.md(`
# Contact

Feel free to reach out and say "Hi!"

**✉ Email:** [hi@nixpig.dev](mailto:hi@nixpig.dev)`, c.contentWidth),

		c.formModel.View(),
	}, "")
}

func (c contactModel) Title() string {
	return c.title
}

func (c contactModel) Description() string {
	return c.description
}

func (c contactModel) FilterValue() string {
	return fmt.Sprintf("%s %s", c.title, c.description)
}

type formInputs struct {
	name    textinput.Model
	email   textinput.Model
	message textarea.Model
}

type form struct {
	focusIndex       int
	inputs           formInputs
	cursorMode       cursor.Mode
	helpKeyMap       help.KeyMap
	helpModel        help.Model
	validationErrors []string
	sending          bool
	spinner          spinner.Model
	renderer         *lipgloss.Renderer
}

func NewForm(renderer *lipgloss.Renderer) form {
	helpModel := help.New()
	helpModel.ShortSeparator = " • "
	helpModel.Styles.ShortKey = lipgloss.NewStyle().Bold(true)
	helpModel.Styles.ShortDesc = lipgloss.NewStyle().Faint(true)

	nameInput := textinput.New()
	nameInput.CharLimit = 32
	nameInput.Prompt = "Name:    "
	nameInput.Placeholder = "janedoe"
	nameInput.Focus()

	emailInput := textinput.New()
	emailInput.CharLimit = 64
	emailInput.Prompt = "Email:   "
	emailInput.Placeholder = "jane@example.com"

	messageInput := textarea.New()
	messageInput.ShowLineNumbers = false
	messageInput.Placeholder = "Hi, there!"
	messageInput.CharLimit = 1024

	m := form{
		inputs: formInputs{
			name:    nameInput,
			email:   emailInput,
			message: messageInput,
		},
		helpKeyMap: keys.FormKeys,
		helpModel:  helpModel,
		sending:    false,
		spinner:    newSpinner(),
		renderer:   renderer,
	}

	return m
}

func newSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Points
	return s
}

func (m form) Init() tea.Cmd {
	m.spinner.Tick()
	return nil
}

func (m form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case commands.SendEmailSuccessMsg:
		m.inputs.name.SetValue("")
		m.inputs.email.SetValue("")
		m.inputs.message.SetValue("")
		m.validationErrors = []string{}
		m.sending = false
		m.spinner = newSpinner()

	case commands.SendEmailErrMsg:
		m.validationErrors = append(m.validationErrors, msg.Error())
		m.sending = false
		m.spinner = newSpinner()

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		if m.sending {
			return m, nil
		}
		switch {
		case key.Matches(msg, keys.FormKeys.Next):
			if m.focusIndex == 3 {
				m.focusIndex = 0
			} else {
				m.focusIndex++
			}
		case key.Matches(msg, keys.FormKeys.Enter):
			if m.focusIndex == 3 {
				m.validationErrors = []string{}
				if len(m.inputs.name.Value()) == 0 {
					m.validationErrors = append(
						m.validationErrors,
						"name: no name provided",
					)
				}
				if _, err := mail.ParseAddress(m.inputs.email.Value()); err != nil {
					m.validationErrors = append(
						m.validationErrors,
						err.Error(),
					)
				}
				if len(m.inputs.message.Value()) == 0 {
					m.validationErrors = append(
						m.validationErrors,
						"message: no message provided",
					)
				}

				if len(m.validationErrors) == 0 {
					cmd = commands.SendEmail(
						m.inputs.name.Value(),
						m.inputs.email.Value(),
						m.inputs.message.Value(),
					)
					m.sending = true

					cmds = append(cmds, cmd)
					cmds = append(cmds, m.spinner.Tick)
				}
			} else {
				m.focusIndex++
			}
		}
	}

	m.updateInputs(msg)

	return m, tea.Batch(cmds...)
}

func (m *form) updateInputs(msg tea.Msg) tea.Cmd {
	m.inputs.name.Blur()
	m.inputs.email.Blur()
	m.inputs.message.Blur()

	switch m.focusIndex {
	case 0:
		m.inputs.name.Focus()
		m.inputs.name, _ = m.inputs.name.Update(msg)
	case 1:
		m.inputs.email.Focus()
		m.inputs.email, _ = m.inputs.email.Update(msg)
	case 2:
		m.inputs.message.Focus()
		m.inputs.message, _ = m.inputs.message.Update(msg)
	}

	return nil
}

func (m form) View() string {
	focusedStyle := m.renderer.
		NewStyle().
		Foreground(lipgloss.Color(theme.Dracula.Pink))

	blurredStyle := m.renderer.
		NewStyle().
		Foreground(lipgloss.Color(theme.Dracula.Faint))

	noStyle := m.renderer.
		NewStyle().
		Foreground(lipgloss.Color(theme.Dracula.Foreground))

	focusedButton := focusedStyle.Render("[ Send ]")
	blurredButton := fmt.Sprintf(
		"%s %s %s",
		noStyle.Render("["),
		blurredStyle.Render("Send"),
		noStyle.Render("]"),
	)

	var b strings.Builder

	m.inputs.name.TextStyle = noStyle
	m.inputs.name.PromptStyle = noStyle
	m.inputs.email.TextStyle = noStyle
	m.inputs.email.PromptStyle = noStyle

	switch m.focusIndex {
	case 0:
		m.inputs.name.Cursor.Style = focusedStyle
		m.inputs.name.TextStyle = focusedStyle
		m.inputs.name.PromptStyle = focusedStyle

	case 1:
		m.inputs.email.Cursor.Style = focusedStyle
		m.inputs.email.TextStyle = focusedStyle
		m.inputs.email.PromptStyle = focusedStyle

	case 2:
		m.inputs.message.Cursor.Style = focusedStyle
	}

	m.inputs.name.PlaceholderStyle = blurredStyle
	m.inputs.email.PlaceholderStyle = blurredStyle

	b.WriteString(fmt.Sprintf("%s\n", m.inputs.name.View()))
	b.WriteString(fmt.Sprintf("%s\n", m.inputs.email.View()))
	b.WriteString(noStyle.Render("Message:\n"))
	b.WriteString(fmt.Sprintf("\n%s\n", m.inputs.message.View()))

	if len(m.validationErrors) > 0 {
		for _, e := range m.validationErrors {
			b.WriteString(
				m.renderer.
					NewStyle().
					Foreground(lipgloss.Color(theme.Dracula.Red)).
					Render(fmt.Sprintf("\n⚠ %s", e)),
			)
		}

		b.WriteRune('\n')
	}

	button := &blurredButton
	if m.focusIndex == 3 {
		button = &focusedButton
	}
	if m.sending {
		sending := fmt.Sprintf(
			"%s %s %s",
			blurredStyle.Render("["),
			focusedStyle.Render(m.spinner.View()),
			blurredStyle.Render("]"),
		)
		button = &sending
	}

	b.WriteString(fmt.Sprintf(
		"\n%s\n\n%s",
		*button,
		m.helpModel.View(m.helpKeyMap),
	))

	return m.renderer.NewStyle().
		PaddingLeft(2).
		PaddingRight(2).
		Render(b.String())
}
