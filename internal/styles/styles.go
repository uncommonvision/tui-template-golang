package styles

import "github.com/charmbracelet/lipgloss"

var (
	Primary   = lipgloss.Color("#7C3AED")
	Secondary = lipgloss.Color("#EC4899")
	Success   = lipgloss.Color("#10B981")
	Warning   = lipgloss.Color("#F59E0B")
	Error     = lipgloss.Color("#EF4444")
	Info      = lipgloss.Color("#3B82F6")

	Foreground = lipgloss.Color("#FFFFFF")
	Background = lipgloss.Color("#1F2937")
	Muted      = lipgloss.Color("#6B7280")
	Border     = lipgloss.Color("#374151")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Primary).
			MarginBottom(1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Foreground).
			Background(Primary).
			Padding(0, 1)

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Border).
			Padding(1, 2)

	ButtonStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Foreground).
			Background(Primary).
			Padding(0, 3).
			Margin(0, 1)

	ButtonActiveStyle = ButtonStyle.Copy().
				Background(Secondary)

	HelpStyle = lipgloss.NewStyle().
			Foreground(Muted).
			MarginTop(1)

	ErrorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Error)

	SuccessStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Success)

	InfoStyle = lipgloss.NewStyle().
			Foreground(Info)

	WarningStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Warning)

	ListItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	ListSelectedStyle = lipgloss.NewStyle().
				PaddingLeft(1).
				Foreground(Primary).
				Bold(true)

	StatusBarStyle = lipgloss.NewStyle().
			Background(Border).
			Foreground(Foreground).
			Padding(0, 1)
)
