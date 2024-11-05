package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	id          uuid.UUID
	title       string
	description string
	createdAt time.Time
	updatedAt time.Time
}

func NewTodo(title, description string) (*Todo, error) {
	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	return newTodo(primaryKey, title, description), nil
}

func ReconstructTodo(id uuid.UUID, title, description string) *Todo {
	return newTodo(id, title, description)
}

func newTodo(id uuid.UUID, title, description string) *Todo {
	return &Todo{
		id:          id,
		title:       title,
		description: description,
	}
}

func (t *Todo) ID() uuid.UUID {
	return t.id
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Description() string {
	return t.description
}

func (t *Todo) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Todo) UpdatedAt() time.Time {
	return t.updatedAt
}
