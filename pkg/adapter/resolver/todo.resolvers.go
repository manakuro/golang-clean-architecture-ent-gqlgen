package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/ent/todo"
	"golang-clean-architecture-ent-gqlgen/graph/generated"
	"golang-clean-architecture-ent-gqlgen/pkg/util/datetime"
)

func (r *queryResolver) Todo(ctx context.Context, id *int) (*ent.Todo, error) {
	u, err := r.client.Todo.Query().Where(todo.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
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
