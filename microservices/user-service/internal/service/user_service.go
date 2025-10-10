package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"taskmanager/microservices/user-service/ent"
	"taskmanager/microservices/user-service/ent/user"
	pb "taskmanager/microservices/user-service/pb"
)

var jwtKey = []byte("12345678901234567890123456789012")

func generateJWT(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

type UserService struct {
	pb.UnimplementedUserServiceServer
	client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{client: client}
}

// CreateUser: hash password trước khi lưu
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return nil, err
	}

	u, err := s.client.User.
		Create().
		SetUsername(req.Username).
		SetPassword(string(hashedPassword)).
		Save(ctx)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return nil, err
	}

	token, err := generateJWT(u.ID, u.Username)
	if err != nil {
		log.Printf("failed to generate token: %v", err)
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(u.ID),
		Username: u.Username,
		Token:    token,
	}, nil
}

// Login: kiểm tra username + password, trả JWT
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	u, err := s.client.User.Query().
		Where(user.UsernameEQ(req.Username)).
		Only(ctx)
	if err != nil {
		log.Printf("login failed: user not found: %v", err)
		return nil, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		log.Printf("login failed: wrong password for user %s", req.Username)
		return nil, errors.New("invalid username or password")
	}

	token, err := generateJWT(u.ID, u.Username)
	if err != nil {
		log.Printf("failed to generate token: %v", err)
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
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

// UpdateUser: nếu password được cập nhật, hash trước khi lưu
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	update := s.client.User.UpdateOneID(int(req.Id)).
		SetUsername(req.Username)

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("failed to hash password: %v", err)
			return nil, err
		}
		update.SetPassword(string(hashedPassword))
	}

	u, err := update.Save(ctx)
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
