package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"github.com/libp2p/go-libp2p"
	libp2pgrpc "github.com/paralin/go-libp2p-grpc"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "git-grpc-wrapper",
		Usage: "A CLI wrapper for Git commands with GRPC support for specific tasks",
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

func executeGitCommand(c *cli.Context) error {
	// Use os/exec to run git commands, args are available in c.Args().Slice()
	fmt.Println("Executing Git command:", c.Args().Slice())
	// Implementation here...
	return nil
}

func executeViaGRPC(c *cli.Context) error {
	ctx := context.Background()

	// Setup libp2p host
	host, err := libp2p.New(ctx)
	if err != nil {
		return err
	}

	// Setup GRPC over libp2p
	grpcHost := libp2pgrpc.NewGrpcHost(host)

	// Assuming you have a service definition for your GRPC services
	// Register your GRPC services here with grpcHost

	fmt.Println("Executing via GRPC:", c.Args().Slice())
	// Implementation here...
	return nil
}
