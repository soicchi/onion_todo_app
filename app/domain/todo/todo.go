package todo

import (
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	id          uuid.UUID
	title       string
	description string
	priorityID  uuid.UUID
	statusID    uuid.UUID
}

func NewTodo(title, description, priorityID, statusID string) (*Todo, error) {
	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	priorityUUID, err := uuid.Parse(priorityID)
	if err != nil {
		return nil, fmt.Errorf("error parsing priority ID: %w", err)
	}

	statusUUID, err := uuid.Parse(statusID)
	if err != nil {
		return nil, fmt.Errorf("error parsing status ID: %w", err)
	}

	return &Todo{
		id:          primaryKey,
		title:       title,
		description: description,
		priorityID:  priorityUUID,
		statusID:    statusUUID,
	}, nil
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

func (t *Todo) PriorityID() uuid.UUID {
	return t.priorityID
}

func (t *Todo) StatusID() uuid.UUID {
	return t.statusID
}
