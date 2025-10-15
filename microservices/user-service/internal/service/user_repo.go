package service

import (
	"context"
	"taskmanager/internal/commonrepo"
	"taskmanager/microservices/user-service/ent"
)

func NewUserRepo(client *ent.Client) *commonrepo.CRUDRepo[ent.User, *ent.Client] {
	return &commonrepo.CRUDRepo[ent.User, *ent.Client]{
		CreateFn: func(ctx context.Context, c *ent.Client, args ...any) (*ent.User, error) {
			return c.User.Create().SetUsername(args[0].(string)).SetPassword(args[1].(string)).Save(ctx)
		},
		GetFn: func(ctx context.Context, c *ent.Client, id int) (*ent.User, error) {
			return c.User.Get(ctx, id)
		},
		UpdateFn: func(ctx context.Context, c *ent.Client, id int, args ...any) (*ent.User, error) {
			return c.User.UpdateOneID(id).SetUsername(args[0].(string)).SetPassword(args[1].(string)).Save(ctx)
		},
		DeleteFn: func(ctx context.Context, c *ent.Client, id int) error {
			return c.User.DeleteOneID(id).Exec(ctx)
		},
		ListFn: func(ctx context.Context, c *ent.Client) ([]*ent.User, error) {
			return c.User.Query().All(ctx)
		},
		Client: client,
	}
}
