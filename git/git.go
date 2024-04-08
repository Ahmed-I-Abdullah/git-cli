package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Options structure contains verbose (indicates whether to show verbose output) for Git operations.
type Options struct {
	Verbose bool
}

// CloneOptions structure contains options (URL, Directory, and Branch) for the clone operation.
type CloneOptions struct {
	URL    string
	Dir    string
	Branch string
	Options
}

// PullOptions structure contains options (URL, and Branch) for the Pull operation.
type PullOptions struct {
	URL    string
	Branch string
	Options
}

// PushOptions structure contains options (Remote, Branch, PushAll, IncludeTags, and ForcePush) for the Push operation.
type PushOptions struct {
	Remote      string
	Branch      string
	PushAll     bool
	IncludeTags bool
	ForcePush   bool
	Options
}

/*
 * This function handles the git clone command for a repository with the specified options.
 */
func Clone(opts CloneOptions) error {
	//Need to check if a URL and Directory options are set, if not return an error
	if opts.URL == "" {
		return fmt.Errorf("URL is required for cloning")
	}
	if opts.Dir == "" {
		return fmt.Errorf("destination directory is required for cloning")
	}
	//Append the clone command and the options to the argument
	args := []string{"clone", opts.URL}
	if opts.Branch != "" {
		args = append(args, "--branch", opts.Branch)
	}
	args = append(args, opts.Dir)
	//execute the clone command using the command prompts
	cmd := exec.Command("git", args...)
	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

/*
 * This function handles the git Pull command for a repository with the specified options.
 */
func Pull(opts PullOptions) error {
	//Need to check if a URL option is set, if not return an error
	if opts.URL == "" {
		return fmt.Errorf("remote repository URL is required for pulling")
	}
	//Append the pull command and the options to the argument
	args := []string{"pull"}
	if opts.Verbose {
		args = append(args, "--verbose")
	}
	if opts.Branch != "" {
		args = append(args, opts.URL, opts.Branch)
	} else {
		args = append(args, opts.URL)
	}
	//execute the pull command using the command prompts
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

/*
 * This function handles the git Push command for a repository with the specified options.
 */
func Push(opts PushOptions) error {
	//Need to check if a remote option is set, if not return an error
	if opts.Remote == "" {
		return fmt.Errorf("remote repository name is required for pushing")
	}

	//Append the push command and the options to the argument
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

	//execute the push command using the command prompts
	cmd := exec.Command("git", args...)
	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
