package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/config"
	"github.com/ammar-y62/git-cli/constants"
	"github.com/urfave/cli/v2"
)

/*
 * This function is a command handler for setting the peer URL configuration.
 * Returns an error if any.
 */
func setConfig(c *cli.Context) error {
	//Extract the peer URL from the command-line flags
	peerURL := c.String("peer-url")
	//Need to check if a peerURL is set, if not return an error
	if peerURL == "" {
		return fmt.Errorf("peer URL is required")
	}
	//Create a configuration object with the provided peer URL
	conf := config.Config{
		PeerURL: peerURL,
	}
	//Save the configuration to the config file
	if err := config.SaveConfig(conf, constants.ConfigFile); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}
	//Print a confirmation message indicating the successful setting of the peer URL
	fmt.Printf("Peer URL set to: %s\n", peerURL)
	return nil
}
