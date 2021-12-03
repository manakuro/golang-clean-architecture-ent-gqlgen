// Code generated by entc, DO NOT EDIT.

package ent

import (
	"golang-clean-architecture-ent-gqlgen/ent/todo"
	"time"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Name      *string
	Status    *todo.Status
	Priority  *int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	UserID    *int
}

// Mutate applies the CreateTodoInput on the TodoCreate builder.
func (i *CreateTodoInput) Mutate(m *TodoCreate) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the create builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c)
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	ID        int
	Name      *string
	Status    *todo.Status
	Priority  *int
	UserID    *int
	ClearUser bool
}

// Mutate applies the UpdateTodoInput on the TodoMutation.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the update builder.
func (u *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateTodoInput on the update-one builder.
func (u *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name      string
	Age       int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	TodoIDs   []int
}

// Mutate applies the CreateUserInput on the UserCreate builder.
func (i *CreateUserInput) Mutate(m *UserCreate) {
	m.SetName(i.Name)
	m.SetAge(i.Age)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if ids := i.TodoIDs; len(ids) > 0 {
		m.AddTodoIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the create builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c)
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	ID            int
	Name          *string
	Age           *int
	AddTodoIDs    []int
	RemoveTodoIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Age; v != nil {
		m.SetAge(*v)
	}
	if ids := i.AddTodoIDs; len(ids) > 0 {
		m.AddTodoIDs(ids...)
	}
	if ids := i.RemoveTodoIDs; len(ids) > 0 {
		m.RemoveTodoIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the update builder.
func (u *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateUserInput on the update-one builder.
func (u *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}