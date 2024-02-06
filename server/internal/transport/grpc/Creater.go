package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_management_service/service_api/api"
)

type Creater interface {
	CreateTask(ctx context.Context, req *api.CreateRequest) error
}

type CreaterServer struct {
	service Creater
}

func NewCreaterServer(service Creater) *CreaterServer {
	return &CreaterServer{
		service: service,
	}
}

func (h *CreaterServer) mustEmbedUnimplementedCreaterServer() {
}

func (h *CreaterServer) Create(ctx context.Context, req *api.CreateRequest) (*api.Empty, error) {
	err := h.service.CreateTask(ctx, req)

	return &api.Empty{}, err
}
