package commands

import (
	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func resetViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	commit := c.Args().First()

	_, err = grpcClient.Reset(ctx, &pb.RepoResetRequest{Commit: commit})
	if err != nil {
		return err
	}

	err = git.Reset(git.ResetOptions{Commit: commit})
	if err != nil {
		return err
	}

	return nil
}
