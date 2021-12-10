package testutil

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/ent/enttest"
	"golang-clean-architecture-ent-gqlgen/pkg/infrastructure/datastore"
	"testing"

	"entgo.io/ent/dialect"
)

// NewDBClient loads database for test.
func NewDBClient(t *testing.T) *ent.Client {
	d := datastore.New()
	return enttest.Open(t, dialect.MySQL, d)
}

// DropAll drops all data from database.
func DropAll(t *testing.T, client *ent.Client) {
	t.Log("drop data from database")
	DropUser(t, client)
	DropTodo(t, client)
}

// DropUser drops data from users.
func DropUser(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.User.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// DropTodo drops data from todos.
func DropTodo(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.Todo.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
