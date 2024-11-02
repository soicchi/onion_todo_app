package repository

import (
	"fmt"

	domain "onion_todo_app/domain/todo"
	"onion_todo_app/infrastructure/postgres/database"

	"github.com/labstack/echo/v4"
)

type TodoRepository struct{}

func (tr TodoRepository) Create(ctx echo.Context, todo *domain.Todo) error {
	// Get the database connection
	db := database.GetDB()

	if err := db.Create(&database.Todo{
		Base: database.Base{
			ID: todo.ID(),
		},
		Title:       todo.Title(),
		Description: todo.Description(),
		PriorityID:  todo.PriorityID(),
		StatusID:    todo.StatusID(),
	}).Error; err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}

	return nil
}
