package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/internal/api/pb"
	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting the application...")

	app := &cli.App{
		Name:  "git-grpc-wrapper",
		Usage: "A CLI wrapper for Git commands with GRPC support for specific tasks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "peer-url",
				Usage: "URL of the GRPC peer to connect to for remote operations",
			},
		},
		Before: setupGRPCClient,
		Commands: []*cli.Command{
			{
				Name:    "git",
				Aliases: []string{"g"},
				Usage:   "Execute Git commands directly",
				Action:  executeGitCommand,
			},
			{
				Name:   "init",
				Usage:  "Initialize a new Git repository via GRPC",
				Action: initViaGRPC,
			},
			// Placeholder for clone, push, and pull commands
			// Add similar structs for each operation as needed
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	mainCleanup()
}

// Global GRPC client connection
var grpcClientConn *grpc.ClientConn

func setupGRPCClient(c *cli.Context) error {
	peerURL := c.String("peer-url")
	if peerURL == "" {
		return nil // No GRPC setup required if no peer URL provided
	}

	var err error
	grpcClientConn, err = grpc.Dial(peerURL, grpc.WithInsecure()) // Consider secure options for production
	if err != nil {
		return fmt.Errorf("failed to connect to GRPC peer: %v", err)
	}

	return nil
}

func executeGitCommand(c *cli.Context) error {
	args := c.Args().Slice()
	if len(args) < 1 {
		return fmt.Errorf("no Git command provided")
	}

	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git command error: %v, output: %s", err, output)
	}

	fmt.Printf("Git command output:\n%s\n", output)
	return nil
}

func initViaGRPC(c *cli.Context) error {
	repoName := c.Args().First()
	if repoName == "" {
		return fmt.Errorf("you must provide a repository name for initialization")
	}

	if grpcClientConn == nil {
		return fmt.Errorf("GRPC client is not setup. Please provide a peer URL.")
	}

	grpcClient := pb.NewRepositoryClient(grpcClientConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := grpcClient.Init(ctx, &pb.RepoInitRequest{Name: repoName, FromCli: true})
	if err != nil {
		return fmt.Errorf("failed to initialize repository via GRPC: %v", err)
	}

	fmt.Printf("Repository initialized successfully. Server Response: %s\n", response.Message)
	return nil
}

// Placeholders for clone, push, and pull gRPC actions
// Implement these functions similar to initViaGRPC

func mainCleanup() {
	if grpcClientConn != nil {
		grpcClientConn.Close()
	}
}
