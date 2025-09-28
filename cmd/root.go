package cmd

import (
	"fmt"
	"os"

	"tui-template/internal/screens/form"
	"tui-template/internal/screens/home"
	"tui-template/internal/screens/list"
	"tui-template/internal/screens/router"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mytui",
	Short: "A TUI starter template using Bubble Tea",
	Long: `A TUI starter template built with Bubble Tea from Charm. This template provides examples and patterns for building
interactive terminal applications in Go.

Running 'mytui' without arguments will launch the main TUI application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := router.New(
			map[string]tea.Model{
				"home": home.New(),
				"form": form.New(),
				"list": list.New(),
			},
			router.Options{
				ReadyView: "home",
			},
		)

		p := tea.NewProgram(r, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug mode")
	rootCmd.PersistentFlags().String("config", "", "Config file path")
	rootCmd.PersistentFlags().String("theme", "default", "Color theme (default, dark, light)")
}
