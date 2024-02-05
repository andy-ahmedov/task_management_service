package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_management_service/server/internal/domain"
)

type Task interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetTask(ctx context.Context, id int64) (domain.Task, error)
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	DeleteTask(ctx context.Context, id int64) error
	UpdateTask(ctx context.Context, id int64, input domain.UpdateTaskInput) error
}

type GRPCServer struct {
	tasksService Task
}

func NewGRPCServer(task Task) *GRPCServer {
	return &GRPCServer{
		tasksService: task,
	}
}

// func (g *GRPCServer) Add(ctx context.Context, req *api.CreateRequest) (*api.Empty, error) {
// 	return api.CreaterServer.Create(ctx, req)
// }
