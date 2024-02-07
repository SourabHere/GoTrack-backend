package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database interface {
	InitDB() (*sql.DB, error)
}

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB() (*sql.DB, error) {
	connectionString := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL database: %w", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("error pinging PostgreSQL database: %w", err)
	}

	return db, nil
}
