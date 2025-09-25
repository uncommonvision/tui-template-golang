package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"tui-template/internal/styles"
)

type Screen int

const (
	HomeScreen Screen = iota
	SettingsScreen
	HelpScreen
)

type MainModel struct {
	currentScreen Screen
	width         int
	height        int
	ready         bool
}

func NewMainModel() MainModel {
	return MainModel{
		currentScreen: HomeScreen,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1", "h":
			m.currentScreen = HomeScreen
		case "2", "s":
			m.currentScreen = SettingsScreen
		case "3", "?":
			m.currentScreen = HelpScreen
		}
	}

	return m, nil
}

func (m MainModel) View() string {
	if !m.ready {
		return "Loading..."
	}

	var content string

	switch m.currentScreen {
	case HomeScreen:
		content = m.renderHome()
	case SettingsScreen:
		content = m.renderSettings()
	case HelpScreen:
		content = m.renderHelp()
	}

	header := m.renderHeader()
	footer := m.renderFooter()

	contentHeight := m.height - lipgloss.Height(header) - lipgloss.Height(footer) - 4
	contentBox := styles.BorderStyle.Copy().
		Width(m.width - 4).
		Height(contentHeight).
		Render(content)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		contentBox,
		footer,
	)
}

func (m MainModel) renderHeader() string {
	var title string
	switch m.currentScreen {
	case HomeScreen:
		title = "🏠 Home"
	case SettingsScreen:
		title = "⚙️  Settings"
	case HelpScreen:
		title = "❓ Help"
	}

	return styles.HeaderStyle.Copy().
		Width(m.width).
		Render(title)
}

func (m MainModel) renderFooter() string {
	help := []string{
		"1/h: Home",
		"2/s: Settings",
		"3/?: Help",
		"q: Quit",
	}

	helpText := strings.Join(help, " • ")
	return styles.HelpStyle.Copy().
		Width(m.width).
		Render(helpText)
}

func (m MainModel) renderHome() string {
	welcome := styles.TitleStyle.Render("Welcome to TUI Template!")

	features := []string{
		"✨ Built with Bubble Tea & Lipgloss",
		"🎨 Customizable themes and styling",
		"📱 Multiple screen navigation",
		"⌨️  Comprehensive keyboard shortcuts",
		"🧩 Reusable components",
		"📖 Example implementations",
	}

	var featureList []string
	for _, feature := range features {
		featureList = append(featureList, styles.ListItemStyle.Render(feature))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		welcome,
		"",
		"This template includes:",
		"",
		strings.Join(featureList, "\n"),
		"",
		styles.InfoStyle.Render("Navigate using the keyboard shortcuts shown below."),
	)

	return content
}

func (m MainModel) renderSettings() string {
	title := styles.TitleStyle.Render("Settings")

	settings := []string{
		fmt.Sprintf("Screen Size: %d x %d", m.width, m.height),
		"Theme: Default",
		"Debug Mode: Off",
		"Auto-save: On",
	}

	var settingsList []string
	for _, setting := range settings {
		settingsList = append(settingsList, styles.ListItemStyle.Render("• "+setting))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(settingsList, "\n"),
		"",
		styles.WarningStyle.Render("Settings are read-only in this template."),
	)

	return content
}

func (m MainModel) renderHelp() string {
	title := styles.TitleStyle.Render("Help & Keyboard Shortcuts")

	shortcuts := [][]string{
		{"Navigation", ""},
		{"1, h", "Go to Home screen"},
		{"2, s", "Go to Settings screen"},
		{"3, ?", "Show this Help screen"},
		{"", ""},
		{"Actions", ""},
		{"q, Ctrl+C", "Quit application"},
		{"", ""},
		{"Examples", ""},
		{"mytui examples list", "Run list example"},
		{"mytui examples form", "Run form example"},
		{"mytui examples dashboard", "Run dashboard example"},
	}

	var helpContent []string
	for _, shortcut := range shortcuts {
		if shortcut[0] == "" {
			helpContent = append(helpContent, "")
			continue
		}
		if shortcut[1] == "" {
			helpContent = append(helpContent, styles.HeaderStyle.Copy().
				Background(styles.Secondary).
				Render(shortcut[0]))
			continue
		}

		line := fmt.Sprintf("%-15s %s", shortcut[0], shortcut[1])
		helpContent = append(helpContent, styles.ListItemStyle.Render(line))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(helpContent, "\n"),
	)

	return content
}
