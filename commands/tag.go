package commands

import (
	"fmt"

	"github.com/ammar-y62/git-cli/git"
	"github.com/urfave/cli/v2"
)

func tag(c *cli.Context) error {
	tagName := c.Args().First()
	if tagName == "" {
		return fmt.Errorf("a tag name is required")
	}

	message := c.String("m")

	err := git.Tag(git.TagOptions{
		Name:    tagName,
		Message: message,
	})
	if err != nil {
		return fmt.Errorf("failed to create tag '%s': %v", tagName, err)
	}

	fmt.Printf("Tag '%s' created successfully.\n", tagName)
	return nil
}
