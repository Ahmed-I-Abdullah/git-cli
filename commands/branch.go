package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func branchViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	branchName := c.Args().First()

	_, err = grpcClient.Branch(ctx, &pb.RepoBranchRequest{Name: branchName})
	if err != nil {
		return fmt.Errorf("failed to perform branch operation via GRPC: %v", err)
	}

	if branchName == "" {
		branches, err := git.Branch(git.BranchOptions{List: true, Options: git.Options{Verbose: true}})
		if err != nil {
			return fmt.Errorf("failed to list branches: %v", err)
		}
		fmt.Println(branches)
	} else {
		_, err := git.Branch(git.BranchOptions{Create: true, Name: branchName, Options: git.Options{Verbose: true}})
		if err != nil {
			return fmt.Errorf("failed to create branch '%s': %v", branchName, err)
		}
		fmt.Printf("Branch '%s' created successfully.\n", branchName)
	}

	return nil
}
