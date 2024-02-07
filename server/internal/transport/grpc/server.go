package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/andy-ahmedov/task_management_service/service_api/api"
	"google.golang.org/grpc"
)

// type Task interface {
// 	CreateTask(ctx context.Context, task *domain.Task) error
// 	GetTask(ctx context.Context, id int64) (domain.Task, error)
// 	GetAllTasks(ctx context.Context) ([]domain.Task, error)
// 	DeleteTask(ctx context.Context, id int64) error
// 	UpdateTask(ctx context.Context, id int64, input domain.UpdateTaskInput) error
// }

type Server struct {
	grpcSrv    *grpc.Server
	taskServer api.TaskServiceServer
}

func New(taskServ api.TaskServiceServer) *Server {
	return &Server{
		grpcSrv:    grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor)),
		taskServer: taskServ,
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
	return resp, err
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	api.RegisterTaskServiceServer(s.grpcSrv, s.taskServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return err
}

// func (g *GRPCServer) Add(ctx context.Context, req *api.CreateRequest) (*api.Empty, error) {
// 	return api.CreaterServer.Create(ctx, req)
// }
