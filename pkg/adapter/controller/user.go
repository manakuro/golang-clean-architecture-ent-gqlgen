package controller

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
	"golang-clean-architecture-ent-gqlgen/pkg/usecase/usecase"
)

type user struct {
	userUsecase usecase.User
}

// User of interface
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error)
}

// NewUserController returns user controller
func NewUserController(uu usecase.User) User {
	return &user{userUsecase: uu}
}

func (u *user) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return u.userUsecase.Get(ctx, id)
}

func (u *user) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error) {
	return u.userUsecase.List(ctx, after, first, before, last, where)
}

func (u *user) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return u.userUsecase.Create(ctx, input)
}

func (u *user) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return u.userUsecase.Update(ctx, input)
}
