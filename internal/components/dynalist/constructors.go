package dynalist

import keymapview "tui-template/internal/components/keymapview"

func New(width, height int) Model {
	return Model{
		AnnotationStyle:         DefaultAnnotationStyle(),
		AnnotationSelectedStyle: DefaultAnnotationSelectedStyle(),
		ListStyle:               DefaultListStyle(),
		RowStyle:                DefaultRowStyle(),
		RowSelectedStyle:        DefaultRowSelectedStyle(),
		width:                   width,
		height:                  height,
		selectedIndex:           0,
		keyMapView:              keymapview.New(DefaultKeyMap),
	}
}

func NewRow(item *RowItem) Row {
	return Row{
		annotation:  item.Annotation,
		title:       item.Title,
		description: item.Description,
		meta:        item.Meta,
	}
}
