package registry

import (
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/controller"
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/repository"
	"golang-clean-architecture-ent-gqlgen/pkg/usecase/usecase"
)

func (r *registry) NewTodoController() controller.Todo {
	repo := repository.NewTodoRepository(r.client)
	u := usecase.NewTodoUsecase(repo)

	return controller.NewTodoController(u)
}
