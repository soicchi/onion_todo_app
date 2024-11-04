package entity

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoRepository interface {
	Create(ctx echo.Context, todo *Todo, priorityID, statusID uuid.UUID) error
	FetchAll(ctx echo.Context) ([]*TodoDetail, error)
}

type PriorityRepository interface {
	Create(ctx echo.Context, priority *Priority) error
}

type StatusRepository interface {
	Create(ctx echo.Context, status *Status) error
}
