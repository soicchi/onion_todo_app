package router

import (
	"onion_todo_app/presentation/handlers/priority"
	"onion_todo_app/presentation/handlers/status"
	"onion_todo_app/presentation/handlers/todo"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo) {
	const basePath = "/api/v1"
	baseGroup := e.Group(basePath)

	TodoRouter(e, baseGroup)
}

func TodoRouter(e *echo.Echo, group *echo.Group) {
	todoGroup := e.Group("/todos")

	handler := todo.NewHandler()
	todoGroup.POST("/", handler.CreateTodo)
}

func StatusRouter(e *echo.Echo, group *echo.Group) {
	statusGroup := e.Group("/statuses")

	handler := status.NewHandler()
	statusGroup.POST("/", handler.CreateStatus)
}

func PriorityRouter(e *echo.Echo, group *echo.Group) {
	priorityGroup := e.Group("/priorities")

	handler := priority.NewHandler()
	priorityGroup.POST("/", handler.CreatePriority)
}
