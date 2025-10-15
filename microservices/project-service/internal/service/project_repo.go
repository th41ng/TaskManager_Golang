package service

import (
	"context"
	"taskmanager/internal/commonrepo"
	"taskmanager/microservices/project-service/ent"
	"taskmanager/microservices/project-service/ent/project"
)

func NewProjectRepo(client *ent.Client) *commonrepo.CRUDRepo[ent.Project, *ent.Client] {
	return &commonrepo.CRUDRepo[ent.Project, *ent.Client]{
		CreateFn: func(ctx context.Context, c *ent.Client, args ...any) (*ent.Project, error) {
			return c.Project.Create().SetName(args[0].(string)).SetOwnerID(args[1].(int)).Save(ctx)
		},
		GetFn: func(ctx context.Context, c *ent.Client, id int) (*ent.Project, error) {
			return c.Project.Query().Where(project.ID(id)).Only(ctx)
		},
		UpdateFn: func(ctx context.Context, c *ent.Client, id int, args ...any) (*ent.Project, error) {
			return c.Project.UpdateOneID(id).SetName(args[0].(string)).SetOwnerID(args[1].(int)).Save(ctx)
		},
		DeleteFn: func(ctx context.Context, c *ent.Client, id int) error {
			return c.Project.DeleteOneID(id).Exec(ctx)
		},
		ListFn: func(ctx context.Context, c *ent.Client) ([]*ent.Project, error) {
			return c.Project.Query().All(ctx)
		},
		Client: client,
	}
}
