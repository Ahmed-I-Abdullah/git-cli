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

/*
 * This function is a command handler function for pushing changes to a remote repository via gRPC.
 * Returns an error if there is any
 */
func pushViaGRPC(c *cli.Context) error {
	//Establish a connection to the gRPC server
	client, err := grpc.GetConnection()

	if err != nil {
		return err
	}

	defer client.Close()

	//Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	repoName := filepath.Base(currentDir)

	//Create a gRPC client for interacting with the repository service
	grpcClient := pb.NewRepositoryClient(client.GetConn())

	//Get the context with timeout
	ctx, cancel := client.GetContextWithTimeout()
	defer cancel()

	//Get the leader URL for the repository
	response, err := grpcClient.GetLeaderUrl(ctx, &pb.LeaderUrlRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("Failed to get leader url via GRPC: %v", err)
	}

	//Print with the leader response and the leader's gRPC address
	fmt.Printf("Leader response is %v", response)

	fmt.Printf("Leader grpc address: -%s-", response.GrpcAddress)

	//Resolve current addresses
	currentConnAddress, err := resolveAddress(client.GetConn().Target())
	if err != nil {
		return fmt.Errorf("Failed to resolve current connection address: %v", err)
	}

	//Resolve leader addresses
	resLeaderGrpcAddr, err := resolveAddress(response.GrpcAddress)
	if err != nil {
		return fmt.Errorf("Failed to resolve leader gRPC address: %v", err)
	}

	var leaderConn *google_grpc.ClientConn

	//If the current address is the leader, get the client connection straight away
	//If not, then connect to the leader
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

	//Once connected to the leader, attempt to acquire lock
	fmt.Printf("\nAttempting to acquire lock for repository %s", repoName)

	acquirLockResponse, err := leaderGrpcClient.AcquireLock(ctx, &pb.AcquireLockRequest{
		RepoName: repoName,
	})

	//If an error occurs when trying to acquire the lock, respond with an error
	if err != nil {
		fmt.Errorf("\nFailed to aquire lock from leader for repo %s. Error: %v", repoName, err)
		return fmt.Errorf("Failed to acquire lock from leader: %v", err)
	}

	//If a lock is taken, respond with a message saying that another push is in progress
	if !acquirLockResponse.Ok {
		fmt.Printf("\nFailed to acquired lock from leader for reposiotry %s. Another push is in progress.", repoName)
		return fmt.Errorf("Failed to acquire reposiotry push as another push is in progress. Please try again later")
	}
	//Once a lock is acquired, push changes to the remote repository
	fmt.Printf("\nSuccessfully acquired lock from leader for reposiotry %s", repoName)

	err = git.Push(git.PushOptions{
		Remote:  response.GitRepoAddress,
		PushAll: true,
		Options: git.Options{
			Verbose: true,
		},
	})

	//If the push failed, respond with an error message
	if err != nil {
		return fmt.Errorf("Failed to push repository to leader: %v", err)
	}

	//If a push succeeds, notify the leader about the completion of the push operation
	notifyResponse, err := leaderGrpcClient.NotifyPushCompletion(ctx, &pb.NotifyPushCompletionRequest{Name: repoName})

	//If it failed to notify the leader, respond with an error message
	if err != nil {
		return fmt.Errorf("Failed to notify leader about repository push: %v", err)
	}

	//If a notify succeeds, respond with a message that the leader has been notified
	fmt.Printf("\nLeader notify response: %s", notifyResponse.Message)
	return nil
}

/*
 * This function resolves the address to an IP address.
 * Returns IP address from the host:port
 */
func resolveAddress(address string) (string, error) {
	//Split the address into host and port
	host, port, err := net.SplitHostPort(address)
	//if an error occurs, return an empty string with the error
	if err != nil {
		return "", err
	}

	//Lookup IP addresses associated with the given hostname
	ips, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}

	//If no IP address was found, return an error
	if len(ips) == 0 {
		return "", fmt.Errorf("No IPs found for the hostname: %s", host)
	}

	//Iterate through the list of IP addresses
	for _, ip := range ips {
		//Check if the IP address is IPv4
		if ip4 := ip.To4(); ip4 != nil {
			//return the joined IPv4 address with the port
			return net.JoinHostPort(ip4.String(), port), nil
		}
	}

	//If no IPv4 addresses were found, return the first IP address with the port
	return net.JoinHostPort(ips[0].String(), port), nil
}
