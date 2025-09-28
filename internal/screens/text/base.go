package text

import (
	"tui-template/internal/color"
	"tui-template/internal/components/keymapview"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

type Model struct {
	width      int
	height     int
	value      string
	style      lipgloss.Style
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
		case key.Matches(msg, DefaultKeyMap.Quit):
			cmds = append(cmds, tea.Quit)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.style.
		Align(lipgloss.Center, lipgloss.Center).
		Foreground(color.HotPink).
		Width(m.width).
		Height(m.height)

	_, t, r, b, l := style.GetBorder()

	if t {
		style = style.Height(style.GetHeight() - 1)
	}

	if r {
		style = style.Width(style.GetWidth() - 1)
	}

	if b {
		style = style.Height(style.GetHeight() - 1)
	}

	if l {
		style = style.Width(style.GetWidth() - 1)
	}

	return style.
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.value,
				m.keyMapView.View(),
			),
		)
}
