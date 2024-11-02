package status

import (
	"net/http"

	"onion_todo_app/usecase/status"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	createStatusUseCase *status.CreateStatusUseCase
}

func NewHandler() *StatusHandler {
	return &StatusHandler{
		createStatusUseCase: status.NewCreateStatusUseCase(),
	}
}

func (h *StatusHandler) CreateStatus(ctx echo.Context) error {
	var req CreateStatusRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dto := status.CreateStatusInputDTO{
		State: req.State,
	}

	if err := h.createStatusUseCase.Execute(ctx, dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, nil)
}
