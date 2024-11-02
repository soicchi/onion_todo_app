package status

import (
	"fmt"

	"github.com/google/uuid"
)

type Status struct {
	id uuid.UUID
	state string
}

func NewStatus(state string) (*Status, error) {
	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	return &Status{
		id: primaryKey,
		state: state,
	}, nil
}

func (s *Status) ID() uuid.UUID {
	return s.id
}

func (s *Status) State() string {
	return s.state
}
