package priority

import (
	"net/http"

	"onion_todo_app/usecase/priority"

	"github.com/labstack/echo/v4"
)

type PriorityHandler struct {
	createPriorityUseCase *priority.CreatePriorityUseCase
}

func NewHandler() *PriorityHandler {
	return &PriorityHandler{
		createPriorityUseCase: priority.NewCreatePriorityUseCase(),
	}
}

func (h *PriorityHandler) CreatePriority(ctx echo.Context) error {
	var req CreatePriorityRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dto := priority.CreatePriorityInputDTO{
		Level: req.Level,
	}

	if err := h.createPriorityUseCase.Execute(ctx, dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, nil)
}
