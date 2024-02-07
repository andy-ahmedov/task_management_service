package transport

import (
	"context"
	"fmt"

	"github.com/andy-ahmedov/task_management_service/service_api/api"
	"google.golang.org/grpc"
)

type Client struct {
	conn          *grpc.ClientConn
	CreaterClient api.CreaterClient
}

func NewClient(port int) (*Client, error) {
	// var conn *grpc.ClientConn
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	addr := fmt.Sprintf(":%d", port)

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:          conn,
		CreaterClient: api.NewCreaterClient(conn),
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) Create(ctx context.Context, req *api.CreateRequest) error {
	_, err := c.CreaterClient.Create(ctx, req)

	return err
}
