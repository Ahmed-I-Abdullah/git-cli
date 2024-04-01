package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func mergeViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	branchName := c.Args().First()
	if branchName == "" {
		return fmt.Errorf("a branch name is required for merge")
	}

	_, err = grpcClient.Merge(ctx, &pb.RepoMergeRequest{BranchName: branchName})
	if err != nil {
		return err
	}

	err = git.Merge(git.MergeOptions{Branch: branchName})
	if err != nil {
		return err
	}

	return nil
}
