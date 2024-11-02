package status

import (
	"github.com/labstack/echo/v4"
)

type StatusRepository interface {
	Create(ctx echo.Context, status *Status) error
}
