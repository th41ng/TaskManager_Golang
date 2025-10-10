package service

import (
	"context"
	"taskmanager/microservices/task-service/ent"
)

type TaskRepository interface {
	Create(ctx context.Context, title string, projectID, priority int) (*ent.Task, error)
	GetByID(ctx context.Context, id int) (*ent.Task, error)
	Update(ctx context.Context, id int, title string, done bool, priority, projectID int) (*ent.Task, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*ent.Task, error)
}
