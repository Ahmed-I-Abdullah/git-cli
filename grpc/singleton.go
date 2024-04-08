package grpc

import (
	"fmt"

	"github.com/ammar-y62/git-cli/config"
	"github.com/ammar-y62/git-cli/constants"
)

/*
 * This function establishes a gRPC connection with the configured peer URL.
 * It loads the configuration, checks if a peer URL is provided, and sets up the gRPC client.
 * Returns a gRPC client instance and an error, if any.
 */
func GetConnection() (*Client, error) {
	// Load the configuration using the config file
	config, err := config.LoadConfig(constants.ConfigFile)
	if err != nil {
		return nil, err
	}
	//check if a peer URL is provided
	if config.PeerURL == "" {
		return nil, fmt.Errorf("No peerUrl found. Please provide a peer URL using cli config")
	}
	// Set up the gRPC client
	grpcClient, err := SetupGRPCClient(config.PeerURL)
	if err != nil {
		return nil, err
	}
	//return the gRPC client
	return grpcClient, nil
}
