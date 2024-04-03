package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func status(c *cli.Context) error {
	status, err := git.Status()
	if err != nil {
		return fmt.Errorf("failed to get status: %v", err)
	}

	fmt.Println(status)
	return nil
}
