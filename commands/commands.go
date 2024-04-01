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
		Name:   "status",
		Usage:  "Show the working tree status",
		Action: statusViaGRPC,
	},
	{
		Name:   "branch",
		Usage:  "List, create, or delete branches",
		Action: branchViaGRPC,
	},
	{
		Name:   "checkout",
		Usage:  "Switch branches or restore working tree files",
		Action: checkoutViaGRPC,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "branch",
				Aliases:  []string{"b"},
				Usage:    "Branch name to checkout",
				Required: true,
			},
		},
	},
	{
		Name:   "diff",
		Usage:  "Show the git diff",
		Action: diffViaGRPC,
	},
	{
		Name:   "fetch",
		Usage:  "Fetch branches and/or tags from another repository, and update the remote tracking branches",
		Action: fetchViaGRPC,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "remote",
				Aliases: []string{"r"},
				Usage:   "The name of the remote repository to fetch from",
			},
		},
	},
	{
		Name:   "log",
		Usage:  "Show commit logs",
		Action: logViaGRPC,
	},
	{
		Name:   "merge",
		Usage:  "Join two or more development histories together",
		Action: mergeViaGRPC,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "branch",
				Aliases:  []string{"b"},
				Usage:    "Branch name to merge into the current branch",
				Required: true,
			},
		},
	},
	{
		Name:   "rebase",
		Usage:  "Reapply commits on top of another base tip",
		Action: rebaseViaGRPC,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "branch",
				Aliases:  []string{"b"},
				Usage:    "Branch name to rebase onto the current branch",
				Required: true,
			},
		},
	},
	{
		Name:   "remote",
		Usage:  "Manage set of tracked repositories",
		Action: remoteViaGRPC,
	},
	{
		Name:   "reset",
		Usage:  "Reset current HEAD to the specified state",
		Action: resetViaGRPC,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "commit",
				Usage:    "Commit hash to reset to",
				Required: false, // Make required false if you want to allow reset without specifying a commit, which would default to HEAD
			},
		},
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
