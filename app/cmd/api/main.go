package main

import (
	"onion_todo_app/infrastructure/postgres"

	"github.com/labstack/echo/v4"
)

func main() {
	dbConfig, err := postgres.NewDBConfig()
	if err != nil {
		panic(err)
	}

	// Connect to the database
	if err := dbConfig.ConnectDB(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
