package todo

import (
	"net/http"

	useCase "onion_todo_app/usecase/todo"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	createTodoUseCase *useCase.CreateTodoUseCase
}

func NewHandler() *Handler {
	return &Handler{
		createTodoUseCase: useCase.NewCreateTodoUseCase(),
	}
}

func (h *Handler) CreateTodo(ctx echo.Context) error {
	var req CreateRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dto := useCase.CreateTodoInputDTO{
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
