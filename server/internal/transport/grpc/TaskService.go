package grpc

import (
	"context"

	"github.com/andy-ahmedov/task_management_service/server/internal/domain"
	"github.com/andy-ahmedov/task_management_service/service_api/api"
	"github.com/andy-ahmedov/task_management_service/service_api/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskService interface {
	CreateTask(ctx context.Context, req *api.CreateRequest) error
	GetTask(ctx context.Context, id int64) (domain.Task, error)
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	DeleteTask(ctx context.Context, id int64) error
	UpdateTask(ctx context.Context, req *api.UpdateRequest) error
}

type TaskServiceServer struct {
	service TaskService
	api.UnimplementedTaskServiceServer
	logger *logrus.Logger
}

func NewCreaterServer(service TaskService, logger *logrus.Logger) *TaskServiceServer {
	return &TaskServiceServer{
		service: service,
		logger:  logger,
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

func (h *TaskServiceServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	task, err := h.service.GetTask(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	apiTask := ConvertToApiTask(task)

	resp := api.GetResponse{
		Task: apiTask,
	}

	return &resp, err
}

func (h *TaskServiceServer) GetAll(ctx context.Context, req *api.GetAllRequest) (*api.GetAllResponse, error) {
	tasks, err := h.service.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	apiTasks := make([]*api.Task, 0)

	for _, task := range tasks {
		apiTask := ConvertToApiTask(task)
		apiTasks = append(apiTasks, apiTask)
	}

	res := api.GetAllResponse{
		Tasks: apiTasks,
	}

	return &res, nil
}

func (h *TaskServiceServer) Delete(ctx context.Context, req *api.DeleteRequest) (*api.Empty, error) {
	err := h.service.DeleteTask(ctx, req.ID)

	return &api.Empty{}, err
}

func (h *TaskServiceServer) Update(ctx context.Context, req *api.UpdateRequest) (*api.Empty, error) {
	err := h.service.UpdateTask(ctx, req)

	return &api.Empty{}, err
}

func ConvertToApiTask(task domain.Task) *api.Task {
	created_at := timestamppb.New(task.Created_at)

	apitask := api.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   created_at,
	}

	return &apitask
}
