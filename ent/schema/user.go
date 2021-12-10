package schema

import (
	"golang-clean-architecture-ent-gqlgen/ent/mixin"
	"golang-clean-architecture-ent-gqlgen/pkg/const/globalid"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the Test entity.
type User struct {
	ent.Schema
}

// UserMixin defines Fields
type UserMixin struct {
	entMixin.Schema
}

// Fields of the User.
func (UserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Int("age").Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().User.Prefix),
		UserMixin{},
		mixin.NewDatetime(),
	}
}
