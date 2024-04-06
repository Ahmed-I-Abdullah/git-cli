package commands

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	"github.com/ammar-y62/git-cli/git"
	"github.com/ammar-y62/git-cli/grpc"
	"github.com/urfave/cli/v2"
	google_grpc "google.golang.org/grpc"
)

func pushViaGRPC(c *cli.Context) error {
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}

	defer client.Close()

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	repoName := filepath.Base(currentDir)

	grpcClient := pb.NewRepositoryClient(client.GetConn())

	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	response, err := grpcClient.GetLeaderUrl(ctx, &pb.LeaderUrlRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("Failed to get leader url via GRPC: %v", err)
	}

	err = git.Push(git.PushOptions{
		Remote:  response.GitRepoAddress,
		PushAll: true,
		Options: git.Options{
			Verbose: true,
		},
	})

	if err != nil {
		return fmt.Errorf("Failed to push repository to leader: %v", err)
	}

	grpcCtx, grpcCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer grpcCancel()
	leaderConn, err := google_grpc.DialContext(grpcCtx, response.GrpcAddress, google_grpc.WithInsecure(), google_grpc.WithBlock())

	if err != nil {
		return fmt.Errorf("Failed to connect to leader with gRPC: %v", err)
	}
	defer leaderConn.Close()

	leaderGrpcClient := pb.NewRepositoryClient(leaderConn)

	notifyResponse, err := leaderGrpcClient.NotifyPushCompletion(ctx, &pb.NotifyPushCompletionRequest{Name: repoName})

	if err != nil {
		return fmt.Errorf("Failed to notify leader about repository push: %v", err)
	}

	fmt.Printf("Leader notify response: %s", notifyResponse.Message)

	return nil
}
