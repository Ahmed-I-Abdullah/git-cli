package git

import (
	"fmt"
	"os"
	"os/exec"
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

type AddOptions struct {
	FilePaths []string
	Options
}

type CommitOptions struct {
	Message string
	Options
}

type StatusOptions struct {
	Options
}

type BranchOptions struct {
	List   bool
	Create bool
	Delete bool
	Name   string
	Options
}

type CheckoutOptions struct {
	Branch string
	Options
}

type DiffOptions struct {
	Options
}

type FetchOptions struct {
	Remote string
	Options
}

type LogOptions struct {
	Options
}

type MergeOptions struct {
	Branch string
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

	cmd := exec.Command("git", args...)
	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

func Add(opts AddOptions) error {
	if len(opts.FilePaths) == 0 {
		return fmt.Errorf("at least one file path is required for adding")
	}

	args := []string{"add"}
	args = append(args, opts.FilePaths...)

	cmd := exec.Command("git", args...)
	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

func Commit(opts CommitOptions) error {
	if opts.Message == "" {
		return fmt.Errorf("commit message is required")
	}

	args := []string{"commit", "-m", opts.Message}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
func Status(opts StatusOptions) (string, error) {
	args := []string{"status"}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get status: %v", err)
	}

	statusOutput := "Current status of the repo."

	return statusOutput, nil
}

func Branch(opts BranchOptions) (string, error) {
	args := []string{"branch"}
	if opts.List {
		// No additional arguments needed to list branches
	} else if opts.Create {
		args = append(args, opts.Name)
	} else if opts.Delete {
		args = append(args, "-d", opts.Name)
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to manage branches: %v", err)
	}
	branchOutput := "Branch operation executed successfully."
	return branchOutput, nil
}

func Checkout(opts CheckoutOptions) error {
	if opts.Branch == "" {
		return fmt.Errorf("branch name is required for checkout")
	}

	args := []string{"checkout", opts.Branch}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Diff(opts DiffOptions) error {
	args := []string{"diff"}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Fetch(opts FetchOptions) error {
	args := []string{"fetch"}
	if opts.Remote != "" {
		args = append(args, opts.Remote)
	}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Log(opts LogOptions) error {
	args := []string{"log"}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Merge(opts MergeOptions) error {
	if opts.Branch == "" {
		return fmt.Errorf("branch name is required for merge")
	}

	args := []string{"merge", opts.Branch}
	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
