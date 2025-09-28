package dynalist

import (
	lg "github.com/charmbracelet/lipgloss"
)

type MetaInfo int

type RowItem struct {
	Annotation  string
	Title       string
	Description string
	Meta        map[MetaInfo]string
}

type Row struct {
	Style lg.Style

	annotation  string
	title       string
	description string
	meta        map[MetaInfo]string
}
