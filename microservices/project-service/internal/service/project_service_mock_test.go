package service

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"taskmanager/microservices/project-service/ent"
// 	"taskmanager/microservices/project-service/pb"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockProjectRepository struct {
// 	mock.Mock
// }

// func (m *MockProjectRepository) Create(ctx context.Context, name string, ownerID int) (*ent.Project, error) {
// 	args := m.Called(ctx, name, ownerID)
// 	return args.Get(0).(*ent.Project), args.Error(1)
// }
// func (m *MockProjectRepository) GetByID(ctx context.Context, id int) (*ent.Project, error) {
// 	args := m.Called(ctx, id)
// 	return args.Get(0).(*ent.Project), args.Error(1)
// }
// func (m *MockProjectRepository) Update(ctx context.Context, id int, name string, ownerID int) (*ent.Project, error) {
// 	args := m.Called(ctx, id, name, ownerID)
// 	return args.Get(0).(*ent.Project), args.Error(1)
// }
// func (m *MockProjectRepository) Delete(ctx context.Context, id int) error {
// 	args := m.Called(ctx, id)
// 	return args.Error(0)
// }
// func (m *MockProjectRepository) List(ctx context.Context) ([]*ent.Project, error) {
// 	args := m.Called(ctx)
// 	return args.Get(0).([]*ent.Project), args.Error(1)
// }

// type ProjectServiceWithRepo struct {
// 	repo ProjectRepository
// }

// func NewProjectServiceWithRepo(repo ProjectRepository) *ProjectServiceWithRepo {
// 	return &ProjectServiceWithRepo{repo: repo}
// }

// func (s *ProjectServiceWithRepo) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.ProjectResponse, error) {
// 	p, err := s.repo.Create(ctx, req.Name, int(req.OwnerId))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.ProjectResponse{Id: int32(p.ID), Name: p.Name, OwnerId: int32(p.OwnerID)}, nil
// }

// func (s *ProjectServiceWithRepo) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.ProjectResponse, error) {
// 	p, err := s.repo.GetByID(ctx, int(req.Id))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.ProjectResponse{Id: int32(p.ID), Name: p.Name, OwnerId: int32(p.OwnerID)}, nil
// }

// func (s *ProjectServiceWithRepo) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.ProjectResponse, error) {
// 	p, err := s.repo.Update(ctx, int(req.Id), req.Name, int(req.OwnerId))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.ProjectResponse{Id: int32(p.ID), Name: p.Name, OwnerId: int32(p.OwnerID)}, nil
// }

// func (s *ProjectServiceWithRepo) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
// 	err := s.repo.Delete(ctx, int(req.Id))
// 	if err != nil {
// 		return &pb.DeleteProjectResponse{Success: false}, err
// 	}
// 	return &pb.DeleteProjectResponse{Success: true}, nil
// }

// func (s *ProjectServiceWithRepo) ListProjects(ctx context.Context, req *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
// 	projects, err := s.repo.List(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp := &pb.ListProjectsResponse{}
// 	for _, p := range projects {
// 		resp.Projects = append(resp.Projects, &pb.ProjectResponse{
// 			Id: int32(p.ID), Name: p.Name, OwnerId: int32(p.OwnerID),
// 		})
// 	}
// 	return resp, nil
// }

// func TestCreateProject_Success(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	project := &ent.Project{ID: 1, Name: "proj1", OwnerID: 2}
// 	repo.On("Create", ctx, "proj1", 2).Return(project, nil)
// 	resp, err := svc.CreateProject(ctx, &pb.CreateProjectRequest{Name: "proj1", OwnerId: 2})
// 	assert.NoError(t, err)
// 	assert.Equal(t, int32(1), resp.Id)
// 	assert.Equal(t, "proj1", resp.Name)
// 	assert.Equal(t, int32(2), resp.OwnerId)
// }

