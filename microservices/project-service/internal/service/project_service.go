package service

import (
	"context"
	"log"

	"taskmanager/microservices/project-service/ent"
	"taskmanager/microservices/project-service/ent/project"
	pb "taskmanager/microservices/project-service/pb"
)

type ProjectService struct {
	pb.UnimplementedProjectServiceServer
	client *ent.Client
}

// Constructor
func NewProjectService(client *ent.Client) *ProjectService {
	return &ProjectService{client: client}
}

// CreateProject
func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.ProjectResponse, error) {
	p, err := s.client.Project.
		Create().
		SetName(req.Name).
		SetOwnerID(int(req.OwnerId)).
		Save(ctx)
	if err != nil {
		log.Printf("failed to create project: %v", err)
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
	p, err := s.client.Project.
		Query().
		Where(project.ID(int(req.Id))).
		Only(ctx)
	if err != nil {
		log.Printf("project %d not found: %v", req.Id, err)
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
	p, err := s.client.Project.
		UpdateOneID(int(req.Id)).
		SetName(req.Name).
		SetOwnerID(int(req.OwnerId)).
		Save(ctx)
	if err != nil {
		log.Printf("failed to update project %d: %v", req.Id, err)
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
	err := s.client.Project.
		DeleteOneID(int(req.Id)).
		Exec(ctx)
	if err != nil {
		log.Printf("failed to delete project %d: %v", req.Id, err)
		return &pb.DeleteProjectResponse{Success: false}, err
	}

	return &pb.DeleteProjectResponse{Success: true}, nil
}

// ListProjects
func (s *ProjectService) ListProjects(ctx context.Context, req *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
	projects, err := s.client.Project.Query().All(ctx)
	if err != nil {
		log.Printf("failed to list projects: %v", err)
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
