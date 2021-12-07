package registry

import (
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/controller"
	"golang-clean-architecture-ent-gqlgen/pkg/adapter/repository"
	"golang-clean-architecture-ent-gqlgen/pkg/usecase/usecase"
)

// NewUserController conforms to interface
func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)

	return controller.NewUserController(u)
}
