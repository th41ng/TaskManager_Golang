package main

import (
	"context"
	"log"
	"net"

	"taskmanager/microservices/project-service/internal/service"
	pb "taskmanager/microservices/project-service/pb"

	"taskmanager/microservices/project-service/ent" // import ent package

	"google.golang.org/grpc"
)

func main() {
	// Kết nối database
	client, err := ent.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Auto migration
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Tạo listener cho gRPC
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Tạo service và truyền ent client vào
	projectService := service.NewProjectService(client)
	pb.RegisterProjectServiceServer(grpcServer, projectService)

	log.Println("ProjectService listening on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
