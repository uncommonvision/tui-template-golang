package form

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"tui-template/internal/styles"
)

type Model struct {
	inputs    []textinput.Model
	focused   int
	submitted bool
	width     int
	height    int
}

func NewModel() Model {
	inputs := make([]textinput.Model, 3)

	inputs[0] = textinput.New()
	inputs[0].Placeholder = "Enter your name"
	inputs[0].Focus()

	inputs[1] = textinput.New()
	inputs[1].Placeholder = "Enter your email"

	inputs[2] = textinput.New()
	inputs[2].Placeholder = "Enter your company"

	return Model{
		inputs:  inputs,
		focused: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "tab", "down":
			m.nextInput()
		case "shift+tab", "up":
			m.prevInput()
		case "enter":
			if m.focused == len(m.inputs)-1 {
				m.submitted = true
				return m, tea.Quit
			} else {
				m.nextInput()
			}
		}
	}

	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *Model) nextInput() {
	m.inputs[m.focused].Blur()
	m.focused = (m.focused + 1) % len(m.inputs)
	m.inputs[m.focused].Focus()
}

func (m *Model) prevInput() {
	m.inputs[m.focused].Blur()
	m.focused = (m.focused - 1 + len(m.inputs)) % len(m.inputs)
	m.inputs[m.focused].Focus()
}

func (m Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.submitted {
		return styles.SuccessStyle.Render(fmt.Sprintf(
			"Form submitted successfully!\nName: %s\nEmail: %s\nCompany: %s",
			m.inputs[0].Value(),
			m.inputs[1].Value(),
			m.inputs[2].Value(),
		))
	}

	title := styles.TitleStyle.Render("User Registration Form")

	labels := []string{"Name:", "Email:", "Company:"}

	var formFields []string
	formFields = append(formFields, title)
	formFields = append(formFields, "")

	for i, input := range m.inputs {
		label := styles.InfoStyle.Render(labels[i])
		field := input.View()
		if i == m.focused {
			field = styles.BorderStyle.Render(field)
		}
		formFields = append(formFields, label)
		formFields = append(formFields, field)
		formFields = append(formFields, "")
	}

	help := styles.HelpStyle.Render("Tab/Enter: Next field • Shift+Tab: Previous field • Enter on last field: Submit • q/Esc: Quit")
	formFields = append(formFields, help)

	return strings.Join(formFields, "\n")
}
