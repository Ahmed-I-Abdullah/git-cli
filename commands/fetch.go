package commands

import (
	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func fetchViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	_, err = grpcClient.Fetch(ctx, &pb.RepoFetchRequest{})
	if err != nil {
		return err
	}

	err = git.Fetch(git.FetchOptions{Remote: c.String("remote")})
	if err != nil {
		return err
	}

	return nil
}
