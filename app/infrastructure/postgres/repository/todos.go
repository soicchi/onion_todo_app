package todos

import (
	"fmt"

	"onion_todo_app/domain/todo"
	"onion_todo_app/infrastructure/postgres/database"

	"github.com/labstack/echo/v4"
)

type TodoRepository struct{}

func (tr *TodoRepository) Create(ctx echo.Context, todo *todo.Todo) error {
	// Get the database connection
	db := getDB()
	todoModel := &database.Todo{
		ID:          todo.ID(),
		Title:       todo.Title(),
		Description: todo.Description(),
		StatusID:    todo.StatusID(),
		PriorityID:  todo.PriorityID(),
	}

	if err := db.Create(todoModel).Error; err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}

	return nil
}
