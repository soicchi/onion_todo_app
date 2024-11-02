package database

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}

type Todo struct {
	Base        `gorm:"embedded"`
	Title       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	StatusID    uuid.UUID `gorm:"type:uuid"`
	PriorityID  uuid.UUID `gorm:"type:uuid"`
}

type Status struct {
	Base  `gorm:"embedded"`
	State string `gorm:"type:varchar(64);unique;not null"`
	Todos []Todo `gorm:"foreignKey:StatusID"`
}

type Priority struct {
	Base  `gorm:"embedded"`
	Level string `gorm:"type:varchar(32);unique;not null"`
	Todos []Todo `gorm:"foreignKey:PriorityID"`
}
