package form

import (
	"tui-template/internal/components/keymapview"

	"github.com/charmbracelet/bubbles/textinput"
)

func New() Model {
	inputs := make([]textinput.Model, 3)

	inputs[0] = textinput.New()
	inputs[0].Placeholder = "Enter your name"
	inputs[0].Focus()

	inputs[1] = textinput.New()
	inputs[1].Placeholder = "Enter your email"

	inputs[2] = textinput.New()
	inputs[2].Placeholder = "Enter your company"

	return Model{
		inputs:     inputs,
		focused:    0,
		keyMapView: keymapview.New(DefaultKeyMap),
	}
}