// func TestCreateProject_Error(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	repo.On("Create", ctx, "proj1", 2).Return((*ent.Project)(nil), errors.New("fail"))
// 	resp, err := svc.CreateProject(ctx, &pb.CreateProjectRequest{Name: "proj1", OwnerId: 2})
// 	assert.Error(t, err)
// 	assert.Nil(t, resp)
// }

// func TestGetProject_Success(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	project := &ent.Project{ID: 1, Name: "proj1", OwnerID: 2}
// 	repo.On("GetByID", ctx, 1).Return(project, nil)
// 	resp, err := svc.GetProject(ctx, &pb.GetProjectRequest{Id: 1})
// 	assert.NoError(t, err)
// 	assert.Equal(t, int32(1), resp.Id)
// 	assert.Equal(t, "proj1", resp.Name)
// 	assert.Equal(t, int32(2), resp.OwnerId)
// }

// func TestGetProject_Error(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	repo.On("GetByID", ctx, 1).Return((*ent.Project)(nil), errors.New("not found"))
// 	resp, err := svc.GetProject(ctx, &pb.GetProjectRequest{Id: 1})
// 	assert.Error(t, err)
// 	assert.Nil(t, resp)
// }

// func TestUpdateProject_Success(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	project := &ent.Project{ID: 1, Name: "proj1", OwnerID: 2}
// 	repo.On("Update", ctx, 1, "proj1", 2).Return(project, nil)
// 	resp, err := svc.UpdateProject(ctx, &pb.UpdateProjectRequest{Id: 1, Name: "proj1", OwnerId: 2})
// 	assert.NoError(t, err)
// 	assert.Equal(t, int32(1), resp.Id)
// 	assert.Equal(t, "proj1", resp.Name)
// 	assert.Equal(t, int32(2), resp.OwnerId)
// }

// func TestUpdateProject_Error(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	repo.On("Update", ctx, 1, "proj1", 2).Return((*ent.Project)(nil), errors.New("fail"))
// 	resp, err := svc.UpdateProject(ctx, &pb.UpdateProjectRequest{Id: 1, Name: "proj1", OwnerId: 2})
// 	assert.Error(t, err)
// 	assert.Nil(t, resp)
// }

// func TestDeleteProject_Success(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	repo.On("Delete", ctx, 1).Return(nil)
// 	resp, err := svc.DeleteProject(ctx, &pb.DeleteProjectRequest{Id: 1})
// 	assert.NoError(t, err)
// 	assert.True(t, resp.Success)
// }

// func TestDeleteProject_Error(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	repo.On("Delete", ctx, 1).Return(errors.New("fail"))
// 	resp, err := svc.DeleteProject(ctx, &pb.DeleteProjectRequest{Id: 1})
// 	assert.Error(t, err)
// 	assert.False(t, resp.Success)
// }

// func TestListProjects_Success(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	projects := []*ent.Project{{ID: 1, Name: "proj1", OwnerID: 2}, {ID: 2, Name: "proj2", OwnerID: 3}}
// 	repo.On("List", ctx).Return(projects, nil)
// 	resp, err := svc.ListProjects(ctx, &pb.ListProjectsRequest{})
// 	assert.NoError(t, err)
// 	assert.Len(t, resp.Projects, 2)
// 	assert.Equal(t, int32(1), resp.Projects[0].Id)
// 	assert.Equal(t, "proj1", resp.Projects[0].Name)
// 	assert.Equal(t, int32(2), resp.Projects[0].OwnerId)
// 	assert.Equal(t, int32(2), resp.Projects[1].Id)
// 	assert.Equal(t, "proj2", resp.Projects[1].Name)
// 	assert.Equal(t, int32(3), resp.Projects[1].OwnerId)
// }

// func TestListProjects_Error(t *testing.T) {
// 	repo := new(MockProjectRepository)
// 	svc := NewProjectServiceWithRepo(repo)
// 	ctx := context.Background()
// 	repo.On("List", ctx).Return([]*ent.Project{}, errors.New("fail"))
// 	resp, err := svc.ListProjects(ctx, &pb.ListProjectsRequest{})
// 	assert.Error(t, err)
// 	assert.Nil(t, resp)
// }
