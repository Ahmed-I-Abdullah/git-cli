package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/ammar-y62/git-cli/constants"
	"google.golang.org/grpc"
)

// Structure of a gRPC client
type Client struct {
	conn *grpc.ClientConn
}

/*
 * This function sets up a gRPC client connection with the specified peer URL.
 * Returns a pointer to the client or an error if the operation fails.
 */
func SetupGRPCClient(peerURL string) (*Client, error) {
	//if no peer URL is provided return an error
	if peerURL == "" {
		return nil, errors.New("no peer URL provided")
	}
	//attempt to dial peer with gRPC using the peer URL
	grpcClientConn, err := grpc.Dial(peerURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to GRPC peer: %v", err)
	}
	//returns pointer to the client
	return &Client{conn: grpcClientConn}, nil
}

/*
 * This function closes the gRPC client connection.
 */
func (c *Client) Close() error {
	return c.conn.Close()
}

/*
 * This function gets the gRPC client connection.
 */
func (c *Client) GetConn() *grpc.ClientConn {
	return c.conn
}

/*
 * This function gets a context with a timeout based on the predefined TimeoutDuration.
 */
func (c *Client) GetContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), constants.TimeoutDuration)
}
