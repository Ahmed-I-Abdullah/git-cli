package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func tagViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	tagName := c.Args().First()
	message := c.String("m")

	_, err = grpcClient.Tag(ctx, &pb.RepoTagRequest{Name: tagName, Message: message})
	if err != nil {
		return err
	}

	err = git.Tag(git.TagOptions{Name: tagName, Message: message})
	if err != nil {
		return err
	}

	fmt.Printf("Tag '%s' created successfully.\n", tagName)
	return nil
}
