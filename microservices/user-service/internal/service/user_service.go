package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"taskmanager/internal/commonrepo"
	"taskmanager/microservices/user-service/ent"
	pb "taskmanager/microservices/user-service/pb"
)

var jwtKey = []byte("12345678901234567890123456789012")

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo  *commonrepo.CRUDRepo[ent.User, *ent.Client]
	query UserQuery
}

func NewUserService(repo *commonrepo.CRUDRepo[ent.User, *ent.Client], query UserQuery) *UserService {
	return &UserService{repo: repo, query: query}
}

func generateJWT(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ✅ 1. CreateUser
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u, err := s.repo.Create(ctx, req.Username, string(hashed))
	if err != nil {
		return nil, err
	}

	token, _ := generateJWT(u.ID, u.Username)

	return &pb.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
		Token:    token,
	}, nil
}

// ✅ 2. Login
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	u, err := s.query.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		return nil, errors.New("invalid username or password")
	}

	token, _ := generateJWT(u.ID, u.Username)
	return &pb.LoginResponse{Token: token}, nil
}

// ✅ 3. GetUser
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	u, err := s.repo.GetByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
	}, nil
}

// ✅ 4. UpdateUser
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u, err := s.repo.Update(ctx, int(req.Id), req.Username, string(hashed))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
	}, nil
}

// ✅ 5. DeleteUser
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.repo.Delete(ctx, int(req.Id))
	if err != nil {
		return &pb.DeleteUserResponse{Success: false}, err
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}

// ✅ 6. ListUsers
func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
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
