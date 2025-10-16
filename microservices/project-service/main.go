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
)

func main() {
	// Lấy DSN MySQL từ biến môi trường (ưu tiên Railway)
	dsn := os.Getenv("MYSQL_URL")
	if dsn != "" && len(dsn) > 9 && dsn[:8] == "mysql://" {
		// Chuyển mysql://user:pass@host:port/dbname?params thành user:pass@tcp(host:port)/dbname?params
		dsn = dsn[8:]
		atIdx := -1
		for i, c := range dsn {
			if c == '@' {
				atIdx = i
				break
			}
		}
		if atIdx != -1 {
			auth := dsn[:atIdx]
			hostDb := dsn[atIdx+1:]
			// hostDb: host:port/dbname?params
			slashIdx := -1
			for i, c := range hostDb {
				if c == '/' {
					slashIdx = i
					break
				}
			}
			if slashIdx != -1 {
				host := hostDb[:slashIdx]
				dbAndParams := hostDb[slashIdx:]
				host = "@tcp(" + host + ")"
				dsn = auth + host + dbAndParams
			}
		}
	}
	if dsn == "" {
		dsn = "root:123123@tcp(127.0.0.1:3306)/project-db?charset=utf8mb4&parseTime=True&loc=Local"
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

	log.Println("✅ ProjectService is running on port :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
