package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ammar-y62/git-cli/constants"
)

// This structure represents the structure of the configuration.
type Config struct {
	PeerURL string `json:"peer_url"`
}

/*
 * This function loads the configuration from the specified file.
 * Returns the loaded configuration and an error if any.
 */
func LoadConfig(configFile string) (Config, error) {
	var config Config
	//Attempt to read the data of the configuration file
	data, err := os.ReadFile(filepath.Join(GetConfigDir(), configFile))
	//If the configuration file cannot be read, return the configuration and the error
	if err != nil {
		return config, err
	}
	//If no error occured, unmarshal the json data, return the configuration and the unmarshalled json
	err = json.Unmarshal(data, &config)
	return config, err
}

/*
 * This function saves the configuration to the specified file.
 * Returns an error if the operation fails.
 */
func SaveConfig(config Config, configFile string) error {
	//Get the directory path where configuration files are stored
	configDir := GetConfigDir()
	//Check if the directory exists, if not, create it
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %v", err)
		}
	}
	//Marshal the configuration into a JSON format
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	//Write the JSON data to the specified configuration file
	return os.WriteFile(filepath.Join(configDir, configFile), data, 0644)
}

/*
 * This function gets the directory path where configuration files are stored.
 */
func GetConfigDir() string {
	//Get the user's home directory and join the filepath to the configeration directory
	home, _ := os.UserHomeDir()
	return filepath.Join(home, constants.ConfigDir)
}
