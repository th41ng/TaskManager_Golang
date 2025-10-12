package service

import (
	"context"
	"testing"

	pb "taskmanager/microservices/project-service/pb"

	"github.com/stretchr/testify/assert"
)

func TestProjectService_CRUD(t *testing.T) {
	ctx := context.Background()
	repo := NewMockRepo()
	svc := NewProjectService(repo)

	// Create
	createResp, err := svc.CreateProject(ctx, &pb.CreateProjectRequest{Name: "Demo", OwnerId: 1})
	assert.NoError(t, err)
	assert.Equal(t, "Demo", createResp.Name)

	// Get
	getResp, err := svc.GetProject(ctx, &pb.GetProjectRequest{Id: createResp.Id})
	assert.NoError(t, err)
	assert.Equal(t, "Demo", getResp.Name)

	// Update
	updResp, err := svc.UpdateProject(ctx, &pb.UpdateProjectRequest{Id: createResp.Id, Name: "Updated", OwnerId: 2})
	assert.NoError(t, err)
	assert.Equal(t, "Updated", updResp.Name)

	// List
	listResp, err := svc.ListProjects(ctx, &pb.ListProjectsRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Projects, 1)

	// Delete
	delResp, err := svc.DeleteProject(ctx, &pb.DeleteProjectRequest{Id: createResp.Id})
	assert.NoError(t, err)
	assert.True(t, delResp.Success)

	// Get again (should fail)
	_, err = svc.GetProject(ctx, &pb.GetProjectRequest{Id: createResp.Id})
	assert.Error(t, err)
}
