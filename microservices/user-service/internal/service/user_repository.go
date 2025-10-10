package service

import (
	"context"
	"taskmanager/microservices/user-service/ent"
)

type UserRepository interface {
	Create(ctx context.Context, username, password string) (*ent.User, error)
	GetByID(ctx context.Context, id int) (*ent.User, error)
	GetByUsername(ctx context.Context, username string) (*ent.User, error)
	Update(ctx context.Context, id int, username, password string) (*ent.User, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*ent.User, error)
}
