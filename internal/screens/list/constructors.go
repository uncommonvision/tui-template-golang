package list

import (
	"tui-template/internal/components/keymapview"
	"tui-template/internal/styles"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

func New() Model {
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
	l.SetShowHelp(true)
	l.SetFilteringEnabled(true)

	l.KeyMap.Quit.SetEnabled(false)

	l.Help.Styles.ShortKey = styles.BorderNone
	l.Help.Styles.FullKey = styles.BorderNone

	l.Styles.Title = styles.TitleStyle
	l.Styles.PaginationStyle = styles.HelpStyle
	l.Styles.HelpStyle = styles.HelpStyle

	var DefaultKeyMap = KeyMap{
		Back: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("left", "back"),
		),
	}

	return Model{
		list:       l,
		keyMapView: keymapview.New(DefaultKeyMap),
	}
}
