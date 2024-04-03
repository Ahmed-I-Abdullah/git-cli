package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func merge(c *cli.Context) error {
	branchName := c.Args().First()
	if branchName == "" {
		return fmt.Errorf("a branch name is required for merge")
	}

	err := git.Merge(git.MergeOptions{Branch: branchName})
	if err != nil {
		return fmt.Errorf("failed to merge branch '%s': %v", branchName, err)
	}

	fmt.Printf("Merged branch '%s' successfully.\n", branchName)
	return nil
}
