package main

import (
	"context"
	"log"

	"github.com/andy-ahmedov/task_management_service/server/internal/repository/postgres"
	"github.com/andy-ahmedov/task_management_service/server/internal/service"
	"github.com/andy-ahmedov/task_management_service/server/pkg/psql"
	"github.com/andy-ahmedov/task_management_service/service_api/config"
	"github.com/andy-ahmedov/task_management_service/service_api/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.NewLogger()
	logg.Info(cfg)

	db, err := psql.ConnectToDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	repo := postgres.NewTaskRepository(db)
	service := service.NewTaskStorage(repo)

	service.CreateTask(context.Background(), nil)

}
