package main

import (
	"context"
	"log"

	grpc_client "github.com/andy-ahmedov/task_management_service/client/internal/transport"
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

	createClient, err := grpc_client.NewClient(cfg.Srvr.Port)
	if err != nil {
		logg.Fatal(err)
	}

	// err = createClient.Create(context.Background(), &api.CreateRequest{
	// 	Name:        "HEYYO GRPC",
	// 	Description: "BELIEVE?",
	// 	Status:      "I DONT KNOW",
	// })

	err = createClient.Create(context.Background(), "DOUBLE", "WHY", "OMG")
	if err != nil {
		logg.Fatal(err)
	}
}
