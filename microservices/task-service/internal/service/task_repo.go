package service

import (
	"context"
	"taskmanager/internal/commonrepo"
	"taskmanager/microservices/task-service/ent"
	"taskmanager/microservices/task-service/ent/task"
)

func NewTaskRepo(client *ent.Client) *commonrepo.CRUDRepo[ent.Task, *ent.Client] {
	return &commonrepo.CRUDRepo[ent.Task, *ent.Client]{
		CreateFn: func(ctx context.Context, c *ent.Client, args ...any) (*ent.Task, error) {
			return c.Task.Create().SetTitle(args[0].(string)).SetProjectID(args[1].(int)).SetPriority(args[2].(int)).SetDone(false).Save(ctx)
		},
		GetFn: func(ctx context.Context, c *ent.Client, id int) (*ent.Task, error) {
			return c.Task.Query().Where(task.ID(id)).Only(ctx)
		},
		UpdateFn: func(ctx context.Context, c *ent.Client, id int, args ...any) (*ent.Task, error) {
			return c.Task.UpdateOneID(id).SetTitle(args[0].(string)).SetDone(args[1].(bool)).SetPriority(args[2].(int)).SetProjectID(args[3].(int)).Save(ctx)
		},
		DeleteFn: func(ctx context.Context, c *ent.Client, id int) error {
			return c.Task.DeleteOneID(id).Exec(ctx)
		},
		ListFn: func(ctx context.Context, c *ent.Client) ([]*ent.Task, error) {
			return c.Task.Query().All(ctx)
		},
		Client: client,
	}
}
