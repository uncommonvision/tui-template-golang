package keymapview

func New(km any) Model {
	return Model{
		Width:  0,
		Height: 0,
		keymap: km,
	}
}
