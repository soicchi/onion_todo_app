package repository

import (
	"fmt"

	"onion_todo_app/domain/todo"
	"onion_todo_app/infrastructure/postgres/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoRepository struct{}

func (tr TodoRepository) Create(ctx echo.Context, todo *todo.Todo) error {
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

func (tr TodoRepository) FetchAll(ctx echo.Context) ([]*todo.TodoDetail, error) {
	db := database.GetDB()

	var todos []*database.Todo
	if err := db.Preload("Priority", func(db *gorm.DB) *gorm.DB {
		return db.Select("priorities.id", "priorities.level")
	}).Preload("Status", func(db *gorm.DB) *gorm.DB {
		return db.Select("statuses.id", "statuses.state")
	}).Find(&todos).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}

	var result []*todo.TodoDetail
	for _, t := range todos {
		if t.Priority == nil || t.Status == nil {
			return nil, fmt.Errorf("related Priority or Status is missing for Todo ID: %v", t.ID)
		}
		result = append(result, todo.ReconstructTodoDetail(
			t.ID,
			t.Priority.ID,
			t.Status.ID,
			t.Title,
			t.Description,
			t.Priority.Level,
			t.Status.State,
		))
	}

	return result, nil
}
