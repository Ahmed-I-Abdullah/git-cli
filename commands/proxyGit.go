package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ProxyToGit(ctx *cli.Context, command string, args ...string) {
	fullArgs := append([]string{command}, args...)

	cmd := exec.Command("git", fullArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error executing git command: %v\n", err)
		os.Exit(1)
	}
}
