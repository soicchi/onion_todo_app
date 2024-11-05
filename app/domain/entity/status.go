package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Status struct {
	id    uuid.UUID
	state string
	createdAt time.Time // zero value when it is initialized
	updatedAt time.Time // zero value when it is initialized
}

func NewStatus(state string) (*Status, error) {
	if state == "" {
		return nil, fmt.Errorf("state cannot be empty")
	}

	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	return newStatus(primaryKey, state), nil
}

func ReconstructStatus(id uuid.UUID, state string) *Status {
	return newStatus(id, state)
}

func newStatus(id uuid.UUID, state string) *Status {
	return &Status{
		id:    id,
		state: state,
	}
}

func (s *Status) ID() uuid.UUID {
	return s.id
}

func (s *Status) State() string {
	return s.state
}

func (s *Status) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Status) UpdatedAt() time.Time {
	return s.updatedAt
}
