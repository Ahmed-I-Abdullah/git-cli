package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ammar-y62/git-cli/constants"
)

type Config struct {
	PeerURL string `json:"peer_url"`
}

func LoadConfig(configFile string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filepath.Join(GetConfigDir(), configFile))
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

func SaveConfig(config Config, configFile string) error {
	configDir := GetConfigDir()
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %v", err)
		}
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(configDir, configFile), data, 0644)
}

func GetConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, constants.ConfigDir)
}
