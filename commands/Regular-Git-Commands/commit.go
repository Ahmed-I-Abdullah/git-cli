package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func commitViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	commitMessage := c.String("message")
	if commitMessage == "" {
		return fmt.Errorf("commit message is required")
	}

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	_, err = grpcClient.Commit(ctx, &pb.RepoCommitRequest{Message: commitMessage})
	if err != nil {
		return fmt.Errorf("failed to commit via GRPC: %v", err)
	}

	err = git.Commit(git.CommitOptions{
		Message: commitMessage,
		Options: git.Options{
			Verbose: true,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to commit files: %v", err)
	}

	fmt.Println("Files committed successfully.")
	return nil
}
