all: run


run: 
	go run server/cmd/main.go

up:
	docker-compose up -d db

stop_and_delete_container:
	docker stop task_manager
	docker rm task_manager
	docker image rmi task_management_service-db:latest

create_table:
	docker exec -it task_manager psql -U postgres -d task_service -c "\i script.sql"

proto:
	protoc -I service_api/proto --go_out=service_api/api service_api/proto/task_manager.proto
	protoc -I service_api/proto --go-grpc_out=service_api/api service_api/proto/task_manager.proto
