package commands

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	{
		Name:   "init",
		Usage:  "Initialize a new Git repository via GRPC",
		Action: initViaGRPC,
	},
	{
		Name:   "pull",
		Usage:  "Pull changes from a Git repository via GRPC",
		Action: pullViaGRPC,
	},
	{
		Name:   "clone",
		Usage:  "Clone a Git repository via GRPC",
		Action: cloneViaGRPC,
	},
	{
		Name:   "push",
		Usage:  "Push Git repository via GRPC",
		Action: pushViaGRPC,
	},
	{
		Name:   "add",
		Usage:  "Add file contents to the index",
		Action: addViaGRPC,
	},
	{
		Name:   "commit",
		Usage:  "Record changes to the repository",
		Action: commitViaGRPC,
	},
	{
		Name:   "config",
		Usage:  "Set configuration options",
		Action: setConfig,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "peer-url",
				Usage: "Set the URL of the GRPC peer",
			},
		},
	},
}
