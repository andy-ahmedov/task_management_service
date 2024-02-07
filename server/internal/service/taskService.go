package service

import (
	"context"
	"fmt"
	"time"

	"github.com/andy-ahmedov/task_management_service/server/internal/domain"
	"github.com/andy-ahmedov/task_management_service/service_api/api"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	Get(ctx context.Context, id int64) (domain.Task, error)
	GetAll(ctx context.Context) ([]domain.Task, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, updTask domain.UpdateTaskInput) error
}

type TasksStorage struct {
	repo TaskRepository
}

func NewTaskStorage(repo TaskRepository) *TasksStorage {
	return &TasksStorage{
		repo: repo,
	}
}

func (t *TasksStorage) CreateTask(ctx context.Context, req *api.CreateRequest) error {
	task := domain.Task{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Created_at:  time.Now(),
	}

	return t.repo.Create(ctx, &task)

}

func (t *TasksStorage) GetTask(ctx context.Context, id int64) (domain.Task, error) {
	return t.repo.Get(ctx, id)
}

func (t *TasksStorage) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	return t.repo.GetAll(ctx)
}

func (t *TasksStorage) DeleteTask(ctx context.Context, id int64) error {
	return t.repo.Delete(ctx, id)
}

func (t *TasksStorage) UpdateTask(ctx context.Context, req *api.UpdateRequest) error {
	// task := domain.UpdateTaskInput{
	// 	Name:        &req.Task.Name,
	// 	Description: &req.Task.Description,
	// 	Status:      &req.Task.Status,
	// }

	var name, description, status string

	if req.Task.Name != nil {
		fmt.Println("taskService.go", req.Task.Name)
		name = req.Task.Name.Value
	}

	if req.Task.Description != nil {
		fmt.Println("taskService.go", req.Task.Description)
		description = req.Task.Description.Value
	}

	if req.Task.Status != nil {
		fmt.Println("taskService.go", req.Task.Status)
		status = req.Task.Status.Value
	}

	task := domain.UpdateTaskInput{
		Name:        &name,
		Description: &description,
		Status:      &status,
	}

	return t.repo.Update(ctx, req.ID, task)
}
