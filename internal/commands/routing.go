package commands

import tea "github.com/charmbracelet/bubbletea"

type RouteToViewMsg struct {
	Section string
}

func routeToView(section string) tea.Cmd {
	return func() tea.Msg {
		return RouteToViewMsg{
			Section: section,
		}
	}
}

func RouteToForm() tea.Cmd {
	return routeToView("form")
}

func RouteToHome() tea.Cmd {
	return routeToView("home")
}

func RouteToList() tea.Cmd {
	return routeToView("list")
}
