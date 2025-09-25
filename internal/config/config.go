package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Theme    string `json:"theme"`
	Debug    bool   `json:"debug"`
	AutoSave bool   `json:"auto_save"`
	LogLevel string `json:"log_level"`
}

func DefaultConfig() Config {
	return Config{
		Theme:    "default",
		Debug:    false,
		AutoSave: true,
		LogLevel: "info",
	}
}

func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "mytui")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(configDir, "config.json"), nil
}

func Load() (Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return DefaultConfig(), err
	}

	data, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		return DefaultConfig(), nil
	}
	if err != nil {
		return DefaultConfig(), err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return DefaultConfig(), err
	}

	return config, nil
}

func (c Config) Save() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func (c Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	return string(data)
}

func Init() error {
	config := DefaultConfig()
	if err := config.Save(); err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}

	configPath, _ := GetConfigPath()
	fmt.Printf("Config file created at: %s\n", configPath)
	return nil
}
