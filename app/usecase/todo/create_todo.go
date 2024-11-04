package todo

import (
	"fmt"

	"onion_todo_app/domain/entity"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateTodoUseCase struct {
	todoRepository entity.TodoRepository
}

func NewCreateTodoUseCase() *CreateTodoUseCase {
	return &CreateTodoUseCase{
		todoRepository: repository.NewTodoRepository(),
	}
}

type CreateTodoInputDTO struct {
	Title       string
	Description string
	PriorityID  string
	StatusID    string
}

func (uc *CreateTodoUseCase) Execute(ctx echo.Context, dto CreateTodoInputDTO) error {
	todo, err := entity.NewTodo(dto.Title, dto.Description)
	if err != nil {
		return fmt.Errorf("error creating new todo: %w", err)
	}

	priorityUUID, err := uuid.Parse(dto.PriorityID)
	if err != nil {
		return fmt.Errorf("error parsing priority ID: %w", err)
	}

	statusUUID, err := uuid.Parse(dto.StatusID)
	if err != nil {
		return fmt.Errorf("error parsing status ID: %w", err)
	}

	if err := uc.todoRepository.Create(ctx, todo, priorityUUID, statusUUID); err != nil {
		return fmt.Errorf("error creating todo: %w", err)
	}

	return nil
}
