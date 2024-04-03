package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func branch(c *cli.Context) error {
	branchName := c.Args().First()

	if branchName == "" {
		branchOutput, err := git.Branch(git.BranchOptions{List: true})
		if err != nil {
			return fmt.Errorf("failed to list branches: %v", err)
		}
		fmt.Println(branchOutput)
	} else {
		branchOutput, err := git.Branch(git.BranchOptions{Create: true, Name: branchName})
		if err != nil {
			return fmt.Errorf("failed to create branch '%s': %v", branchName, err)
		}
		fmt.Println(branchOutput)
	}

	return nil
}
