package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func remote(c *cli.Context) error {
	remotes, err := git.Remote()
	if err != nil {
		return fmt.Errorf("failed to list remotes: %v", err)
	}

	fmt.Println(remotes)
	return nil
}
