package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"tui-template/examples/form"
	"tui-template/examples/list"
)

var examplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "Run example TUI applications",
	Long:  `Run various example TUI applications to see different patterns and components.`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Run the list navigation example",
	Long:  `Demonstrates a filterable list with navigation and selection.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := list.NewModel()
		p := tea.NewProgram(model, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running list example: %v\n", err)
			os.Exit(1)
		}
	},
}

var formCmd = &cobra.Command{
	Use:   "form",
	Short: "Run the form input example",
	Long:  `Demonstrates form input handling with multiple text fields and navigation.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := form.NewModel()
		p := tea.NewProgram(model, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running form example: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(examplesCmd)
	examplesCmd.AddCommand(listCmd)
	examplesCmd.AddCommand(formCmd)
}
