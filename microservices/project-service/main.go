package main

import (
	"context"
	"log"
	"net"
	"os"

	"taskmanager/microservices/project-service/ent"
	"taskmanager/microservices/project-service/internal/service"
	pb "taskmanager/microservices/project-service/pb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Kết nối database
	dsn := os.Getenv("MYSQL_URL")
	if dsn == "" {
		dsn = "root:Thang@2004@tcp(mysql.service.consul:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
	}
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// migrate schema
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//Khởi tạo repository thật
	projectRepo := service.NewProjectRepo(client)

	//Tạo service và inject repository vào
	projectService := service.NewProjectService(projectRepo)

	//Khởi tạo server gRPC
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProjectServiceServer(grpcServer, projectService)

	// Register reflection for runtime introspection (grpcurl)
	reflection.Register(grpcServer)

	log.Println("✅ ProjectService is running on port :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
