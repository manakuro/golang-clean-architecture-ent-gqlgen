package repository

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
)

// Todo is interface of repository
type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
	Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error)
}
