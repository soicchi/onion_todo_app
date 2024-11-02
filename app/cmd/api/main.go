package main

import (
	"onion_todo_app/infrastructure/postgres/database"

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

	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
