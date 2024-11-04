package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type dbConnector interface {
	GetDB(ctx echo.Context) *gorm.DB
}
