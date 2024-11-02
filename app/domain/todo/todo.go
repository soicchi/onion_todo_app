package todo

import (
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	id          string
	title       string
	description string
	priorityID  string
	statusID    string
}

func NewTodo(title, description, priorityID, statusID string) (*Todo, error) {
	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	return &Todo{
		id:          primaryKey.String(),
		title:       title,
		description: description,
		priorityID:  priorityID,
		statusID:    statusID,
	}, nil
}

func (t *Todo) ID() string {
	return t.id
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Description() string {
	return t.description
}

func (t *Todo) PriorityID() string {
	return t.priorityID
}

func (t *Todo) StatusID() string {
	return t.statusID
}
