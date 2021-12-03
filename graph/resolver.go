package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"golang-clean-architecture-ent-gqlgen/ent"
	"golang-clean-architecture-ent-gqlgen/graph/generated"

	"github.com/99designs/gqlgen/graphql"
)

// Resolver is a context struct
type Resolver struct{ client *ent.Client }

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
