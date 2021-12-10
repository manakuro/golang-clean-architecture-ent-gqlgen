package controller

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
	"golang-clean-architecture-ent-gqlgen/pkg/usecase/usecase"
)

// Todo is an interface of controller
type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
	Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error)
}

type testTodo struct {
	testTodoUsecase usecase.Todo
}

// NewTodoController generates test user controller
func NewTodoController(tu usecase.Todo) Todo {
	return &testTodo{
		testTodoUsecase: tu,
	}
}

func (t *testTodo) Get(ctx context.Context, id *model.ID) (*model.Todo, error) {
	return t.testTodoUsecase.Get(ctx, id)
}

func (t *testTodo) List(ctx context.Context) ([]*model.Todo, error) {
	return t.testTodoUsecase.List(ctx)
}

func (t *testTodo) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return t.testTodoUsecase.Create(ctx, input)
}

func (t *testTodo) Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	return t.testTodoUsecase.Update(ctx, input)
}
