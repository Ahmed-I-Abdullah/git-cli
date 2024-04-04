package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func pullViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}

	defer client.Close()

	repoName := c.Args().First()
	if repoName == "" {
		return fmt.Errorf("you must provide a repository name for pulling changes")
	}

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	response, err := grpcClient.Pull(ctx, &pb.RepoPullRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("Failed to get repository url via GRPC: %v", err)
	}

	err = git.Pull(git.PullOptions{
		URL: response.RepoAddress,
		Options: git.Options{
			Verbose: true,
		},
	})

	if err != nil {
		return fmt.Errorf("Failed to pull repository: %v", err)
	}

	return nil
}
