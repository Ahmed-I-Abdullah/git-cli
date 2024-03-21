package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/ammar-y62/git-cli/constants"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
}

func SetupGRPCClient(peerURL string) (*Client, error) {
	if peerURL == "" {
		return nil, errors.New("no peer URL provided")
	}

	grpcClientConn, err := grpc.Dial(peerURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to GRPC peer: %v", err)
	}

	return &Client{conn: grpcClientConn}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *Client) GetContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), constants.TimeoutDuration)
}
