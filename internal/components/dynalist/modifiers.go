package dynalist

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) InsertRow(rowItem *RowItem) {
	m.rows = append(m.rows, NewRow(rowItem))
}

func (m *Model) UpdateRow(index int, rowItem *RowItem) {
	m.rows[index] = NewRow(rowItem)
}

func (m *Model) NextRow() tea.Cmd {
	if m.selectedIndex < len(m.rows)-1 {
		m.selectedIndex = m.selectedIndex + 1
	}

	return rowSelected(m)
}

func (m *Model) PrevRow() tea.Cmd {
	if m.selectedIndex > 0 {
		m.selectedIndex = m.selectedIndex - 1
	}

	return rowSelected(m)
}

func (m *Model) SetRows(rows ...Row) {
	m.rows = rows
}
