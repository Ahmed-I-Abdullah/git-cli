package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func rebase(c *cli.Context) error {
	branchName := c.Args().First()
	if branchName == "" {
		return fmt.Errorf("a branch name is required for rebase")
	}

	err := git.Rebase(git.RebaseOptions{Branch: branchName})
	if err != nil {
		return fmt.Errorf("failed to rebase branch '%s': %v", branchName, err)
	}

	fmt.Printf("Successfully rebased branch '%s'.\n", branchName)
	return nil
}
