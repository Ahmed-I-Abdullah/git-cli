package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func initViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}

	defer client.Close()

	repoName := c.Args().First()

	if repoName == "" {
		return fmt.Errorf("you must provide a repository name for initialization")
	}

	fmt.Println("Initializing repository with name: %s", repoName)

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	response, err := grpcClient.Init(ctx, &pb.RepoInitRequest{Name: repoName, FromCli: true})
	if err != nil {
		return fmt.Errorf("Failed to initialize repository via GRPC: %v", err)
	}

	fmt.Printf("Repository initialized successfully. Server Response: %s\n", response.Message)
	return nil
}
