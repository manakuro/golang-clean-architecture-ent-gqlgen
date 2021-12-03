package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/ent/user"
	"golang-clean-architecture-ent-gqlgen/graph/generated"
	"golang-clean-architecture-ent-gqlgen/pkg/util/datetime"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	u, err := r.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*ent.User, error) {
	u, err := r.client.User.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *queryResolver) User(ctx context.Context, id *int) (*ent.User, error) {
	u, err := r.client.User.Query().Where(user.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
