package service

import (
	"context"
	pb "taskmanager/microservices/project-service/pb"
)

type ProjectService struct {
	pb.UnimplementedProjectServiceServer
	repo ProjectRepo
}

func NewProjectService(repo ProjectRepo) *ProjectService {
	return &ProjectService{repo: repo}
}

// CreateProject
func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.ProjectResponse, error) {
	p, err := s.repo.Create(ctx, req.Name, int(req.OwnerId))
	if err != nil {
		return nil, err
	}

	return &pb.ProjectResponse{
		Id:      int32(p.ID),
		Name:    p.Name,
		OwnerId: int32(p.OwnerID),
	}, nil
}

// GetProject
func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.ProjectResponse, error) {
	p, err := s.repo.GetByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ProjectResponse{
		Id:      int32(p.ID),
		Name:    p.Name,
		OwnerId: int32(p.OwnerID),
	}, nil
}

// UpdateProject
func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.ProjectResponse, error) {
	p, err := s.repo.Update(ctx, int(req.Id), req.Name, int(req.OwnerId))
	if err != nil {
		return nil, err
	}

	return &pb.ProjectResponse{
		Id:      int32(p.ID),
		Name:    p.Name,
		OwnerId: int32(p.OwnerID),
	}, nil
}

// DeleteProject
func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	err := s.repo.Delete(ctx, int(req.Id))
	if err != nil {
		return &pb.DeleteProjectResponse{Success: false}, err
	}
	return &pb.DeleteProjectResponse{Success: true}, nil
}

// ListProjects
func (s *ProjectService) ListProjects(ctx context.Context, req *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
	projects, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListProjectsResponse{}
	for _, p := range projects {
		resp.Projects = append(resp.Projects, &pb.ProjectResponse{
			Id:      int32(p.ID),
			Name:    p.Name,
			OwnerId: int32(p.OwnerID),
		})
	}
	return resp, nil
}
