package service

import (
	"context"
	"log"

	"taskmanager/microservices/task-service/ent"
	"taskmanager/microservices/task-service/ent/task"
	pb "taskmanager/microservices/task-service/pb"
)

type TaskService struct {
	pb.UnimplementedTaskServiceServer
	client *ent.Client
}

// Constructor
func NewTaskService(client *ent.Client) *TaskService {
	return &TaskService{client: client}
}

// CreateTask
func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.client.Task.
		Create().
		SetTitle(req.Title).
		SetProjectID(int(req.ProjectId)).
		SetPriority(int(req.Priority)).
		SetDone(false).
		Save(ctx)
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

// GetTask
func (s *TaskService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.client.Task.Query().
		Where(task.ID(int(req.Id))).
		Only(ctx)
	if err != nil {
		log.Printf("task %d not found: %v", req.Id, err)
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

// UpdateTask
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.client.Task.
		UpdateOneID(int(req.Id)).
		SetTitle(req.Title).
		SetDone(req.Done).
		SetPriority(int(req.Priority)).
		SetProjectID(int(req.ProjectId)).
		Save(ctx)
	if err != nil {
		log.Printf("failed to update task %d: %v", req.Id, err)
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

// DeleteTask
func (s *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.client.Task.DeleteOneID(int(req.Id)).Exec(ctx)
	if err != nil {
		log.Printf("failed to delete task %d: %v", req.Id, err)
		return &pb.DeleteTaskResponse{Success: false}, err
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}

// ListTasks
func (s *TaskService) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, err := s.client.Task.Query().All(ctx)
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
