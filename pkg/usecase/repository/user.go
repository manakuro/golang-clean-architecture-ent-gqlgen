package repository

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
)

// User is interface of repository
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error)
}
