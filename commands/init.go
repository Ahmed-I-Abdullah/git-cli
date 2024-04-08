package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

/*
 * This funciton handles initializing a new repository via gRPC.
 * Returns an output of the server response.
 */
func initViaGRPC(c *cli.Context) error {
	//Establish a connection to the gRPC server
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}

	defer client.Close()

	//Get the repository name from the command-line arguments
	repoName := c.Args().First()

	//If repository name is empty, return an error message
	if repoName == "" {
		return fmt.Errorf("you must provide a repository name for initialization")
	}

	//Print that the repository is being initialized
	fmt.Println("Initializing repository with name: %s", repoName)

	//Create a gRPC client for interacting with the repository service
	grpcClient := pb.NewRepositoryClient(client.GetConn())

	//Get the context with timeout
	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	//Initialize repository via gRPC client
	response, err := grpcClient.Init(ctx, &pb.RepoInitRequest{Name: repoName, FromCli: true})
	if err != nil {
		return fmt.Errorf("Failed to initialize repository via GRPC: %v", err)
	}

	//Respond with the server response indicating that the repository has been initialized
	fmt.Printf("Repository initialized successfully. Server Response: %s\n", response.Message)
	return nil
}
