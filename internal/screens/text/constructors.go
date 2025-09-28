package text

import (
	"github.com/charmbracelet/lipgloss"

	"tui-template/internal/components/keymapview"
	"tui-template/internal/styles"
)

type Options struct {
	Style      *lipgloss.Style
	KeyMapView keymapview.Model
}

func New(value string, options Options) Model {
	if options.Style == nil {
		options.Style = &styles.BorderNone
	}

	return Model{
		value:      value,
		style:      *options.Style,
		keyMapView: options.KeyMapView,
	}
}
