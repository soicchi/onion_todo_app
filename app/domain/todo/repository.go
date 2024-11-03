package todo

import (
	"github.com/labstack/echo/v4"
)

type TodoRepository interface {
	Create(ctx echo.Context, todo *Todo) error
	FetchAll(ctx echo.Context) ([]*TodoDetail, error)
}
