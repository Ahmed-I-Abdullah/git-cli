package grpc

import (
	"fmt"

	"github.com/ammar-y62/git-cli/config"
	"github.com/ammar-y62/git-cli/constants"
)

func GetConnection() (*Client, error) {
	config, err := config.LoadConfig(constants.ConfigFile)
	if err != nil {
		return nil, err
	}

	if config.PeerURL == "" {
		return nil, fmt.Errorf("No peerUrl found. Please provide a peer URL using cli config")
	}

	grpcClient, err := SetupGRPCClient(config.PeerURL)
	if err != nil {
		return nil, err
	}

	return grpcClient, nil
}
