package main

import (
	"log"

	"github.com/andy-ahmedov/task_management_service/client/internal/commands"
	grpc_client "github.com/andy-ahmedov/task_management_service/client/internal/transport/grpc"
	"github.com/andy-ahmedov/task_management_service/service_api/config"
	"github.com/andy-ahmedov/task_management_service/service_api/logger"
	"github.com/spf13/cobra"
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

	var rootCmd = &cobra.Command{Use: "app"}
	cmdCreate := commands.CommandCreate(client, rootCmd)
	cmdUpdate := commands.CommandUpdate(client, rootCmd)
	cmdDelete := commands.CommandDelete(client, rootCmd)
	cmdGetAll := commands.CommandGetAll(client, rootCmd)
	cmdGet := commands.CommandGet(client, rootCmd)

	rootCmd.AddCommand(cmdCreate, cmdUpdate, cmdGet, cmdGetAll, cmdDelete)
	rootCmd.Execute()
}
