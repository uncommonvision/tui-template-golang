package list

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"tui-template/internal/styles"
)

type Item struct {
	title, desc string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

type Model struct {
	list list.Model
}

func NewModel() Model {
	items := []list.Item{
		Item{title: "🍎 Apple", desc: "A crisp red fruit"},
		Item{title: "🍌 Banana", desc: "A yellow curved fruit"},
		Item{title: "🍊 Orange", desc: "A citrus fruit"},
		Item{title: "🍇 Grapes", desc: "Small round fruits in bunches"},
		Item{title: "🍓 Strawberry", desc: "A red berry with seeds"},
		Item{title: "🥝 Kiwi", desc: "A fuzzy brown fruit with green inside"},
		Item{title: "🍑 Cherry", desc: "Small red stone fruit"},
		Item{title: "🍒 Cherries", desc: "Paired red stone fruits"},
		Item{title: "🥭 Mango", desc: "A tropical stone fruit"},
		Item{title: "🍍 Pineapple", desc: "A tropical fruit with spiky exterior"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Fruit List Example"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = styles.TitleStyle
	l.Styles.PaginationStyle = styles.HelpStyle
	l.Styles.HelpStyle = styles.HelpStyle

	return Model{list: l}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 3)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			item := m.list.SelectedItem().(Item)
			return m, tea.Sequence(
				tea.Printf("You selected: %s", item.title),
				tea.Quit,
			)
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	help := styles.HelpStyle.Render("Press enter to select, q to quit, / to filter")
	return fmt.Sprintf("%s\n%s", m.list.View(), help)
}
