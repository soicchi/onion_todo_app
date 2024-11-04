package repository

import (
	"fmt"

	"onion_todo_app/domain/entity"
	"onion_todo_app/infrastructure/postgres/database"

	"github.com/labstack/echo/v4"
)

type PriorityRepository struct {
	dbConn dbConnector
}

func NewPriorityRepository() *PriorityRepository {
	return &PriorityRepository{dbConn: database.DB{}}
}

func (pr *PriorityRepository) Create(ctx echo.Context, priority *entity.Priority) error {
	db := pr.dbConn.GetDB(ctx)

	if err := db.Create(&database.Priority{
		Base: database.Base{
			ID: priority.ID(),
		},
		Level: priority.Level(),
	}).Error; err != nil {
		return fmt.Errorf("error creating priority: %w", err)
	}

	return nil
}
