package service

import (
	"context"
	"errors"
	"taskmanager/microservices/project-service/ent"
)

type MockProjectRepo struct {
	data   map[int]*ent.Project
	nextID int
}

func NewMockRepo() *MockProjectRepo {
	return &MockProjectRepo{
		data:   make(map[int]*ent.Project),
		nextID: 1,
	}
}

func (m *MockProjectRepo) Create(ctx context.Context, name string, ownerID int) (*ent.Project, error) {
	p := &ent.Project{ID: m.nextID, Name: name, OwnerID: ownerID}
	m.data[m.nextID] = p
	m.nextID++
	return p, nil
}

func (m *MockProjectRepo) GetByID(ctx context.Context, id int) (*ent.Project, error) {
	if p, ok := m.data[id]; ok {
		return p, nil
	}
	return nil, errors.New("not found")
}

func (m *MockProjectRepo) Update(ctx context.Context, id int, name string, ownerID int) (*ent.Project, error) {
	if p, ok := m.data[id]; ok {
		p.Name = name
		p.OwnerID = ownerID
		return p, nil
	}
	return nil, errors.New("not found")
}

func (m *MockProjectRepo) Delete(ctx context.Context, id int) error {
	if _, ok := m.data[id]; ok {
		delete(m.data, id)
		return nil
	}
	return errors.New("not found")
}

func (m *MockProjectRepo) List(ctx context.Context) ([]*ent.Project, error) {
	res := []*ent.Project{}
	for _, p := range m.data {
		res = append(res, p)
	}
	return res, nil
}
