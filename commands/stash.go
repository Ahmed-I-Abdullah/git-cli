package commands

import (
	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func stash(c *cli.Context) error {
	err := git.Stash()
	if err != nil {
		return err
	}

	return nil
}
