package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ammar-y62/git-cli/commands"
	"github.com/urfave/cli/v2"
)

var (
	Version  = "0"
	CommitId = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "Git Peer Cli Wrapper"
	app.Usage = "A CLI wrapper for a p2p git implementation"

	app.Version = fmt.Sprintf("%s - %s", Version, CommitId)

	app.Commands = commands.Commands

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
