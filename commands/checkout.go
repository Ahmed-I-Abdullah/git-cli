package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func checkout(c *cli.Context) error {
	branchName := c.Args().First()
	if branchName == "" {
		return fmt.Errorf("a branch name is required for checkout")
	}

	err := git.Checkout(git.CheckoutOptions{Branch: branchName})
	if err != nil {
		return fmt.Errorf("failed to checkout branch '%s': %v", branchName, err)
	}

	fmt.Printf("Switched to branch '%s'.\n", branchName)
	return nil
}
