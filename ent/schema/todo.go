package schema

import (
	"golang-clean-architecture-ent-gqlgen/ent/mixin"
	"golang-clean-architecture-ent-gqlgen/ent/schema/ulid"
	"golang-clean-architecture-ent-gqlgen/pkg/const/globalid"

	"entgo.io/ent/schema/edge"
	entMixin "entgo.io/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// TodoMixin defines Fields
type TodoMixin struct {
	entMixin.Schema
}

// Fields of the Todo.
func (TodoMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			GoType(ulid.ID("")).
			Optional(),
		field.String("name").Default(""),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
		field.Int("priority").Default(0),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("todos").
			Unique().
			Field("user_id"),
	}
}

// Mixin of the Todo.
func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Todo.Prefix),
		TodoMixin{},
		mixin.NewDatetime(),
	}
}
