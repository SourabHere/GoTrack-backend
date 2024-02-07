package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
)

func init() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
	user = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	dbname = os.Getenv("DBNAME")
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Database interface {
	InitDB() (*sql.DB, error)
}

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB() (*sql.DB, error) {

	port, err := strconv.Atoi(port)

	if err != nil {
		return nil, fmt.Errorf("error converting port to integer: %w", err)
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL database: %w", err)
	}

	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(5)

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("error pinging PostgreSQL database: %w", err)
	}

	return db, nil
}
