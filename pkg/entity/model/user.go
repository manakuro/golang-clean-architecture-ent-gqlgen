package model

import "golang-clean-architecture-ent-gqlgen/ent"

// User is the model entity for the User schema.
type User = ent.User

// UserConnection is the connection containing edges to User.
type UserConnection = ent.UserConnection

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput = ent.CreateUserInput

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput = ent.UpdateUserInput
