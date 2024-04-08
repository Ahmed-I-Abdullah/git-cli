package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

/*
 * This funciton handles pulling changes from a remote repository via gRPC.
 * Returns any errors, if any
 */
func pullViaGRPC(c *cli.Context) error {
	//Establish a connection to the gRPC server
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}

	defer client.Close()

	//Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	repoName := filepath.Base(currentDir)

	//Create a gRPC client for interacting with the repository service
	grpcClient := pb.NewRepositoryClient(client.GetConn())

	//Get the context with timeout
	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	//Call the Pull funciton for the gRPC client for the desired repository
	response, err := grpcClient.Pull(ctx, &pb.RepoPullRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("Failed to get repository url via GRPC: %v", err)
	}

	//Pull changes from the remote repository
	err = git.Pull(git.PullOptions{
		URL: response.RepoAddress,
		Options: git.Options{
			Verbose: true,
		},
	})

	//If the pull execution failed, respond with an error message
	if err != nil {
		return fmt.Errorf("Failed to pull repository: %v", err)
	}

	return nil
}
