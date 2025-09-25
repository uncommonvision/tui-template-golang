package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long:  `Display the current version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TUI Template v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
