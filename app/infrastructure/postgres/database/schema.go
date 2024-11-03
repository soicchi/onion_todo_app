package database

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;autoUpdateTime"`
}

type Todo struct {
	Base
	Title       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	StatusID    uuid.UUID `gorm:"type:uuid"`
	PriorityID  uuid.UUID `gorm:"type:uuid"`
	Status      *Status
	Priority    *Priority
}

type Status struct {
	Base
	State string `gorm:"type:varchar(64);unique;not null"`
	Todos []Todo `gorm:"foreignKey:StatusID"`
}

type Priority struct {
	Base
	Level string `gorm:"type:varchar(32);unique;not null"`
	Todos []Todo `gorm:"foreignKey:PriorityID"`
}
