package service

import (
	"context"
	"testing"

	"taskmanager/microservices/project-service/ent"
	"taskmanager/microservices/project-service/pb"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestProjectService_CRUD(t *testing.T) {
	ctx := context.Background()

	// ⚠️ Dòng này rất quan trọng:
	client, err := ent.Open("sqlite", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		t.Fatalf("failed opening sqlite db: %v", err)
	}
	defer client.Close()

	// Auto migrate schema cho DB trong bộ nhớ
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	svc := NewProjectService(client)

	// === Create ===
	resp, err := svc.CreateProject(ctx, &pb.CreateProjectRequest{Name: "proj1", OwnerId: 2})
	assert.NoError(t, err)
	assert.Equal(t, "proj1", resp.Name)
	assert.Equal(t, int32(2), resp.OwnerId)

	// === Get ===
	getResp, err := svc.GetProject(ctx, &pb.GetProjectRequest{Id: resp.Id})
	assert.NoError(t, err)
	assert.Equal(t, resp.Id, getResp.Id)
	assert.Equal(t, "proj1", getResp.Name)

	// === Update ===
	updResp, err := svc.UpdateProject(ctx, &pb.UpdateProjectRequest{Id: resp.Id, Name: "proj1-upd", OwnerId: 3})
	assert.NoError(t, err)
	assert.Equal(t, "proj1-upd", updResp.Name)
	assert.Equal(t, int32(3), updResp.OwnerId)

	// === List ===
	listResp, err := svc.ListProjects(ctx, &pb.ListProjectsRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Projects, 1)
	assert.Equal(t, "proj1-upd", listResp.Projects[0].Name)

	// === Delete ===
	delResp, err := svc.DeleteProject(ctx, &pb.DeleteProjectRequest{Id: resp.Id})
	assert.NoError(t, err)
	assert.True(t, delResp.Success)

	// === Get after delete ===
	_, err = svc.GetProject(ctx, &pb.GetProjectRequest{Id: resp.Id})
	assert.Error(t, err)
}
