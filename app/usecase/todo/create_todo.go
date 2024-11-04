package todo

import (
	"fmt"

	"onion_todo_app/domain/todo"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/labstack/echo/v4"
)

type CreateTodoUseCase struct {
	todoRepository todo.TodoRepository
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
	todo, err := todo.NewTodo(dto.Title, dto.Description, dto.PriorityID, dto.StatusID)
	if err != nil {
		return fmt.Errorf("error creating new todo: %w", err)
	}

	if err := uc.todoRepository.Create(ctx, todo); err != nil {
		return fmt.Errorf("error creating todo: %w", err)
	}

	return nil
}
