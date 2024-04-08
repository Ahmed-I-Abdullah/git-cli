package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Options struct {
	Verbose bool
}

type CloneOptions struct {
	URL    string
	Dir    string
	Branch string
	Options
}

type PullOptions struct {
	URL    string
	Branch string
	Options
}

type PushOptions struct {
	Remote      string
	Branch      string
	PushAll     bool
	IncludeTags bool
	ForcePush   bool
	Options
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

func Push(opts PushOptions) error {
	if opts.Remote == "" {
		return fmt.Errorf("remote repository name is required for pushing")
	}

	args := []string{"push"}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.PushAll {
		args = append(args, "--all")
	}
	if opts.IncludeTags {
		args = append(args, "--tags")
	}
	if opts.ForcePush {
		args = append(args, "--force")
	}
	args = append(args, opts.Remote)

	fmt.Println("Executing command: git", strings.Join(args, " "))

	cmd := exec.Command("git", args...)
	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
