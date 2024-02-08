package service

import (
	"context"
	"time"

	"github.com/andy-ahmedov/task_management_service/server/internal/domain"
	"github.com/andy-ahmedov/task_management_service/service_api/api"
)

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
	task := ConvertToDomainUpdateTask(req)

	return t.repo.Update(ctx, req.ID, task)
}
