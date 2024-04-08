package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

/*
 * This function proxies the given git command to the system's git executable.
 */
func ProxyToGit(ctx *cli.Context, command string, args ...string) {
	//Append the command and its arguments
	fullArgs := append([]string{command}, args...)

	//Execute a cmd instance for running the git command
	cmd := exec.Command("git", fullArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//Execute the cmd command
	err := cmd.Run()
	//If an error occurs running the cmd command, respond with an error message
	if err != nil {
		fmt.Fprintf(os.Stderr, "error executing git command: %v\n", err)
		os.Exit(1)
	}
}
