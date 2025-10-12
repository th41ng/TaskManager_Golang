package service

import (
	"context"
	"testing"

	pb "taskmanager/microservices/task-service/pb"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_CRUD(t *testing.T) {
	repo := NewMockTaskRepo()
	svc := NewTaskService(repo)
	ctx := context.Background()

	// Create
	createResp, err := svc.CreateTask(ctx, &pb.CreateTaskRequest{
		Title: "Task 1", ProjectId: 10, Priority: 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, "Task 1", createResp.Title)

	// Get
	getResp, err := svc.GetTask(ctx, &pb.GetTaskRequest{Id: createResp.Id})
	assert.NoError(t, err)
	assert.Equal(t, "Task 1", getResp.Title)

	// Update
	updResp, err := svc.UpdateTask(ctx, &pb.UpdateTaskRequest{
		Id: createResp.Id, Title: "Updated", Done: true, Priority: 2, ProjectId: 11,
	})
	assert.NoError(t, err)
	assert.Equal(t, "Updated", updResp.Title)
	assert.True(t, updResp.Done)

	// List
	listResp, err := svc.ListTasks(ctx, &pb.ListTasksRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Tasks, 1)

	// Delete
	delResp, err := svc.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: createResp.Id})
	assert.NoError(t, err)
	assert.True(t, delResp.Success)
}
