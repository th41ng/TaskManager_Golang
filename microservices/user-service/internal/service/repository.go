package service

import (
	"context"
	"taskmanager/microservices/user-service/ent"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username, password string) (*ent.User, error)
	GetByUsername(ctx context.Context, username string) (*ent.User, error)
	GetByID(ctx context.Context, id int) (*ent.User, error)
	UpdateUser(ctx context.Context, id int, username, password string) (*ent.User, error)
	DeleteUser(ctx context.Context, id int) error
	ListUsers(ctx context.Context) ([]*ent.User, error)
}
