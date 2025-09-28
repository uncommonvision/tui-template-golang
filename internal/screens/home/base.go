package home

import (
	"fmt"

	"tui-template/internal/commands"
	"tui-template/internal/components/keymapview"
	"tui-template/internal/styles"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	Form key.Binding
	List key.Binding
	Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Form: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "form example"),
	),
	List: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "list example"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

type Model struct {
	width      int
	height     int
	keyMapView keymapview.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Form):
			cmds = append(cmds, commands.RouteToForm())
		case key.Matches(msg, DefaultKeyMap.List):
			cmds = append(cmds, commands.RouteToList())
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return styles.BorderStyle.
		Width(m.width-2).
		Height(m.height-2).
		Align(lipgloss.Center, lipgloss.Center).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				fmt.Sprintf("Header %v %v", m.width, m.height),
				"Home Component",
				m.keyMapView.View(),
			),
		)
}
