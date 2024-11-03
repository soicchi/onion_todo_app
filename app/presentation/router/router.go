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

	TodoRouter(baseGroup)
	StatusRouter(baseGroup)
	PriorityRouter(baseGroup)
}

func TodoRouter(group *echo.Group) {
	todoGroup := group.Group("/todos")

	handler := todo.NewHandler()
	todoGroup.POST("/", handler.CreateTodo)
	todoGroup.GET("/", handler.FetchAllTodos)
}

func StatusRouter(group *echo.Group) {
	statusGroup := group.Group("/statuses")

	handler := status.NewHandler()
	statusGroup.POST("/", handler.CreateStatus)
}

func PriorityRouter(group *echo.Group) {
	priorityGroup := group.Group("/priorities")

	handler := priority.NewHandler()
	priorityGroup.POST("/", handler.CreatePriority)
}
