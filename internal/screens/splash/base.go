package splash

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lg "github.com/charmbracelet/lipgloss"

	"tui-template/internal/styles"
)

type SplashCompleteMsg time.Time

type Model struct {
	height   int
	width    int
	text     textinput.Model
	spinner  spinner.Model
	duration time.Duration
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		tea.Every(m.duration, func(t time.Time) tea.Msg {
			return SplashCompleteMsg(t)
		}),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch {
		case msg.String() == "q":
			cmds = append(cmds, tea.Quit)
		}
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return styles.
		BorderNone.
		Width(m.width).
		Height(m.height).
		Align(lg.Center, lg.Center).
		Render(
			lipgloss.JoinHorizontal(
				lg.Center,
				m.text.View(),
				m.spinner.View(),
			),
		)
}
