package main

import (
	"log"
	"net/http"
	"os"

	"taskmanager/gateway/handlers"
	pb "taskmanager/gateway/pb"

	"taskmanager/gateway/middleware"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Adapter chuyển net/http middleware sang Gin middleware

func AdaptNetHTTPMiddleware(mw func(http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		finished := false
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
			finished = true
		})
		mw(h).ServeHTTP(c.Writer, c.Request)
		if !finished {
			c.Abort()
		}
	}
}

type ginResponseWriter struct {
	http.ResponseWriter
	ResponseWriterGin gin.ResponseWriter
}

func (g *ginResponseWriter) WriteHeader(statusCode int) {
	g.ResponseWriterGin.WriteHeader(statusCode)
	g.ResponseWriter.WriteHeader(statusCode)
}

func (g *ginResponseWriter) Write(b []byte) (int, error) {
	g.ResponseWriterGin.Write(b)
	return g.ResponseWriter.Write(b)
}

func main() {
	// Userconnect, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to connect to user-service: %v", err)
	// }
	// defer Userconnect.Close()

	// Taskconnect, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to connect to user-service: %v", err)
	// }
	// defer Taskconnect.Close()

	// ProjectConnect, err := grpc.Dial("localhost:5003", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to connect to user-service: %v", err)
	// }
	// defer ProjectConnect.Close()
	// Userconnect, err := grpc.Dial("user-service:50051", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to connect to user-service: %v", err)
	// }
	// defer Userconnect.Close()

	// Taskconnect, err := grpc.Dial("task-service:50052", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to connect to task-service: %v", err)
	// }
	// defer Taskconnect.Close()

	// ProjectConnect, err := grpc.Dial("project-service:50053", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to connect to project-service: %v", err)
	// }
	// defer ProjectConnect.Close()

	userAddr := os.Getenv("USER_SERVICE_URL")
	taskAddr := os.Getenv("TASK_SERVICE_URL")
	projectAddr := os.Getenv("PROJECT_SERVICE_URL")

	creds := credentials.NewClientTLSFromCert(nil, "")

	Userconnect, err := grpc.Dial(userAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to user-service: %v", err)
	}
	defer Userconnect.Close()

	Taskconnect, err := grpc.Dial(taskAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to task-service: %v", err)
	}
	defer Taskconnect.Close()

	ProjectConnect, err := grpc.Dial(projectAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to project-service: %v", err)
	}
	defer ProjectConnect.Close()

	userClient := pb.NewUserServiceClient(Userconnect)
	taskClient := pb.NewTaskServiceClient(Taskconnect)
	projectClient := pb.NewProjectServiceClient(ProjectConnect)
	userHandler := handlers.NewUserHandler(userClient)
	taskHandler := handlers.NewTaskHandler(taskClient)
	projectHandler := handlers.NewProjectHandler(projectClient)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorLogger())

	// Thêm các middleware
	r.Use(
		AdaptNetHTTPMiddleware(middleware.CORS),
		AdaptNetHTTPMiddleware(middleware.RateLimit),
	)

	// Public routes (no AuthJWT)
	r.POST("/login", userHandler.Login)
	r.POST("/users", userHandler.CreateUser)

	// Protected routes (require AuthJWT)
	authGroup := r.Group("/")
	authGroup.Use(AdaptNetHTTPMiddleware(middleware.AuthJWT))

	authGroup.GET("/users/:id", userHandler.GetUser)
	authGroup.PUT("/users/:id", userHandler.UpdateUser)
	authGroup.DELETE("/users/:id", userHandler.DeleteUser)

	authGroup.POST("/task", taskHandler.CreateTask)
	authGroup.GET("/task/:id", taskHandler.GetTask)
	authGroup.PUT("/task/:id", taskHandler.UpdateTask)
	authGroup.DELETE("/task/:id", taskHandler.DeleteTask)

	authGroup.POST("/project", projectHandler.CreateProject)
	authGroup.GET("/project/:id", projectHandler.GetProject)
	authGroup.PUT("/project/:id", projectHandler.UpdateProject)
	authGroup.DELETE("/project/:id", projectHandler.DeleteProject)

	log.Println("Gateway running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
