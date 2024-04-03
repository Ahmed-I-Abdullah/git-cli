package commands

import (
	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func diff(c *cli.Context) error {
	err := git.Diff()
	if err != nil {
		return err
	}

	return nil
}
