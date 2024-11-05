package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	name      string
	createdAt time.Time // zero value when it is initialized
	updatedAt time.Time // zero value when it is initialized
}

func NewCreator(name string) (*User, error) {
	primaryKey, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return newCreator(primaryKey, name), nil
}

func ReconstructCreator(id uuid.UUID, name string) *User {
	return newCreator(id, name)
}

func newCreator(id uuid.UUID, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (c *User) ID() uuid.UUID {
	return c.id
}

func (c *User) Name() string {
	return c.name
}

func (c *User) CreatedAt() time.Time {
	return c.createdAt
}

func (c *User) UpdatedAt() time.Time {
	return c.updatedAt
}
