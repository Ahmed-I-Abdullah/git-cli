package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	// Assuming you have a GRPC package for client initialization
	// "mygrpcpackage"
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
				Name:    "grpc",
				Aliases: []string{"rpc"},
				Usage:   "Execute commands via GRPC",
				Action:  executeViaGRPC,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Global GRPC client connection
var grpcClientConn *grpc.ClientConn

func setupGRPCClient(c *cli.Context) error {
	peerURL := c.String("peer-url")
	if peerURL == "" {
		return nil // No GRPC setup required if no peer URL provided
	}

	// Setup GRPC client connection
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

	// Execute Git command
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git command error: %v, output: %s", err, output)
	}

	fmt.Printf("Git command output:\n%s\n", output)
	return nil
}

func executeViaGRPC(c *cli.Context) error {
	args := c.Args().Slice()
	if len(args) < 1 {
		return fmt.Errorf("no command provided for GRPC execution")
	}

	if grpcClientConn == nil {
		return fmt.Errorf("GRPC client is not setup. Please provide a peer URL.")
	}

	// Assuming "MyGRPCClient" is initialized using "grpcClientConn" somewhere after connection setup
	// myGRPCClient := mygrpcpackage.NewMyGRPCClient(grpcClientConn)
	// Perform your GRPC request...

	fmt.Println("GRPC command simulated:", strings.Join(args, " "))
	// Simulate GRPC operation
	return nil
}

func mainCleanup() {
	if grpcClientConn != nil {
		grpcClientConn.Close()
	}
}
