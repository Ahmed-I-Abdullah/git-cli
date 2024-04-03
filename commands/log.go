package commands

import (
	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func log(c *cli.Context) error {
	err := git.Log()
	if err != nil {
		return err
	}

	return nil
}
