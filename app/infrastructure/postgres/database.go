package postgres

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type DBConfig struct {
	host     string
	port     int
	user     string
	password string
	name   string
	sslMode string
}

func NewDBConfig() (*DBConfig, error) {
	intPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Errorf("failed to convert port to int: %w", err)
	}

	return &DBConfig{
		host: os.Getenv("DB_HOST"),
		port: intPort,
		user: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		name: os.Getenv("DB_NAME"),
		sslMode: os.Getenv("DB_SSL_MODE"),
	}, nil
}

func (dc *DBConfig) dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dc.host, dc.port, dc.user, dc.password, dc.name, dc.sslMode)
}

func (dc *DBConfig) ConnectDB() error {
	var err error
	db, err = gorm.Open(postgres.Open(dc.dsn()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return nil
}
