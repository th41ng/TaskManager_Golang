package service

import (
	"context"
	"taskmanager/microservices/project-service/ent"
	"taskmanager/microservices/project-service/ent/project"
)

type EntProjectRepo struct {
	client *ent.Client
}

func NewEntProjectRepo(client *ent.Client) *EntProjectRepo {
	return &EntProjectRepo{client: client}
}

func (r *EntProjectRepo) Create(ctx context.Context, name string, ownerID int) (*ent.Project, error) {
	return r.client.Project.Create().SetName(name).SetOwnerID(ownerID).Save(ctx)
}

func (r *EntProjectRepo) GetByID(ctx context.Context, id int) (*ent.Project, error) {
	return r.client.Project.Query().Where(project.ID(id)).Only(ctx)
}

func (r *EntProjectRepo) Update(ctx context.Context, id int, name string, ownerID int) (*ent.Project, error) {
	return r.client.Project.UpdateOneID(id).SetName(name).SetOwnerID(ownerID).Save(ctx)
}

func (r *EntProjectRepo) Delete(ctx context.Context, id int) error {
	return r.client.Project.DeleteOneID(id).Exec(ctx)
}

func (r *EntProjectRepo) List(ctx context.Context) ([]*ent.Project, error) {
	return r.client.Project.Query().All(ctx)
}
