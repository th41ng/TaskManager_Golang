package service

import (
	"context"
	"errors"
	"testing"

	"taskmanager/microservices/task-service/ent"
	"taskmanager/microservices/task-service/pb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(ctx context.Context, title string, projectID, priority int) (*ent.Task, error) {
	args := m.Called(ctx, title, projectID, priority)
	return args.Get(0).(*ent.Task), args.Error(1)
}
func (m *MockTaskRepository) GetByID(ctx context.Context, id int) (*ent.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*ent.Task), args.Error(1)
}
func (m *MockTaskRepository) Update(ctx context.Context, id int, title string, done bool, priority, projectID int) (*ent.Task, error) {
	args := m.Called(ctx, id, title, done, priority, projectID)
	return args.Get(0).(*ent.Task), args.Error(1)
}
func (m *MockTaskRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockTaskRepository) List(ctx context.Context) ([]*ent.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*ent.Task), args.Error(1)
}

type TaskServiceWithRepo struct {
	repo TaskRepository
}

func NewTaskServiceWithRepo(repo TaskRepository) *TaskServiceWithRepo {
	return &TaskServiceWithRepo{repo: repo}
}

func (s *TaskServiceWithRepo) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.repo.Create(ctx, req.Title, int(req.ProjectId), int(req.Priority))
	if err != nil {
		return nil, err
	}
	return &pb.TaskResponse{Id: int32(t.ID), Title: t.Title, ProjectId: int32(t.ProjectID), Priority: int32(t.Priority), Done: t.Done}, nil
}

func (s *TaskServiceWithRepo) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.repo.GetByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.TaskResponse{Id: int32(t.ID), Title: t.Title, ProjectId: int32(t.ProjectID), Priority: int32(t.Priority), Done: t.Done}, nil
}

func (s *TaskServiceWithRepo) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	t, err := s.repo.Update(ctx, int(req.Id), req.Title, req.Done, int(req.Priority), int(req.ProjectId))
	if err != nil {
		return nil, err
	}
	return &pb.TaskResponse{Id: int32(t.ID), Title: t.Title, ProjectId: int32(t.ProjectID), Priority: int32(t.Priority), Done: t.Done}, nil
}

func (s *TaskServiceWithRepo) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.repo.Delete(ctx, int(req.Id))
	if err != nil {
		return &pb.DeleteTaskResponse{Success: false}, err
	}
	return &pb.DeleteTaskResponse{Success: true}, nil
}

func (s *TaskServiceWithRepo) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListTasksResponse{}
	for _, t := range tasks {
		resp.Tasks = append(resp.Tasks, &pb.TaskResponse{
			Id: int32(t.ID), Title: t.Title, ProjectId: int32(t.ProjectID), Priority: int32(t.Priority), Done: t.Done,
		})
	}
	return resp, nil
}

func TestCreateTask_Success(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	task := &ent.Task{ID: 1, Title: "task1", ProjectID: 2, Priority: 1, Done: false}
	repo.On("Create", ctx, "task1", 2, 1).Return(task, nil)
	resp, err := svc.CreateTask(ctx, &pb.CreateTaskRequest{Title: "task1", ProjectId: 2, Priority: 1})
	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, "task1", resp.Title)
	assert.Equal(t, int32(2), resp.ProjectId)
	assert.Equal(t, int32(1), resp.Priority)
	assert.False(t, resp.Done)
}

func TestCreateTask_Error(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Create", ctx, "task1", 2, 1).Return((*ent.Task)(nil), errors.New("fail"))
	resp, err := svc.CreateTask(ctx, &pb.CreateTaskRequest{Title: "task1", ProjectId: 2, Priority: 1})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestGetTask_Success(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	task := &ent.Task{ID: 1, Title: "task1", ProjectID: 2, Priority: 1, Done: false}
	repo.On("GetByID", ctx, 1).Return(task, nil)
	resp, err := svc.GetTask(ctx, &pb.GetTaskRequest{Id: 1})
	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, "task1", resp.Title)
	assert.Equal(t, int32(2), resp.ProjectId)
	assert.Equal(t, int32(1), resp.Priority)
	assert.False(t, resp.Done)
}

func TestGetTask_Error(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("GetByID", ctx, 1).Return((*ent.Task)(nil), errors.New("not found"))
	resp, err := svc.GetTask(ctx, &pb.GetTaskRequest{Id: 1})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUpdateTask_Success(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	task := &ent.Task{ID: 1, Title: "task1", ProjectID: 2, Priority: 1, Done: true}
	repo.On("Update", ctx, 1, "task1", true, 1, 2).Return(task, nil)
	resp, err := svc.UpdateTask(ctx, &pb.UpdateTaskRequest{Id: 1, Title: "task1", Done: true, Priority: 1, ProjectId: 2})
	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, "task1", resp.Title)
	assert.True(t, resp.Done)
}

func TestUpdateTask_Error(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Update", ctx, 1, "task1", true, 1, 2).Return((*ent.Task)(nil), errors.New("fail"))
	resp, err := svc.UpdateTask(ctx, &pb.UpdateTaskRequest{Id: 1, Title: "task1", Done: true, Priority: 1, ProjectId: 2})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestDeleteTask_Success(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Delete", ctx, 1).Return(nil)
	resp, err := svc.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: 1})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
}

func TestDeleteTask_Error(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Delete", ctx, 1).Return(errors.New("fail"))
	resp, err := svc.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: 1})
	assert.Error(t, err)
	assert.False(t, resp.Success)
}

func TestListTasks_Success(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	tasks := []*ent.Task{{ID: 1, Title: "task1", ProjectID: 2, Priority: 1, Done: false}, {ID: 2, Title: "task2", ProjectID: 3, Priority: 2, Done: true}}
	repo.On("List", ctx).Return(tasks, nil)
	resp, err := svc.ListTasks(ctx, &pb.ListTasksRequest{})
	assert.NoError(t, err)
	assert.Len(t, resp.Tasks, 2)
	assert.Equal(t, int32(1), resp.Tasks[0].Id)
	assert.Equal(t, "task1", resp.Tasks[0].Title)
	assert.Equal(t, int32(2), resp.Tasks[0].ProjectId)
	assert.Equal(t, int32(1), resp.Tasks[0].Priority)
	assert.False(t, resp.Tasks[0].Done)
	assert.Equal(t, int32(2), resp.Tasks[1].Id)
	assert.Equal(t, "task2", resp.Tasks[1].Title)
	assert.Equal(t, int32(3), resp.Tasks[1].ProjectId)
	assert.Equal(t, int32(2), resp.Tasks[1].Priority)
	assert.True(t, resp.Tasks[1].Done)
}

func TestListTasks_Error(t *testing.T) {
	repo := new(MockTaskRepository)
	svc := NewTaskServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("List", ctx).Return([]*ent.Task{}, errors.New("fail"))
	resp, err := svc.ListTasks(ctx, &pb.ListTasksRequest{})
	assert.Error(t, err)
	assert.Nil(t, resp)
}
