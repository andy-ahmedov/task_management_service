package service

import (
	"context"
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
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Status:      req.GetStatus(),
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

func (t *TasksStorage) UpdateTask(ctx context.Context, id int64, input domain.UpdateTaskInput) error {
	return t.repo.Update(ctx, id, input)
}
