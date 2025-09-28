package form

import (
	"fmt"
	"strings"

	"tui-template/internal/commands"
	"tui-template/internal/components/keymapview"
	"tui-template/internal/styles"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Next     key.Binding
	Previous key.Binding
	Submit   key.Binding
	Back     key.Binding
}

var DefaultKeyMap = KeyMap{
	Next: key.NewBinding(
		key.WithKeys("tab", "down"),
		key.WithHelp("tab/▼", "next"),
	),
	Previous: key.NewBinding(
		key.WithKeys("shift+tab", "up"),
		key.WithHelp("shift+tab/▲", "previous"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
	Back: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", "back"),
	),
}

type Model struct {
	Width      int
	Height     int
	inputs     []textinput.Model
	focused    int
	keyMapView keymapview.Model
	submitted  bool
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Next):
			m.nextInput()
		case key.Matches(msg, DefaultKeyMap.Previous):
			m.prevInput()
		case key.Matches(msg, DefaultKeyMap.Submit):
			if m.focused == len(m.inputs)-1 {
				m.submitted = true
				// return m, tea.Quit
			} else {
				m.nextInput()
			}
		case key.Matches(msg, DefaultKeyMap.Back):
			cmds = append(cmds, commands.RouteToHome())
		}
	}

	for i := range m.inputs {
		var cmd tea.Cmd
		m.inputs[i], cmd = m.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
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

	formFields = append(formFields, m.keyMapView.View())

	return strings.Join(formFields, "\n")
}
