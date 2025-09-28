package router

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"tui-template/internal/components/keymapview"
	"tui-template/internal/screens/splash"
	"tui-template/internal/screens/text"
)

var DefaultSplashView = "splash"
var DefaultSplashText = "loading"
var DefaultReadyView = "home"

type KeyMap struct {
	Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

type Options struct {
	SplashView string
	SplashText string
	ReadyView  string
}

func New(views map[string]tea.Model, options Options) Model {

	if len(options.SplashView) == 0 {
		options.SplashView = DefaultSplashView
	}

	if len(options.SplashText) == 0 {
		options.SplashText = DefaultSplashText
	}

	if len(options.ReadyView) == 0 {
		options.ReadyView = DefaultReadyView
	}

	if _, exists := views[options.SplashView]; !exists {
		views[options.SplashView] = splash.New(DefaultSplashText, time.Second*2)
	}

	if _, exists := views[options.ReadyView]; !exists {
		views[options.ReadyView] = text.New("Error: Missing Ready View", text.Options{
			KeyMapView: keymapview.New(DefaultKeyMap),
		})
	}

	return Model{
		views:      views,
		section:    options.SplashView,
		splashView: options.SplashView,
		readyView:  options.ReadyView,
	}
}
