package repository

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/ent/todo"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
	ur "golang-clean-architecture-ent-gqlgen/pkg/usecase/repository"
)

type todoRepository struct {
	client *ent.Client
}

// NewTodoRepository generates  user repository
func NewTodoRepository(client *ent.Client) ur.Todo {
	return &todoRepository{client: client}
}

func (r *todoRepository) Get(ctx context.Context, id *model.ID) (*model.Todo, error) {
	q := r.client.Todo.Query()
	if id != nil {
		q.Where(todo.IDEQ(*id))
	}

	u, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"id": id,
			})
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *todoRepository) List(ctx context.Context) ([]*model.Todo, error) {
	ts, err := r.client.
		Todo.
		Query().
		All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return ts, nil
}

func (r *todoRepository) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	u, err := r.client.
		Todo.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *todoRepository) Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	u, err := r.client.
		Todo.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return u, nil
}
