package todo

import (
	"onion_todo_app/domain/entity"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/labstack/echo/v4"
)

type FetchAllTodosUseCase struct {
	todoRepository entity.TodoRepository
}

func NewFetchAllTodosUseCase() *FetchAllTodosUseCase {
	return &FetchAllTodosUseCase{
		todoRepository: repository.NewTodoRepository(),
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

func (uc *FetchAllTodosUseCase) constructDTO(todos []*entity.TodoDetail) *FetchAllTodosOutputDTO {
	var dto FetchAllTodosOutputDTO
	dtoTodos := make([]fetchAllTodoOutputDTO, 0, len(todos))

	for _, t := range todos {
		dtoTodos = append(dtoTodos, fetchAllTodoOutputDTO{
			ID:          t.Todo().ID().String(),
			Title:       t.Todo().Title(),
			Description: t.Todo().Description(),
			Priority: fetchAllPriorityOutputDTO{
				ID:    t.Priority().ID().String(),
				Level: t.Priority().Level(),
			},
			Status: fetchAllStatusOutputDTO{
				ID:    t.Status().ID().String(),
				State: t.Status().State(),
			},
		})
	}
	dto.Todos = dtoTodos

	return &dto
}
