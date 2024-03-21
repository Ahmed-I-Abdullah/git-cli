package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/config"
	"github.com/ammar-y62/git-cli/constants"
	"github.com/urfave/cli/v2"
)

func setConfig(c *cli.Context) error {
	peerURL := c.String("peer-url")
	if peerURL == "" {
		return fmt.Errorf("peer URL is required")
	}

	conf := config.Config{
		PeerURL: peerURL,
	}

	if err := config.SaveConfig(conf, constants.ConfigFile); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	fmt.Printf("Peer URL set to: %s\n", peerURL)
	return nil
}
