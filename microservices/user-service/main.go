package main

import (
	"context"
	"log"
	"net"

	"taskmanager/microservices/user-service/ent"
	"taskmanager/microservices/user-service/internal/service"
	pb "taskmanager/microservices/user-service/pb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	// Kết nối database MySQL
	client, err := ent.Open(
		"mysql",
		// "root:123123@tcp(127.0.0.1:3306)/user-db?charset=utf8mb4&parseTime=True&loc=Local",
		"root:123123@tcp(mysql-user:3306)/user-db?charset=utf8mb4&parseTime=True&loc=Local",
	)
	if err != nil {
		log.Fatalf("❌ failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Tự động migrate schema
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("❌ failed creating schema resources: %v", err)
	}

	// Tạo repository thật dùng Ent (dùng generic repo)
	userRepo := service.NewUserRepo(client)
	userQuery := service.NewUserQuery(client)

	//Tạo UserService (business logic)
	userService := service.NewUserService(userRepo, userQuery)

	// Khởi tạo gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("✅ UserService listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ failed to serve: %v", err)
	}
}
