package service

import (
	"context"
	"log"

	pb "taskmanager/microservices/task-service/pb"
)

type TaskService struct {
	pb.UnimplementedTaskServiceServer
	repo TaskRepo
}

func NewTaskService(repo TaskRepo) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.repo.CreateTask(ctx, req.Title, int(req.ProjectId), int(req.Priority))
	if err != nil {
		log.Printf("failed to create task: %v", err)
		return nil, err
	}
	return &pb.TaskResponse{
		Id:        int32(t.ID),
		Title:     t.Title,
		Done:      t.Done,
		Priority:  int32(t.Priority),
		ProjectId: int32(t.ProjectID),
	}, nil
}

func (s *TaskService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.repo.GetTask(ctx, int(req.Id))
	if err != nil {
		log.Printf("failed to get task: %v", err)
		return nil, err
	}
	return &pb.TaskResponse{
		Id:        int32(t.ID),
		Title:     t.Title,
		Done:      t.Done,
		Priority:  int32(t.Priority),
		ProjectId: int32(t.ProjectID),
	}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.repo.UpdateTask(ctx, int(req.Id), req.Title, req.Done, int(req.Priority), int(req.ProjectId))
	if err != nil {
		log.Printf("failed to update task: %v", err)
		return nil, err
	}
	return &pb.TaskResponse{
		Id:        int32(t.ID),
		Title:     t.Title,
		Done:      t.Done,
		Priority:  int32(t.Priority),
		ProjectId: int32(t.ProjectID),
	}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.repo.DeleteTask(ctx, int(req.Id))
	if err != nil {
		log.Printf("failed to delete task: %v", err)
		return &pb.DeleteTaskResponse{Success: false}, err
	}
	return &pb.DeleteTaskResponse{Success: true}, nil
}

func (s *TaskService) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, err := s.repo.ListTasks(ctx)
	if err != nil {
		log.Printf("failed to list tasks: %v", err)
		return nil, err
	}
	resp := &pb.ListTasksResponse{}
	for _, t := range tasks {
		resp.Tasks = append(resp.Tasks, &pb.TaskResponse{
			Id:        int32(t.ID),
			Title:     t.Title,
			Done:      t.Done,
			Priority:  int32(t.Priority),
			ProjectId: int32(t.ProjectID),
		})
	}
	return resp, nil
}
