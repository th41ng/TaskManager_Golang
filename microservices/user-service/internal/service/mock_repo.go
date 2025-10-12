package service

import (
	"context"
	"errors"
	"taskmanager/microservices/user-service/ent"
)

type MockUserRepo struct {
	users []*ent.User
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{users: []*ent.User{}}
}

func (m *MockUserRepo) CreateUser(ctx context.Context, username, password string) (*ent.User, error) {
	u := &ent.User{ID: len(m.users) + 1, Username: username, Password: password}
	m.users = append(m.users, u)
	return u, nil
}

func (m *MockUserRepo) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	for _, u := range m.users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockUserRepo) GetByID(ctx context.Context, id int) (*ent.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockUserRepo) UpdateUser(ctx context.Context, id int, username, password string) (*ent.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			u.Username = username
			u.Password = password
			return u, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockUserRepo) DeleteUser(ctx context.Context, id int) error {
	for i, u := range m.users {
		if u.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (m *MockUserRepo) ListUsers(ctx context.Context) ([]*ent.User, error) {
	return m.users, nil
}
