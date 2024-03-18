package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Ahmed-I-Abdullah/p2p-code-collaboration/pb"
	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var configDir string
var configFile string

type Config struct {
	PeerURL string `json:"peer_url"`
}

func main() {
	if err := ensureConfigFile(); err != nil {
		log.Fatalf("failed to ensure config file: %v", err)
	}

	app := &cli.App{
		Name:  "git-grpc-wrapper",
		Usage: "A CLI wrapper for Git commands with GRPC support for specific tasks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "peer-url",
				Usage: "URL of the GRPC peer to connect to for remote operations",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Initialize a new Git repository via GRPC",
				Action: initViaGRPC,
			},
			{
				Name:   "pull",
				Usage:  "Pull changes from a Git repository via GRPC",
				Action: pullViaGRPC,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "repo-name",
						Usage:    "Name of the repository to pull changes from",
						Required: true,
					},
				},
			},
			{
				Name:   "clone",
				Usage:  "Clone a Git repository via GRPC",
				Action: cloneViaGRPC,
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
			// Placeholder for clone and push commands
			// Add similar structs for each operation as needed
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	mainCleanup()
}

// Global GRPC client connection
var grpcClientConn *grpc.ClientConn
var config Config

func ensureConfigFile() error {
	configDir = getConfigDir()
	configFile = filepath.Join(configDir, "config.json")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0700); err != nil {
			return fmt.Errorf("failed to create config directory: %v", err)
		}
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err := saveConfig(); err != nil {
			return fmt.Errorf("failed to create config file: %v", err)
		}
	}

	return nil
}

func getConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to determine user's home directory: %v", err)
	}
	return filepath.Join(home, ".git-p2p-wrapper")
}

func setupGRPCClient() error {
	if err := loadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	if config.PeerURL == "" {
		return nil // No GRPC setup required if no peer URL provided
	}

	var err error
	grpcClientConn, err = grpc.Dial(config.PeerURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to GRPC peer: %v", err)
	}

	return nil
}

func executeGitCommand(c *cli.Context) error {
	args := c.Args().Slice()
	if len(args) < 1 {
		return fmt.Errorf("no Git command provided")
	}

	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git command error: %v, output: %s", err, output)
	}

	fmt.Printf("Git command output:\n%s\n", output)
	return nil
}

func initViaGRPC(c *cli.Context) error {
	if err := setupGRPCClient(); err != nil {
		return err
	}

	repoName := c.Args().First()
	if repoName == "" {
		return fmt.Errorf("you must provide a repository name for initialization")
	}
	if grpcClientConn == nil {
		return fmt.Errorf("GRPC client is not setup. Please provide a peer URL.")
	}

	grpcClient := pb.NewRepositoryClient(grpcClientConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := grpcClient.Init(ctx, &pb.RepoInitRequest{Name: repoName, FromCli: true})
	if err != nil {
		return fmt.Errorf("failed to initialize repository via GRPC: %v", err)
	}

	fmt.Printf("Repository initialized successfully. Server Response: %s\n", response.Message)
	return nil
}

func pullViaGRPC(c *cli.Context) error {
	if err := setupGRPCClient(); err != nil {
		return err
	}

	repoName := c.String("repo-name")
	if repoName == "" {
		return fmt.Errorf("you must provide a repository name for pulling changes")
	}

	if grpcClientConn == nil {
		return fmt.Errorf("GRPC client is not setup. Please provide a peer URL.")
	}

	grpcClient := pb.NewRepositoryClient(grpcClientConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := grpcClient.Pull(ctx, &pb.RepoPullRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("failed to pull repository via GRPC: %v", err)
	}

	fmt.Printf("Repository pull successful. Server Response: %s\n", response.RepoAddress)
	return nil
}

func cloneViaGRPC(c *cli.Context) error {
	if err := setupGRPCClient(); err != nil {
		return err
	}

	repoName := c.Args().First()
	if repoName == "" {
		return fmt.Errorf("you must provide a repository name for cloning")
	}

	if grpcClientConn == nil {
		return fmt.Errorf("GRPC client is not setup. Please provide a peer URL.")
	}

	grpcClient := pb.NewRepositoryClient(grpcClientConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := grpcClient.Pull(ctx, &pb.RepoPullRequest{Name: repoName})
	if err != nil {
		return fmt.Errorf("failed to clone repository via GRPC: %v", err)
	}

	if !response.Success {
		return fmt.Errorf("failed to clone repository: server returned unsuccessful response")
	}

	fmt.Printf("Repository clone successful. Repo Address: %s\n", response.RepoAddress)
	return executeInternalGitCommand(c, "clone", response.RepoAddress)
}

func executeInternalGitCommand(c *cli.Context, command string, args ...string) error {
	cmdArgs := append([]string{command}, args...)
	cmd := exec.Command("git", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git command error: %v, output: %s", err, output)
	}

	fmt.Printf("Git command output:\n%s\n", output)
	return nil
}

func setConfig(c *cli.Context) error {
	peerURL := c.String("peer-url")
	if peerURL == "" {
		return fmt.Errorf("peer URL is required")
	}

	config.PeerURL = peerURL

	if err := saveConfig(); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	fmt.Printf("Peer URL set to: %s\n", peerURL)
	return nil
}

func saveConfig() error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, data, 0644)
}

func loadConfig() error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &config)
}

func mainCleanup() {
	if grpcClientConn != nil {
		grpcClientConn.Close()
	}
}
