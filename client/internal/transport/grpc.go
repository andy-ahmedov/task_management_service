package transport

import (
	"context"
	"fmt"

	"github.com/andy-ahmedov/task_management_service/service_api/api"
	"google.golang.org/grpc"
)

type Client struct {
	conn              *grpc.ClientConn
	TaskServiceClient api.TaskServiceClient
}

func NewClient(port int) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	addr := fmt.Sprintf(":%d", port)

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:              conn,
		TaskServiceClient: api.NewTaskServiceClient(conn),
	}, nil
}

func (c *Client) CloseConnection() error {
	return c.conn.Close()
}

func (c *Client) Create(ctx context.Context, name, des, status string) error {

	task := api.ShortTask{
		Name:        name,
		Description: des,
		Status:      status,
	}

	req := api.CreateRequest{
		Task: &task,
	}
	_, err := c.TaskServiceClient.Create(ctx, &req)

	return err
}
