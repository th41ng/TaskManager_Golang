package service

import (
	"context"
	"taskmanager/microservices/task-service/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

type taskCRUDCase struct {
	name        string
	op          string
	inputTitle  string
	inputPID    int
	inputUID    int
	updateTitle string
	updateDone  bool
	updatePID   int
	updateUID   int
	wantErr     bool
}

func TestTaskRepo_CRUD(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	repo := NewTaskRepo(client)

	cases := []taskCRUDCase{
		{"create valid", "create", "task1", 1, 2, "", false, 0, 0, false},
		{"update valid", "update", "task2", 2, 3, "task3", true, 4, 5, false},
		{"update not found", "update", "notfound", 99, 100, "new", true, 101, 102, true},
	}

	var createdID int

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.op {
			case "create":
				taskObj, err := repo.Create(context.Background(), tc.inputTitle, tc.inputPID, tc.inputUID)
				if tc.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.inputTitle, taskObj.Title)
					createdID = taskObj.ID
				}
			case "update":
				// Create first to update
				taskObj, err := repo.Create(context.Background(), tc.inputTitle, tc.inputPID, tc.inputUID)
				require.NoError(t, err)
				t2, err := repo.Update(context.Background(), taskObj.ID, tc.updateTitle, tc.updateDone, tc.updatePID, tc.updateUID)
				require.NoError(t, err)
				require.Equal(t, tc.updateTitle, t2.Title)
			}
		})
	}

	t.Run("get by id", func(t *testing.T) {
		if createdID == 0 {
			t.Skip("no task created")
		}
		t2, err := repo.GetByID(context.Background(), createdID)
		require.NoError(t, err)
		require.Equal(t, createdID, t2.ID)

		// get not found
		_, err = repo.GetByID(context.Background(), 99999)
		require.Error(t, err)
	})

	t.Run("list", func(t *testing.T) {
		tasks, err := repo.List(context.Background())
		require.NoError(t, err)
		require.True(t, len(tasks) > 0)
	})

	t.Run("delete", func(t *testing.T) {
		if createdID == 0 {
			t.Skip("no task created")
		}
		err := repo.Delete(context.Background(), createdID)
		require.NoError(t, err)

		// delete not found
		err = repo.Delete(context.Background(), 99999)
		require.Error(t, err)
	})
}
