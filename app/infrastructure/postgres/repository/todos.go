package repository

import (
	"fmt"

	"onion_todo_app/domain/entity"
	"onion_todo_app/infrastructure/postgres/database"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoRepository struct {
	dbConn dbConnector
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{dbConn: database.DB{}}
}

func (tr *TodoRepository) Create(ctx echo.Context, todo *entity.Todo, priorityID, statusID uuid.UUID) error {
	// Get the database connection
	db := tr.dbConn.GetDB(ctx)

	if err := db.Create(&database.Todo{
		Base: database.Base{
			ID: todo.ID(),
		},
		Title:       todo.Title(),
		Description: todo.Description(),
		PriorityID:  priorityID,
		StatusID:    statusID,
	}).Error; err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}

	return nil
}

func (tr *TodoRepository) FetchAll(ctx echo.Context) ([]*entity.TodoDetail, error) {
	db := tr.dbConn.GetDB(ctx)

	var todos []*database.Todo
	if err := db.Preload("Priority", func(db *gorm.DB) *gorm.DB {
		return db.Select("priorities.id", "priorities.level")
	}).Preload("Status", func(db *gorm.DB) *gorm.DB {
		return db.Select("statuses.id", "statuses.state")
	}).Find(&todos).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}

	var result []*entity.TodoDetail
	for _, t := range todos {
		if t.Priority == nil || t.Status == nil {
			return nil, fmt.Errorf("related Priority or Status is missing for Todo ID: %v", t.ID)
		}
		result = append(result, entity.ReconstructTodoDetail(
			entity.ReconstructTodo(t.ID, t.Title, t.Description),
			entity.ReconstructPriority(t.Priority.ID, t.Priority.Level),
			entity.ReconstructStatus(t.Status.ID, t.Status.State),
		))
	}

	return result, nil
}
