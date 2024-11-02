package router

import (
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
