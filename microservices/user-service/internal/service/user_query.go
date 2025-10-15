package service

import (
	"context"
	"taskmanager/microservices/user-service/ent"
	"taskmanager/microservices/user-service/ent/user"
)

type UserQuery interface {
	GetByUsername(ctx context.Context, username string) (*ent.User, error)
}

type userQuery struct {
	client *ent.Client
}

func NewUserQuery(client *ent.Client) UserQuery {
	return &userQuery{client: client}
}

func (q *userQuery) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return q.client.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
}
