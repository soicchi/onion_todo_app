package priority

import (
	"fmt"

	"github.com/google/uuid"
)

type Priority struct {
	id string
	level string
}

func NewPriority(level string) (*Priority, error) {
	primaryKey, err := uuid.New()
	if err != nil {
		return nil, fmt.Errorf("error generating primary key: %w", err)
	}

	return &Priority{
		id: primaryKey.String(),
		level: level,
	}, nil
}

func (p *Priority) ID() uuid.UUID {
	return p.id
}

func (p *Priority) Level() string {
	return p.level
}
