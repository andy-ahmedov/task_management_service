package main

import (
	"context"
	"log"

	"github.com/andy-ahmedov/task_management_service/client/internal/domain"
	grpc_client "github.com/andy-ahmedov/task_management_service/client/internal/transport/grpc"
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

	client, err := grpc_client.NewClient(cfg.Srvr.Port)
	if err != nil {
		logg.Fatal(err)
	}

	// err = client.Create(context.Background(), "NEW Task", "HRA", "ON")
	// if err != nil {
	// 	logg.Fatal(err)
	// }

	// task, err := client.Get(context.Background(), 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(task)

	// tasks, err := client.GetAll(context.Background())
	// fmt.Println(tasks)

	// err = client.Delete(context.Background(), 2)
	// if err != nil {
	// 	logg.Fatal(err)
	// }

	that := "FAST"

	upd := domain.UpdateTaskInput{
		Name:        &that,
		Description: nil,
		Status:      &that,
	}

	err = client.Update(context.Background(), 5, &upd)
	if err != nil {
		logg.Fatal(err)
	}
}
