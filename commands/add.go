package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func add(c *cli.Context) error {
	filePaths := c.Args().Slice()
	if len(filePaths) == 0 {
		return fmt.Errorf("you must provide at least one file path to add")
	}

	err := git.Add(git.AddOptions{
		FilePaths: filePaths,
	})

	if err != nil {
		return fmt.Errorf("failed to add files: %v", err)
	}

	fmt.Println("Files added successfully.")
	return nil
}
