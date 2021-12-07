package model

import "golang-clean-architecture-ent-gqlgen/ent"

// Todo is the model entity for the Todo schema.
type Todo = ent.Todo

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput = ent.CreateTodoInput

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput = ent.UpdateTodoInput
