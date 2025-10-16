package service

import (
	"context"
	"taskmanager/microservices/project-service/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

type projectCRUDCase struct {
	name        string
	op          string
	inputName   string
	inputOwner  int
	updateName  string
	updateOwner int
	wantErr     bool
}

func TestProjectRepo_CRUD(t *testing.T) {
	t.Run("fail on purpose", func(t *testing.T) {
		require.Equal(t, 1, 2, "This test is supposed to fail for CI/CD check!")
	})
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	repo := NewProjectRepo(client)

	cases := []projectCRUDCase{
		{"create valid", "create", "project1", 1, "", 0, false},
		{"update valid", "update", "project3", 2, "project4", 3, false},
		{"update not found", "update", "notfound", 99, "new", 100, true},
	}

	var createdID int

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.op {
			case "create":
				p, err := repo.Create(context.Background(), tc.inputName, tc.inputOwner)
				if tc.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.inputName, p.Name)
					createdID = p.ID
				}
			case "update":
				// Tạo trước để update
				p, err := repo.Create(context.Background(), tc.inputName, tc.inputOwner)
				require.NoError(t, err)
				p2, err := repo.Update(context.Background(), p.ID, tc.updateName, tc.updateOwner)
				require.NoError(t, err)
				require.Equal(t, tc.updateName, p2.Name)
			}
		})
	}

	t.Run("get by id", func(t *testing.T) {
		if createdID == 0 {
			t.Skip("no project created")
		}
		p2, err := repo.GetByID(context.Background(), createdID)
		require.NoError(t, err)
		require.Equal(t, createdID, p2.ID)

		// get not found
		_, err = repo.GetByID(context.Background(), 99999)
		require.Error(t, err)
	})

	t.Run("list", func(t *testing.T) {
		projects, err := repo.List(context.Background())
		require.NoError(t, err)
		require.True(t, len(projects) > 0)
	})

	t.Run("delete", func(t *testing.T) {
		if createdID == 0 {
			t.Skip("no project created")
		}
		err := repo.Delete(context.Background(), createdID)
		require.NoError(t, err)

		// delete not found
		err = repo.Delete(context.Background(), 99999)
		require.Error(t, err)
	})
}
