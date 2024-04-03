package git

import (
	"bytes"
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

type BranchOptions struct {
	List   bool
	Create bool
	Name   string
	Options
}

type CheckoutOptions struct {
	Branch string
	Options
}

type FetchOptions struct {
	Remote string
	Options
}

type MergeOptions struct {
	Branch string
}
type RebaseOptions struct {
	Branch string
}

type RemoteOptions struct {
	Options
}

type ResetOptions struct {
	Commit string
}

type StashOptions struct {
	Options
}

type TagOptions struct {
	Name    string
	Message string
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
	args := []string{"add"}
	args = append(args, opts.FilePaths...)

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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
func Status() (string, error) {
	cmd := exec.Command("git", "status")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func Branch(opts BranchOptions) (string, error) {
	var args []string
	var out bytes.Buffer

	if opts.List {
		args = []string{"branch", "--list"}
	} else if opts.Create && opts.Name != "" {
		args = []string{"branch", opts.Name}
	} else {
		return "", fmt.Errorf("either List must be true or Create must be true with a non-empty Name")
	}

	if opts.Verbose {
		args = append(args, "--verbose")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	return out.String(), err
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

func Diff() error {
	cmd := exec.Command("git", "diff")
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

func Log() error {
	cmd := exec.Command("git", "log")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	println(out.String())
	return nil
}

func Merge(opts MergeOptions) error {
	args := []string{"merge", opts.Branch}
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Rebase(opts RebaseOptions) error {
	if opts.Branch == "" {
		return fmt.Errorf("branch name is required for rebase")
	}

	cmd := exec.Command("git", "rebase", opts.Branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Remote() (string, error) {
	cmd := exec.Command("git", "remote", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func Reset(opts ResetOptions) error {
	args := []string{"reset"}
	if opts.Commit != "" {
		args = append(args, opts.Commit)
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Stash() error {
	cmd := exec.Command("git", "stash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Tag(opts TagOptions) error {
	var args []string
	if opts.Message != "" {
		args = []string{"tag", "-a", opts.Name, "-m", opts.Message}
	} else {
		args = []string{"tag", opts.Name}
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
