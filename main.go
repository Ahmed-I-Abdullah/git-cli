package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Starting the application...") // This line will confirm the app starts

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
	fmt.Println("Git working with command:", c.Args().Slice())
	return nil
}

func executeViaGRPC(c *cli.Context) error {
	// For testing purposes, let's assume the libp2p and GRPC setup is successful
	// and just print "GRPC working"
	fmt.Println("GRPC working with command:", c.Args().Slice())

	// The code below is commented out because we are not actually setting up
	// a GRPC connection in this test scenario
	/*
		ctx := context.Background()
		host, err := libp2p.New(ctx)
		if err != nil {
			return err
		}
		defer host.Close()
		grpcHost := libp2pgrpc.NewGrpcHost(host)
		// Register your GRPC services here with grpcHost
	*/

	return nil
}
