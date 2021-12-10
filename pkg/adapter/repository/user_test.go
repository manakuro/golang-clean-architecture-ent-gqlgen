package repository_test

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/repository"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
	"golang-clean-architecture-ent-gqlgen/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (client *ent.Client, teardown func()) {
	testutil.ReadConfig()
	c := testutil.NewDBClient(t)

	return c, func() {
		testutil.DropUser(t, c)
		defer c.Close()
	}
}

func TestUserRepository_List(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewUserRepository(client)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (uc *model.UserConnection, err error)
		assert  func(t *testing.T, uc *model.UserConnection, err error)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name: "It should get user's list",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.User.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				users := []struct {
					name string
					age  int
				}{{name: "test", age: 10}, {name: "test2", age: 11}, {name: "test3", age: 12}}
				bulk := make([]*ent.UserCreate, len(users))
				for i, u := range users {
					bulk[i] = client.User.Create().SetName(u.name).SetAge(u.age)
				}

				_, err = client.User.
					CreateBulk(bulk...).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
			},
			act: func(ctx context.Context, t *testing.T) (us *model.UserConnection, err error) {
				first := 5
				return repo.List(ctx, nil, &first, nil, nil, nil)
			},
			assert: func(t *testing.T, got *model.UserConnection, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 3, len(got.Edges))
			},
			args: args{
				ctx: context.Background(),
			},
			teardown: func(t *testing.T) {
				testutil.DropUser(t, client)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
			tt.teardown(t)
		})
	}
}
