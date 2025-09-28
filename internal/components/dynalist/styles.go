package dynalist

import (
	lg "github.com/charmbracelet/lipgloss"
	"tui-template/internal/color"
)

func DefaultAnnotationStyle() lg.Style {
	return lg.NewStyle().PaddingRight(1).Foreground(color.DarkGray)
}

func DefaultAnnotationSelectedStyle() lg.Style {
	return lg.NewStyle().PaddingRight(1).Foreground(color.Lime)
}

func DefaultListStyle() lg.Style {
	return lg.NewStyle().Height(8).Width(80)
}

func DefaultRowStyle() lg.Style {
	return lg.NewStyle()
}

func DefaultRowSelectedStyle() lg.Style {
	return lg.NewStyle().Foreground(color.HotPink)
}
