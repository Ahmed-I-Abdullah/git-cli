package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func statusViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	_, err = grpcClient.Status(ctx, &pb.RepoStatusRequest{})
	if err != nil {
		return fmt.Errorf("failed to send status request via GRPC: %v", err)
	}

	status, err := git.Status(git.StatusOptions{})
	if err != nil {
		return fmt.Errorf("failed to get status: %v", err)
	}

	fmt.Println(status)
	return nil
}
