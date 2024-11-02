package priority

import (
	"fmt"

	"github.com/google/uuid"
)

type Priority struct {
	id uuid.UUID
	level string
}

func NewPriority(level string) (*Priority, error) {
	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	return &Priority{
		id: primaryKey,
		level: level,
	}, nil
}

func (p *Priority) ID() uuid.UUID {
	return p.id
}

func (p *Priority) Level() string {
	return p.level
}
