package constants

import "time"

// TimeoutDuration specifies the duration for timeout in the CLI operations.
// It is currently set to 20 seconds.
const TimeoutDuration = 20 * time.Second

// ConfigDir represents the directory name for configuration files.
const ConfigDir = ".git-p2p-wrapper"

// ConfigFile represents the name of the configuration file.
const ConfigFile = "config.json"
