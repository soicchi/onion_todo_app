package main

import (
	"onion_todo_app/infrastructure/postgres/database"
	"onion_todo_app/presentation/router"
	"onion_todo_app/presentation/validator"

	"github.com/labstack/echo/v4"
)

func main() {
	dbConfig, err := database.NewDBConfig()
	if err != nil {
		panic(err)
	}

	// Connect to the database
	if err := dbConfig.ConnectDB(); err != nil {
		panic(err)
	}

	// Migrate the database
	if err := database.Migrate(); err != nil {
		panic(err)
	}

	// Initialize router
	e := echo.New()
	e.Validator = validator.NewValidator()
	router.NewRouter(e)

	for _, router := range e.Routes() {
		println(router.Method, router.Path)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
