package todo

import (
	"net/http"

	"onion_todo_app/usecase/todo"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	createTodoUseCase    *todo.CreateTodoUseCase
	fetchAllTodosUseCase *todo.FetchAllTodosUseCase
}

func NewHandler() *Handler {
	return &Handler{
		createTodoUseCase:    todo.NewCreateTodoUseCase(),
		fetchAllTodosUseCase: todo.NewFetchAllTodosUseCase(),
	}
}

func (h *Handler) CreateTodo(ctx echo.Context) error {
	var req CreateTodoRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dto := todo.CreateTodoInputDTO{
		Title:       req.Title,
		Description: req.Description,
		PriorityID:  req.PriorityID,
		StatusID:    req.StatusID,
	}

	if err := h.createTodoUseCase.Execute(ctx, dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (h *Handler) FetchAllTodos(ctx echo.Context) error {
	dto, err := h.fetchAllTodosUseCase.Execute(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto)
}
