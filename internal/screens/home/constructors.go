package home

import (
	"tui-template/internal/components/keymapview"
)

func New() Model {
	return Model{
		keyMapView: keymapview.New(DefaultKeyMap),
	}
}
