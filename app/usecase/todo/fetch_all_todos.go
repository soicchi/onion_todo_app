package todo

import (
	"onion_todo_app/domain/todo"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/labstack/echo/v4"
)

type FetchAllTodosUseCase struct {
	todoRepository todo.TodoRepository
}

func NewFetchAllTodosUseCase() *FetchAllTodosUseCase {
	return &FetchAllTodosUseCase{
		todoRepository: repository.TodoRepository{},
	}
}

type fetchAllPriorityOutputDTO struct {
	ID    string
	Level string
}

type fetchAllStatusOutputDTO struct {
	ID    string
	State string
}

type fetchAllTodoOutputDTO struct {
	ID          string
	Title       string
	Description string
	Priority    fetchAllPriorityOutputDTO
	Status      fetchAllStatusOutputDTO
}

type FetchAllTodosOutputDTO struct {
	Todos []fetchAllTodoOutputDTO
}

func (uc *FetchAllTodosUseCase) Execute(ctx echo.Context) (*FetchAllTodosOutputDTO, error) {
	todos, err := uc.todoRepository.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	var dto *FetchAllTodosOutputDTO
	for _, t := range todos {
		dto.Todos = append(dto.Todos, fetchAllTodoOutputDTO{
			ID:          t.ID().String(),
			Title:       t.Title(),
			Description: t.Description(),
			Priority: fetchAllPriorityOutputDTO{
				ID:    t.PriorityID().String(),
				Level: t.Priority(),
			},
			Status: fetchAllStatusOutputDTO{
				ID:    t.StatusID().String(),
				State: t.Status(),
			},
		})
	}

	return dto, nil
}
