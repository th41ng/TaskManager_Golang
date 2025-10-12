package service

import (
	"context"

	"taskmanager/microservices/task-service/ent"
)

// TaskRepo định nghĩa các hàm mà repository thật & mock đều phải có
type TaskRepo interface {
	CreateTask(ctx context.Context, title string, projectID, priority int) (*ent.Task, error)
	GetTask(ctx context.Context, id int) (*ent.Task, error)
	UpdateTask(ctx context.Context, id int, title string, done bool, priority, projectID int) (*ent.Task, error)
	DeleteTask(ctx context.Context, id int) error
	ListTasks(ctx context.Context) ([]*ent.Task, error)
}
