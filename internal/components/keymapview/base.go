package keymapview

import (
	"fmt"
	"reflect"
	"strings"

	"tui-template/internal/styles"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Width  int
	Height int
	keymap any
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	helptext := []string{}

	if m.keymap == nil {
		return styles.HelpStyle.Width(m.Width).Render("")
	}

	v := reflect.ValueOf(m.keymap)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if binding, ok := field.Interface().(key.Binding); ok {
			keys := binding.Keys()
			help := binding.Help()

			if binding.Enabled() && len(keys) > 0 && help.Desc != "" {
				keyStr := strings.Join(keys, "/")
				helpStr := fmt.Sprintf("%s %s", keyStr, help.Desc)
				helptext = append(helptext, helpStr)
			}
		}
	}

	return styles.HelpStyle.
		Width(m.Width).
		Render(strings.Join(helptext, " • "))
}
