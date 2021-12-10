package usecase

import (
	"context"
	"golang-clean-architecture-ent-gqlgen/pkg/entity/model"
	"golang-clean-architecture-ent-gqlgen/pkg/usecase/repository"
)

type testTodo struct {
	testTodoRepository repository.Todo
}

// Todo is an interface of test user
type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
	Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error)
}

// NewTodoUsecase generates test user repository
func NewTodoUsecase(r repository.Todo) Todo {
	return &testTodo{testTodoRepository: r}
}

func (t *testTodo) Get(ctx context.Context, id *model.ID) (*model.Todo, error) {
	return t.testTodoRepository.Get(ctx, id)
}

func (t *testTodo) List(ctx context.Context) ([]*model.Todo, error) {
	return t.testTodoRepository.List(ctx)
}

func (t *testTodo) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return t.testTodoRepository.Create(ctx, input)
}

func (t *testTodo) Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	return t.testTodoRepository.Update(ctx, input)
}
