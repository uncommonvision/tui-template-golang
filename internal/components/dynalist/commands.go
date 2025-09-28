package dynalist

import tea "github.com/charmbracelet/bubbletea"

type RowSelectedMsg struct {
	Index int
	Row   Row
}

func rowSelected(model *Model) tea.Cmd {
	return func() tea.Msg {
		return RowSelectedMsg{
			Index: model.selectedIndex,
			Row:   model.SelectedRow(),
		}
	}
}
