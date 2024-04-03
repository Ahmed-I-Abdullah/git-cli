package commands

import (
	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func reset(c *cli.Context) error {
	commit := c.Args().First()

	err := git.Reset(git.ResetOptions{Commit: commit})
	if err != nil {
		return err
	}

	return nil
}
