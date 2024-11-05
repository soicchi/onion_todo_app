package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Priority struct {
	id        uuid.UUID
	level     string
	createdAt time.Time // zero value when it is initialized
	updatedAt time.Time // zero value when it is initialized
}

func NewPriority(level string) (*Priority, error) {
	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	return newPriority(primaryKey, level), nil
}

func ReconstructPriority(id uuid.UUID, level string) *Priority {
	return newPriority(id, level)
}

func newPriority(id uuid.UUID, level string) *Priority {
	return &Priority{
		id:    id,
		level: level,
	}
}

func (p *Priority) ID() uuid.UUID {
	return p.id
}

func (p *Priority) Level() string {
	return p.level
}

func (p *Priority) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Priority) UpdatedAt() time.Time {
	return p.updatedAt
}
