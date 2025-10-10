package service

import (
	"context"
	"errors"
	"testing"

	"taskmanager/microservices/user-service/ent"
	"taskmanager/microservices/user-service/pb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, username, password string) (*ent.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(*ent.User), args.Error(1)
}
func (m *MockUserRepository) GetByID(ctx context.Context, id int) (*ent.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*ent.User), args.Error(1)
}
func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*ent.User), args.Error(1)
}
func (m *MockUserRepository) Update(ctx context.Context, id int, username, password string) (*ent.User, error) {
	args := m.Called(ctx, id, username, password)
	return args.Get(0).(*ent.User), args.Error(1)
}
func (m *MockUserRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockUserRepository) List(ctx context.Context) ([]*ent.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*ent.User), args.Error(1)
}

type UserServiceWithRepo struct {
	repo UserRepository
}

func NewUserServiceWithRepo(repo UserRepository) *UserServiceWithRepo {
	return &UserServiceWithRepo{repo: repo}
}

func (s *UserServiceWithRepo) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	hashed := req.Password
	u, err := s.repo.Create(ctx, req.Username, hashed)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: int32(u.ID), Username: u.Username}, nil
}

func (s *UserServiceWithRepo) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	u, err := s.repo.GetByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: int32(u.ID), Username: u.Username}, nil
}

func (s *UserServiceWithRepo) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	u, err := s.repo.Update(ctx, int(req.Id), req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: int32(u.ID), Username: u.Username}, nil
}

func (s *UserServiceWithRepo) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.repo.Delete(ctx, int(req.Id))
	if err != nil {
		return &pb.DeleteUserResponse{Success: false}, err
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}

func (s *UserServiceWithRepo) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListUsersResponse{}
	for _, u := range users {
		resp.Users = append(resp.Users, &pb.UserResponse{
			Id: int32(u.ID), Username: u.Username,
		})
	}
	return resp, nil
}

func TestCreateUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	user := &ent.User{ID: 1, Username: "testuser", Password: "hashed"}
	repo.On("Create", ctx, "testuser", "password").Return(user, nil)
	resp, err := svc.CreateUser(ctx, &pb.CreateUserRequest{Username: "testuser", Password: "password"})
	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, "testuser", resp.Username)
}

func TestCreateUser_Error(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Create", ctx, "testuser", "password").Return((*ent.User)(nil), errors.New("fail"))
	resp, err := svc.CreateUser(ctx, &pb.CreateUserRequest{Username: "testuser", Password: "password"})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestGetUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	user := &ent.User{ID: 1, Username: "testuser"}
	repo.On("GetByID", ctx, 1).Return(user, nil)
	resp, err := svc.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, "testuser", resp.Username)
}

func TestGetUser_Error(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("GetByID", ctx, 1).Return((*ent.User)(nil), errors.New("not found"))
	resp, err := svc.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUpdateUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	user := &ent.User{ID: 1, Username: "testuser"}
	repo.On("Update", ctx, 1, "testuser", "newpass").Return(user, nil)
	resp, err := svc.UpdateUser(ctx, &pb.UpdateUserRequest{Id: 1, Username: "testuser", Password: "newpass"})
	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.Id)
	assert.Equal(t, "testuser", resp.Username)
}

func TestUpdateUser_Error(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Update", ctx, 1, "testuser", "newpass").Return((*ent.User)(nil), errors.New("fail"))
	resp, err := svc.UpdateUser(ctx, &pb.UpdateUserRequest{Id: 1, Username: "testuser", Password: "newpass"})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestDeleteUser_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Delete", ctx, 1).Return(nil)
	resp, err := svc.DeleteUser(ctx, &pb.DeleteUserRequest{Id: 1})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
}

func TestDeleteUser_Error(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("Delete", ctx, 1).Return(errors.New("fail"))
	resp, err := svc.DeleteUser(ctx, &pb.DeleteUserRequest{Id: 1})
	assert.Error(t, err)
	assert.False(t, resp.Success)
}

func TestListUsers_Success(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	users := []*ent.User{{ID: 1, Username: "user1"}, {ID: 2, Username: "user2"}}
	repo.On("List", ctx).Return(users, nil)
	resp, err := svc.ListUsers(ctx, &pb.ListUsersRequest{})
	assert.NoError(t, err)
	assert.Len(t, resp.Users, 2)
	assert.Equal(t, int32(1), resp.Users[0].Id)
	assert.Equal(t, "user1", resp.Users[0].Username)
	assert.Equal(t, int32(2), resp.Users[1].Id)
	assert.Equal(t, "user2", resp.Users[1].Username)
}

func TestListUsers_Error(t *testing.T) {
	repo := new(MockUserRepository)
	svc := NewUserServiceWithRepo(repo)
	ctx := context.Background()
	repo.On("List", ctx).Return([]*ent.User{}, errors.New("fail"))
	resp, err := svc.ListUsers(ctx, &pb.ListUsersRequest{})
	assert.Error(t, err)
	assert.Nil(t, resp)
}
