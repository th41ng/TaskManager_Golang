package service

import (
	"context"
	"taskmanager/microservices/user-service/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

type userCRUDCase struct {
	name       string
	op         string
	inputName  string
	inputPass  string
	updateName string
	updatePass string
	wantErr    bool
}

func TestUserRepo_CRUD(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	repo := NewUserRepo(client)

	cases := []userCRUDCase{
		{"create valid", "create", "testuser", "pass", "", "", false},
		{"update valid", "update", "user2", "pass2", "user3", "pass3", false},
		{"update not found", "update", "notfound", "pass", "new", "pass2", true},
	}

	var createdID int

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.op {
			case "create":
				u, err := repo.Create(context.Background(), tc.inputName, tc.inputPass)
				if tc.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.inputName, u.Username)
					createdID = u.ID
				}
			case "update":
				// Create first to update
				u, err := repo.Create(context.Background(), tc.inputName, tc.inputPass)
				require.NoError(t, err)
				u2, err := repo.Update(context.Background(), u.ID, tc.updateName, tc.updatePass)
				require.NoError(t, err)
				require.Equal(t, tc.updateName, u2.Username)
			}
		})
	}

	t.Run("get by id", func(t *testing.T) {
		if createdID == 0 {
			t.Skip("no user created")
		}
		u2, err := repo.GetByID(context.Background(), createdID)
		require.NoError(t, err)
		require.Equal(t, createdID, u2.ID)

		// get not found
		_, err = repo.GetByID(context.Background(), 99999)
		require.Error(t, err)
	})

	t.Run("list", func(t *testing.T) {
		users, err := repo.List(context.Background())
		require.NoError(t, err)
		require.True(t, len(users) > 0)
	})

	t.Run("delete", func(t *testing.T) {
		if createdID == 0 {
			t.Skip("no user created")
		}
		err := repo.Delete(context.Background(), createdID)
		require.NoError(t, err)

		// delete not found
		err = repo.Delete(context.Background(), 99999)
		require.Error(t, err)
	})
}
