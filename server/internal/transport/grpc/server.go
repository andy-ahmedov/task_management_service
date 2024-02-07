package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/andy-ahmedov/task_management_service/server/internal/domain"
	"github.com/andy-ahmedov/task_management_service/service_api/api"
	"google.golang.org/grpc"
)

type Task interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetTask(ctx context.Context, id int64) (domain.Task, error)
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	DeleteTask(ctx context.Context, id int64) error
	UpdateTask(ctx context.Context, id int64, input domain.UpdateTaskInput) error
}

type Server struct {
	grpcSrv    *grpc.Server
	taskServer api.TaskServiceServer
}

func New(taskServ api.TaskServiceServer) *Server {
	return &Server{
		grpcSrv:    grpc.NewServer(),
		taskServer: taskServ,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// api.RegisterCreaterServer(s.grpcSrv, s.taskServer)
	api.RegisterTaskServiceServer(s.grpcSrv, s.taskServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return err
}

// func (g *GRPCServer) Add(ctx context.Context, req *api.CreateRequest) (*api.Empty, error) {
// 	return api.CreaterServer.Create(ctx, req)
// }
