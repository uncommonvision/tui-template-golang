package dynalist

import (
	"tui-template/internal/components/keymapview"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	Down key.Binding
	Up   key.Binding
}

var DefaultKeyMap = KeyMap{
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j", "down"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k", "up"),
	),
}

type Model struct {
	AnnotationStyle         lg.Style
	AnnotationSelectedStyle lg.Style
	ListStyle               lg.Style
	RowStyle                lg.Style
	RowSelectedStyle        lg.Style

	rows          []Row
	selectedIndex int
	width         int
	height        int
	keyMapView    keymapview.Model
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		return RowSelectedMsg{
			Index: m.selectedIndex,
			Row:   m.SelectedRow(),
		}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	m.width = msg.Width
	// 	m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyDown.String(), "j":
			cmds = append(cmds, m.NextRow())
		case tea.KeyUp.String(), "k":
			cmds = append(cmds, m.PrevRow())
			// case "o":
			// 	m.SelectedRow().OpenURL()
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	rows := []string{}

	listitem := func(annotationStyle lg.Style, textStyle lg.Style, row Row) string {
		text := row.title + "\n" + row.description
		height := lg.Height(text) + 1

		return lg.JoinHorizontal(
			lg.Left,
			lg.JoinVertical(
				lg.Top,
				annotationStyle.Render(row.annotation),
			),
			lg.JoinVertical(
				lg.Top,
				textStyle.Height(height).Render(text),
			),
		)
	}

	for i, row := range m.rows {
		if i == m.selectedIndex {
			rows = append(rows, listitem(m.AnnotationSelectedStyle, m.RowSelectedStyle, row))
		} else {
			rows = append(rows, listitem(m.AnnotationStyle, m.RowStyle, row))
		}
	}

	rows = append(rows, m.keyMapView.View())

	return lg.JoinVertical(lg.Top, rows...)
}
