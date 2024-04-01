package commands

import (
	"fmt"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
)

func checkoutViaGRPC(c *cli.Context) error {
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
		return fmt.Errorf("a branch name is required for checkout")
	}

	_, err = grpcClient.Checkout(ctx, &pb.RepoCheckoutRequest{BranchName: branchName})
	if err != nil {
		return fmt.Errorf("failed to perform checkout via GRPC: %v", err)
	}

	err = git.Checkout(git.CheckoutOptions{Branch: branchName})
	if err != nil {
		return fmt.Errorf("failed to checkout branch '%s': %v", branchName, err)
	}

	fmt.Printf("Switched to branch '%s'.\n", branchName)
	return nil
}