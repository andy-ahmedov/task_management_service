package transport

import (
	"context"
	"fmt"

	"github.com/andy-ahmedov/task_management_service/client/internal/domain"
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

	req := api.CreateRequest{
		Name:        name,
		Description: des,
		Status:      status,
	}
	_, err := c.TaskServiceClient.Create(ctx, &req)

	return err
}

func (c *Client) Get(ctx context.Context, id int64) (domain.Task, error) {
	req := api.GetRequest{
		ID: id,
	}

	task, err := c.TaskServiceClient.Get(ctx, &req)
	if err != nil {
		return domain.Task{}, err
	}

	return ConverteTime(task.Task), nil
}

func (c *Client) GetAll(ctx context.Context) ([]domain.Task, error) {
	req := api.GetAllRequest{}

	task, err := c.TaskServiceClient.GetAll(ctx, &req)
	if err != nil {
		return nil, err
	}

	tasks := ConvertTasks(task.Tasks)

	return tasks, nil
}

func (c *Client) Delete(ctx context.Context, id int64) error {
	req := api.DeleteRequest{
		ID: id,
	}

	_, err := c.TaskServiceClient.Delete(ctx, &req)

	return err
}

func (c *Client) Update(ctx context.Context, id int64, upd *domain.UpdateTaskInput) error {
	task := createShortTask(upd)

	if task.Name != nil {
		fmt.Println("transport/grpc.go", task.Name)
	}
	if task.Description != nil {
		fmt.Println("transport/grpc.go", task.Description)
	}
	if task.Status != nil {
		fmt.Println("transport/grpc.go", task.Status)
	}

	req := api.UpdateRequest{
		ID:   id,
		Task: task,
	}

	_, err := c.TaskServiceClient.Update(ctx, &req)

	return err
}

func ConverteTime(task *api.Task) domain.Task {
	newTask := domain.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		Created_at:  task.CreatedAt.AsTime(),
	}

	return newTask
}

func ConvertTasks(apiTasks []*api.Task) []domain.Task {
	tasks := make([]domain.Task, 0)

	for _, apiTask := range apiTasks {
		task := ConverteTime(apiTask)
		tasks = append(tasks, task)
	}

	return tasks
}

func createShortTask(upd *domain.UpdateTaskInput) *api.ShortTask {
	task := &api.ShortTask{}

	if upd.Name != nil {
		fmt.Println(upd.Name)
	} else if upd.Name == nil {
		fmt.Println("Name = nil")
		task.Name = nil
	}

	if upd.Description != nil {
		fmt.Println(upd.Description)
	} else if upd.Description == nil {
		fmt.Println("Descrip = nil")
	}

	if upd.Status != nil {
		fmt.Println(upd.Status)
	} else if upd.Status == nil {
		fmt.Println("Status = nil")
	}

	return task
}
