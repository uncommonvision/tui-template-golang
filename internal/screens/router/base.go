package router

import (
	tea "github.com/charmbracelet/bubbletea"

	"tui-template/internal/commands"
	"tui-template/internal/screens/splash"
)

type Model struct {
	width   int
	height  int
	views   map[string]tea.Model
	section string

	splashView string
	readyView  string
}

func (m Model) Init() tea.Cmd {
	return m.views[m.splashView].Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		for key, view := range m.views {
			var cmd tea.Cmd
			m.views[key], cmd = view.Update(msg)
			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)
	case splash.SplashCompleteMsg:
		m.section = m.readyView
		cmds = append(cmds, m.views[m.section].Init())
	case commands.RouteToViewMsg:
		m.section = msg.Section
	}

	var cmd tea.Cmd
	m.views[m.section], cmd = m.views[m.section].Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.views[m.section].View()
}
