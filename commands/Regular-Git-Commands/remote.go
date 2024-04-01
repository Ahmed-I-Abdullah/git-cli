package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func remoteViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	_, err = grpcClient.Remote(ctx, &pb.RepoRemoteRequest{})
	if err != nil {
		return err
	}

	remotes, err := git.Remote(git.RemoteOptions{})
	if err != nil {
		return err
	}

	fmt.Println(remotes)
	return nil
}
