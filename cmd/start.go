package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"tui-template/internal/models"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the main TUI application",
	Long:  `Launch the main TUI application with navigation between Home, Settings, and Help screens.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := models.NewMainModel()
		p := tea.NewProgram(model, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
