package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"tui-template/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage application configuration",
	Long:  `Manage the application configuration file including initialization and viewing current settings.`,
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize default configuration file",
	Long:  `Create a default configuration file in the user's config directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Init(); err != nil {
			fmt.Printf("Error initializing config: %v\n", err)
			return
		}
		fmt.Println("Configuration file initialized successfully!")
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  `Display the current configuration settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		configPath, _ := config.GetConfigPath()
		fmt.Printf("Configuration file: %s\n\n", configPath)
		fmt.Println(cfg.String())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
}
