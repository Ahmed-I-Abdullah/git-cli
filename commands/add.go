package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func addViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	filePaths := c.Args().Slice()
	if len(filePaths) == 0 {
		return fmt.Errorf("you must provide at least one file path to add")
	}

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	_, err = grpcClient.Add(ctx, &pb.RepoAddRequest{FilePaths: filePaths})
	if err != nil {
		return fmt.Errorf("failed to add files via GRPC: %v", err)
	}

	err = git.Add(git.AddOptions{
		FilePaths: filePaths,
	})

	if err != nil {
		return fmt.Errorf("failed to add files: %v", err)
	}

	fmt.Println("Files added successfully.")
	return nil
}
