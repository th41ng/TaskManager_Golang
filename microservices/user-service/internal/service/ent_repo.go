package service

import (
	"context"
	"taskmanager/microservices/user-service/ent"
	"taskmanager/microservices/user-service/ent/user"
)

type EntUserRepo struct {
	client *ent.Client
}

func NewEntUserRepo(client *ent.Client) *EntUserRepo {
	return &EntUserRepo{client: client}
}

func (r *EntUserRepo) CreateUser(ctx context.Context, username, password string) (*ent.User, error) {
	return r.client.User.Create().
		SetUsername(username).
		SetPassword(password).
		Save(ctx)
}

func (r *EntUserRepo) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return r.client.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
}

func (r *EntUserRepo) GetByID(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
}

func (r *EntUserRepo) UpdateUser(ctx context.Context, id int, username, password string) (*ent.User, error) {
	return r.client.User.UpdateOneID(id).
		SetUsername(username).
		SetPassword(password).
		Save(ctx)
}

func (r *EntUserRepo) DeleteUser(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}

func (r *EntUserRepo) ListUsers(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}
