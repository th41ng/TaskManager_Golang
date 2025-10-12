package service

import (
	"context"

	"taskmanager/microservices/task-service/ent"
	"taskmanager/microservices/task-service/ent/task"
)

type EntTaskRepo struct {
	client *ent.Client
}

func NewEntTaskRepo(client *ent.Client) *EntTaskRepo {
	return &EntTaskRepo{client: client}
}

func (r *EntTaskRepo) CreateTask(ctx context.Context, title string, projectID, priority int) (*ent.Task, error) {
	return r.client.Task.Create().
		SetTitle(title).
		SetProjectID(projectID).
		SetPriority(priority).
		SetDone(false).
		Save(ctx)
}

func (r *EntTaskRepo) GetTask(ctx context.Context, id int) (*ent.Task, error) {
	return r.client.Task.Query().
		Where(task.ID(id)).
		Only(ctx)
}

func (r *EntTaskRepo) UpdateTask(ctx context.Context, id int, title string, done bool, priority, projectID int) (*ent.Task, error) {
	return r.client.Task.UpdateOneID(id).
		SetTitle(title).
		SetDone(done).
		SetPriority(priority).
		SetProjectID(projectID).
		Save(ctx)
}

func (r *EntTaskRepo) DeleteTask(ctx context.Context, id int) error {
	return r.client.Task.DeleteOneID(id).Exec(ctx)
}

func (r *EntTaskRepo) ListTasks(ctx context.Context) ([]*ent.Task, error) {
	return r.client.Task.Query().All(ctx)
}
