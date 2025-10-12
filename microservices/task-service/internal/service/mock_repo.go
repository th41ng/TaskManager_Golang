package service

import (
	"context"
	"errors"

	"taskmanager/microservices/task-service/ent"
)

type MockTaskRepo struct {
	tasks  map[int]*ent.Task
	nextID int
}

func NewMockTaskRepo() *MockTaskRepo {
	return &MockTaskRepo{
		tasks:  make(map[int]*ent.Task),
		nextID: 1,
	}
}

func (m *MockTaskRepo) CreateTask(ctx context.Context, title string, projectID, priority int) (*ent.Task, error) {
	t := &ent.Task{ID: m.nextID, Title: title, ProjectID: projectID, Priority: priority, Done: false}
	m.tasks[m.nextID] = t
	m.nextID++
	return t, nil
}

func (m *MockTaskRepo) GetTask(ctx context.Context, id int) (*ent.Task, error) {
	if t, ok := m.tasks[id]; ok {
		return t, nil
	}
	return nil, errors.New("not found")
}

func (m *MockTaskRepo) UpdateTask(ctx context.Context, id int, title string, done bool, priority, projectID int) (*ent.Task, error) {
	if t, ok := m.tasks[id]; ok {
		t.Title = title
		t.Done = done
		t.Priority = priority
		t.ProjectID = projectID
		return t, nil
	}
	return nil, errors.New("not found")
}

func (m *MockTaskRepo) DeleteTask(ctx context.Context, id int) error {
	delete(m.tasks, id)
	return nil
}

func (m *MockTaskRepo) ListTasks(ctx context.Context) ([]*ent.Task, error) {
	list := []*ent.Task{}
	for _, t := range m.tasks {
		list = append(list, t)
	}
	return list, nil
}
