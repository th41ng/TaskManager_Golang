package main

import (
	"context"
	"log"
	"net"

	"taskmanager/microservices/task-service/ent"
	"taskmanager/microservices/task-service/internal/service"
	pb "taskmanager/microservices/task-service/pb"

	"google.golang.org/grpc"
)

func main() {
	client, err := ent.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	taskService := service.NewTaskService(client)
	pb.RegisterTaskServiceServer(grpcServer, taskService)

	log.Println("TaskService listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
