package commands

import (
	"context"
	"fmt"
	"net"
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

	fmt.Printf("Leader response is %v", response)

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

	fmt.Printf("Leader grpc address: -%s-", response.GrpcAddress)

	currentConnAddress, err := resolveAddress(client.GetConn().Target())
	if err != nil {
		return fmt.Errorf("Failed to resolve current connection address: %v", err)
	}

	resLeaderGrpcAddr, err := resolveAddress(response.GrpcAddress)
	if err != nil {
		return fmt.Errorf("Failed to resolve leader gRPC address: %v", err)
	}

	var leaderConn *google_grpc.ClientConn

	if currentConnAddress == resLeaderGrpcAddr {
		leaderConn = client.GetConn()
	} else {
		grpcCtx, grpcCancel := context.WithTimeout(context.Background(), time.Second*5)
		defer grpcCancel()
		leaderConn, err = google_grpc.DialContext(grpcCtx, resLeaderGrpcAddr, google_grpc.WithInsecure(), google_grpc.WithBlock())

		if err != nil {
			return fmt.Errorf("Failed to connect to leader with gRPC: %v", err)
		}
		defer leaderConn.Close()
	}

	leaderGrpcClient := pb.NewRepositoryClient(leaderConn)

	acquirLockResponse, err := leaderGrpcClient.AcquireLock(ctx, &pb.AcquireLockRequest{
		RepoName: repoName,
	})

	if err != nil {
		return fmt.Errorf("Failed to acquire lock from leader: %v", err)
	}

	if !acquirLockResponse.Ok {
		return fmt.Errorf("Failed to acquire reposiotry push as another push is in progress. Please try again later")
	}

	notifyResponse, err := leaderGrpcClient.NotifyPushCompletion(ctx, &pb.NotifyPushCompletionRequest{Name: repoName})

	if err != nil {
		return fmt.Errorf("Failed to notify leader about repository push: %v", err)
	}

	fmt.Printf("Leader notify response: %s", notifyResponse.Message)
	return nil
}

func resolveAddress(address string) (string, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", err
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}

	if len(ips) == 0 {
		return "", fmt.Errorf("No IPs found for the hostname: %s", host)
	}

	for _, ip := range ips {
		if ip4 := ip.To4(); ip4 != nil {
			return net.JoinHostPort(ip4.String(), port), nil
		}
	}

	return net.JoinHostPort(ips[0].String(), port), nil
}
