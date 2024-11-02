package database

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	StatusID uuid.UUID `gorm:"type:uuid"`
	PriorityID uuid.UUID `gorm:"type:uuid"`
}

type Status struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	State string `gorm:"type:varchar(64);unique"`
	Todos []Todo `gorm:"foreignKey:StatusID"`
}

type Priority struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Level string `gorm:"type:varchar(32);unique"`
	Todos []Todo `gorm:"foreignKey:PriorityID"`
}
