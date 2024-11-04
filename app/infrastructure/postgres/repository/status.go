package repository

import (
	"fmt"

	domain "onion_todo_app/domain/status"
	"onion_todo_app/infrastructure/postgres/database"

	"github.com/labstack/echo/v4"
)

type StatusRepository struct {
	dbConn dbConnector
}

func NewStatusRepository() *StatusRepository {
	return &StatusRepository{dbConn: database.DB{}}
}

func (sr StatusRepository) Create(ctx echo.Context, status *domain.Status) error {
	// Get the database connection
	db := sr.dbConn.GetDB(ctx)

	if err := db.Create(&database.Status{
		Base: database.Base{
			ID: status.ID(),
		},
		State: status.State(),
	}).Error; err != nil {
		return fmt.Errorf("failed to create status: %w", err)
	}

	return nil
}
