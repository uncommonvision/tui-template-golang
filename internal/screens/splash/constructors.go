package splash

import (
	"time"
	"tui-template/internal/styles"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

func New(value string, duration time.Duration) tea.Model {
	text := textinput.New()
	text.Prompt = ""
	text.SetValue(value)
	text.TextStyle = lg.NewStyle().Foreground(styles.Foreground)

	return Model{
		duration: duration,
		spinner:  spinner.New(spinner.WithSpinner(spinner.Points)),
		text:     text,
	}
}
