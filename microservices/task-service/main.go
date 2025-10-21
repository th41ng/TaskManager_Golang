package main

import (
	"context"
	"log"
	"net"

	"taskmanager/microservices/task-service/ent"
	"taskmanager/microservices/task-service/internal/service"
	pb "taskmanager/microservices/task-service/pb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	// Kết nối database
	//client, err := ent.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/task-db?charset=utf8mb4&parseTime=True&loc=Local")
	client, err := ent.Open("mysql", "root:123123@tcp(mysql-task:3306)/task-db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Auto migration
	// 2️⃣ Tự động migrate schema
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// 3️⃣ Tạo repository thật dùng Ent (dùng generic repo)
	taskRepo := service.NewTaskRepo(client)

	// 4️⃣ Tạo TaskService (business logic)
	taskService := service.NewTaskService(taskRepo)

	// 5️⃣ Khởi tạo gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, taskService)

	log.Println("✅ TaskService listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
