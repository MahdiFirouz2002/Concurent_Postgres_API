package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var database *sqlx.DB

const (
	DbUser     = "postgres"
	DbPassword = "your_password"
	DbName     = "your_database"
	DbHost     = "localhost"
	DbSSLMode  = "disable"
)

func ConnectToDB() error {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s", DbUser, DbName, DbSSLMode, DbPassword, DbHost)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	database = db
	return nil
}

func Close_database() {
	database.Close()
}

func checkConnection() error {
	if err := database.Ping(); err != nil {
		return err
	}

	return nil
}
