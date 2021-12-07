package usecase

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
	"golang-clean-architecture-ent-gqlgen/pkg/usecase/repository"
)

type user struct {
	userRepository repository.User
}

// User of usecase.
type User interface {
	Get(ctx context.Context, id *int) (*model.User, error)
}

// NewUserUsecase returns user usecase.
func NewUserUsecase(r repository.User) User {
	return &user{userRepository: r}
}

func (u *user) Get(ctx context.Context, id *int) (*model.User, error) {
	return u.userRepository.Get(ctx, id)
}
