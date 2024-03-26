package git

import (
	"fmt"
	"os"
	"os/exec"
)

type CloneOptions struct {
	URL     string
	Dir     string
	Branch  string
	Verbose bool
}

func Clone(opts CloneOptions) error {
	if opts.URL == "" {
		return fmt.Errorf("URL is required for cloning")
	}
	if opts.Dir == "" {
		return fmt.Errorf("destination directory is required for cloning")
	}

	args := []string{"clone", opts.URL}
	if opts.Branch != "" {
		args = append(args, "--branch", opts.Branch)
	}
	args = append(args, opts.Dir)

	cmd := exec.Command("git", args...)
	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

type PullOptions struct {
	URL     string
	Branch  string
	Verbose bool
}

func Pull(opts PullOptions) error {
	if opts.URL == "" {
		return fmt.Errorf("remote repository URL is required for pulling")
	}

	args := []string{"pull"}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Branch != "" {
		args = append(args, opts.URL, opts.Branch)
	} else {
		args = append(args, opts.URL)
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
