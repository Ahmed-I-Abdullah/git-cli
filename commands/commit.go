package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func commit(c *cli.Context) error {
	commitMessage := c.String("message")
	if commitMessage == "" {
		return fmt.Errorf("commit message is required")
	}

	err := git.Commit(git.CommitOptions{
		Message: commitMessage,
	})

	if err != nil {
		return fmt.Errorf("failed to commit files: %v", err)
	}

	fmt.Println("Files committed successfully.")
	return nil
}
