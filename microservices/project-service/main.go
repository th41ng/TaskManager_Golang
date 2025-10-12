package main

import (
	"context"
	"log"
	"net"

	"taskmanager/microservices/project-service/ent"
	"taskmanager/microservices/project-service/internal/service"
	pb "taskmanager/microservices/project-service/pb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	// 1️⃣ Kết nối database MySQL (repo thật)
	// client, err := ent.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/project-db?charset=utf8mb4&parseTime=True&loc=Local")
	client, err := ent.Open("mysql", "root:123123@tcp(mysql-project:3306)/project-db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// 2️⃣ Tự động migrate schema
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// 3️⃣ Khởi tạo repository thật
	repo := service.NewEntProjectRepo(client)

	// 4️⃣ Tạo service và inject repository vào
	projectService := service.NewProjectService(repo)

	// 5️⃣ Khởi tạo server gRPC
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProjectServiceServer(grpcServer, projectService)

	log.Println("✅ ProjectService is running on port :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
