package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func cloneViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}
	defer client.Close()

	repoName := c.Args().First()
	if repoName == "" {
		return fmt.Errorf("You must provide a repository name for cloning")
	}

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	response, err := grpcClient.Pull(ctx, &pb.RepoPullRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("failed to get repository url via GRPC: %v", err)
	}

	if !response.Success {
		return fmt.Errorf("failed to clone repository: server returned unsuccessful response")
	}

	fmt.Printf("Repository URL retrieval successful. Repo Address: %s\n", response.RepoAddress)

	err = git.Clone(git.CloneOptions{
		URL: response.RepoAddress,
		Dir: fmt.Sprintf("./%s", repoName),
		Options: git.Options{
			Verbose: true,
		},
	})

	if err != nil {
		return fmt.Errorf("Failed clone repository: %v", err)
	}

	return nil
}
