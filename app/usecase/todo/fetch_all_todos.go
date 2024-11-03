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
	ID    string `json:"id"`
	Level string `json:"level"`
}

type fetchAllStatusOutputDTO struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

type fetchAllTodoOutputDTO struct {
	ID          string                    `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Priority    fetchAllPriorityOutputDTO `json:"priority"`
	Status      fetchAllStatusOutputDTO   `json:"status"`
}

type FetchAllTodosOutputDTO struct {
	Todos []fetchAllTodoOutputDTO `json:"todos"`
}

func (uc *FetchAllTodosUseCase) Execute(ctx echo.Context) (*FetchAllTodosOutputDTO, error) {
	todos, err := uc.todoRepository.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return uc.constructDTO(todos), nil
}

func (uc *FetchAllTodosUseCase) constructDTO(todos []*todo.TodoDetail) *FetchAllTodosOutputDTO {
	var dto FetchAllTodosOutputDTO
	dtoTodos := make([]fetchAllTodoOutputDTO, 0, len(todos))

	for _, t := range todos {
		dtoTodos = append(dtoTodos, fetchAllTodoOutputDTO{
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
	dto.Todos = dtoTodos

	return &dto
}
