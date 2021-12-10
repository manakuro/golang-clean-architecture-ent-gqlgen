package repository

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/ent/user"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
	usecaseRepository "golang-clean-architecture-ent-gqlgen/pkg/usecase/repository"
)

type userRepository struct {
	client *ent.Client
}

// NewUserRepository is specific implementation of the interface
func NewUserRepository(client *ent.Client) usecaseRepository.User {
	return &userRepository{client: client}
}

func (r *userRepository) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	u, err := r.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error) {
	us, err := r.client.
		User.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (r *userRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	u, err := r.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}

func (r *userRepository) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	u, err := r.client.User.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return u, nil
}
