package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_management_service/service_api/api"
	"github.com/andy-ahmedov/task_management_service/service_api/logger"
)

type TaskService interface {
	CreateTask(ctx context.Context, req *api.CreateRequest) error
}

type TaskServiceServer struct {
	service TaskService
	api.UnimplementedTaskServiceServer
}

func NewCreaterServer(service TaskService) *TaskServiceServer {
	return &TaskServiceServer{
		service: service,
	}
}

func (h *TaskServiceServer) Create(ctx context.Context, req *api.CreateRequest) (*api.Empty, error) {
	err := h.service.CreateTask(ctx, req)
	if err == nil {
		logg := logger.NewLogger()
		logg.Info("COMPLETE")
	}

	return &api.Empty{}, err
}
