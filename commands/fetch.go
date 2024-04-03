package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func fetch(c *cli.Context) error {
	remoteName := c.String("remote")

	err := git.Fetch(git.FetchOptions{
		Remote: remoteName,
	})

	if err != nil {
		return fmt.Errorf("failed to fetch from remote '%s': %v", remoteName, err)
	}

	fmt.Printf("Fetched from remote '%s' successfully.\n", remoteName)
	return nil
}
