package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_management_service/service_api/api"
	"github.com/andy-ahmedov/task_management_service/service_api/logger"
)

type Creater interface {
	CreateTask(ctx context.Context, req *api.CreateRequest) error
}

type CreaterServer struct {
	service Creater
	api.UnimplementedCreaterServer
}

func NewCreaterServer(service Creater) *CreaterServer {
	return &CreaterServer{
		service: service,
	}
}

func (h *CreaterServer) Create(ctx context.Context, req *api.CreateRequest) (*api.Empty, error) {
	err := h.service.CreateTask(ctx, req)
	if err == nil {
		logg := logger.NewLogger()
		logg.Info("COMPLETE")
	}

	return &api.Empty{}, err
}
