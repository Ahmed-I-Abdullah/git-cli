package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

/*
 * This funciton handles cloning repositories from a remote repository via gRPC.
 * Returns any errors, if any
 */
func cloneViaGRPC(c *cli.Context) error {
	//Establish a connection to the gRPC server
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}
	defer client.Close()

	//Get the repository name from the command-line arguments
	repoName := c.Args().First()
	if repoName == "" {
		return fmt.Errorf("You must provide a repository name for cloning")
	}

	//Create a gRPC client for interacting with the repository service
	grpcClient := pb.NewRepositoryClient(client.GetConn())

	//Get the context with timeout
	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	//Call the Pull funciton for the gRPC client for the newly initialized repository
	response, err := grpcClient.Pull(ctx, &pb.RepoPullRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("failed to get repository url via GRPC: %v", err)
	}

	//If the pull response is unsuccessful, respond with an error message
	if !response.Success {
		return fmt.Errorf("failed to clone repository: server returned unsuccessful response")
	}

	//Print the repository URL in the response
	fmt.Printf("Repository URL retrieval successful. Repo Address: %s\n", response.RepoAddress)

	//Clone repository using git command
	err = git.Clone(git.CloneOptions{
		URL: response.RepoAddress,
		Dir: fmt.Sprintf("./%s", repoName),
		Options: git.Options{
			Verbose: true,
		},
	})

	//If an error occurs when cloning repository, respond with an error message
	if err != nil {
		return fmt.Errorf("Failed clone repository: %v", err)
	}

	return nil
}
