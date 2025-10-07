package service

import (
	"context"
	"log"

	"taskmanager/microservices/user-service/ent"
	"taskmanager/microservices/user-service/ent/user"
	pb "taskmanager/microservices/user-service/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{client: client}
}

// CreateUser
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	u, err := s.client.User.
		Create().
		SetUsername(req.Username).
		SetPassword(req.Password).
		Save(ctx)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
	}, nil
}

// GetUser
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	u, err := s.client.User.Query().
		Where(user.ID(int(req.Id))).
		Only(ctx)
	if err != nil {
		log.Printf("user %d not found: %v", req.Id, err)
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
	}, nil
}

// UpdateUser
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	u, err := s.client.User.
		UpdateOneID(int(req.Id)).
		SetUsername(req.Username).
		SetPassword(req.Password).
		Save(ctx)
	if err != nil {
		log.Printf("failed to update user %d: %v", req.Id, err)
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
	}, nil
}

// DeleteUser
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.client.User.DeleteOneID(int(req.Id)).Exec(ctx)
	if err != nil {
		log.Printf("failed to delete user %d: %v", req.Id, err)
		return &pb.DeleteUserResponse{Success: false}, err
	}

	return &pb.DeleteUserResponse{Success: true}, nil
}

// ListUsers
func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.client.User.Query().All(ctx)
	if err != nil {
		log.Printf("failed to list users: %v", err)
		return nil, err
	}

	resp := &pb.ListUsersResponse{}
	for _, u := range users {
		resp.Users = append(resp.Users, &pb.UserResponse{
			Id:       int32(u.ID),
			Username: u.Username,
		})
	}

	return resp, nil
}
