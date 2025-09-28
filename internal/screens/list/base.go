package list

import (
	"tui-template/internal/commands"
	"tui-template/internal/components/keymapview"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	list.KeyMap
	Back key.Binding
}

var DefaultKeyMap = KeyMap{
	Back: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", "back"),
	),
}

type Item struct {
	title, desc string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

type Model struct {
	list       list.Model
	keyMapView keymapview.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 3)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Back):
			cmds = append(cmds, commands.RouteToHome())
		default:
			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return lg.JoinVertical(
		lg.Left,
		m.list.View(),
		m.keyMapView.View(),
	)
}
