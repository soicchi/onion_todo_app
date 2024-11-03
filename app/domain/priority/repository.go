package priority

import (
	"github.com/labstack/echo/v4"
)

type PriorityRepository interface {
	Create(ctx echo.Context, priority *Priority) error
}
