package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ammar-y62/git-cli/commands"
	"github.com/urfave/cli/v2"
)

// Version and CommitId store the version and commit ID of the application.
var (
	Version  = "0"
	CommitId = "0"
)

/*
 * This function acts as the main entry point to the Git Peer Cli Wrapper application.
 * It sets application information including name, usage, and version.
 * Adds commands to the application, handles command not found scenarios by proxying to git.
 * Runs the application, logging any errors encountered during execution
 */
func main() {
	//Create a new CLI application
	app := cli.NewApp()
	app.Name = "Git Peer Cli Wrapper"
	app.Usage = "A CLI wrapper for a p2p git implementation"

	// Set the version and commit ID of the application
	app.Version = fmt.Sprintf("%s - %s", Version, CommitId)

	//Add commands to the application
	app.Commands = commands.Commands

	//Handle command not found scenarios by proxying to git
	app.CommandNotFound = func(ctx *cli.Context, command string) {
		// proxy commands with no grpc handler to git
		args := os.Args[2:]
		commands.ProxyToGit(ctx, command, args...)
	}

	//Run the CLI application
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
