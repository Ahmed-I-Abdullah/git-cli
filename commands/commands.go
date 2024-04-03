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
		Action: add,
	},
	{
		Name:   "commit",
		Usage:  "Record changes to the repository",
		Action: commit,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "message",
				Aliases:  []string{"m"},
				Usage:    "Commit message",
				Required: true,
			},
		},
	},
	{
		Name:   "branch",
		Usage:  "List, create, or delete branches",
		Action: branch,
	},
	{
		Name:   "status",
		Usage:  "Show the working tree status",
		Action: status,
	},
	{
		Name:   "checkout",
		Usage:  "Switch branches or restore working tree files",
		Action: checkout,
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
		Usage:  "Show the changes between commits, commit and working tree, etc",
		Action: diff,
	},
	{
		Name:   "fetch",
		Usage:  "Download objects and refs from another repository",
		Action: fetch,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "remote",
				Usage:   "The name of the remote repository to fetch from",
				Aliases: []string{"r"},
			},
		},
	},
	{
		Name:   "log",
		Usage:  "Show commit logs",
		Action: log,
	},
	{
		Name:   "merge",
		Usage:  "Join two or more development histories together",
		Action: merge,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "branch",
				Usage:    "Branch name to merge into the current branch",
				Required: true,
			},
		},
	},
	{
		Name:   "rebase",
		Usage:  "Reapply commits on top of another base tip",
		Action: rebase,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "branch",
				Usage:    "Branch name to rebase onto the current branch",
				Required: true,
			},
		},
	},
	{
		Name:   "remote",
		Usage:  "List the remote connections you have to other repositories",
		Action: remote,
	},
	{
		Name:   "reset",
		Usage:  "Reset current HEAD to the specified state",
		Action: reset,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "commit",
				Usage: "Commit hash to reset to",
			},
		},
	},
	{
		Name:   "stash",
		Usage:  "Stash the changes in a dirty working directory",
		Action: stash,
	},
	{
		Name:   "tag",
		Usage:  "Create, list, delete or verify a tag object signed with GPG",
		Action: tag,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "m",
				Usage:    "Tag message",
				Required: false,
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

	// Commands to add:
	// add, commit, status, branch, checkout, diff, fetch, log, merge, rebase, remote, reset, stash, tag, config

}
