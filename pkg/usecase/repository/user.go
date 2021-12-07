package repository

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
)

// User is interface of repository
type User interface {
	Get(ctx context.Context, id *int) (*model.User, error)
}
