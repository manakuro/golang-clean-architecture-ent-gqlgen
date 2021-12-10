package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/ent/schema/ulid"
	"golang-clean-architecture-ent-gqlgen/graph/generated"
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/handler"
	"golang-clean-architecture-ent-gqlgen/pkg/util/datetime"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	t, err := r.controller.Todo.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input ent.UpdateTodoInput) (*ent.Todo, error) {
	t, err := r.controller.Todo.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) Todo(ctx context.Context, id *ulid.ID) (*ent.Todo, error) {
	t, err := r.controller.Todo.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	ts, err := r.controller.Todo.List(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return ts, nil
}

func (r *todoResolver) CreatedAt(ctx context.Context, obj *ent.Todo) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *todoResolver) UpdatedAt(ctx context.Context, obj *ent.Todo) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
