package main

import (
	"context"
	"log"
	"net"
	"os"

	"taskmanager/microservices/user-service/ent"
	"taskmanager/microservices/user-service/internal/service"
	pb "taskmanager/microservices/user-service/pb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
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

	// Register reflection service on gRPC server so tools like grpcurl can introspect
	reflection.Register(grpcServer)

	log.Println("✅ UserService listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ failed to serve: %v", err)
	}
}
