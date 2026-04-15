package resources

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectPostgres() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open Postgres connection: %w", err)
	}

	// Verify the connection is alive
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping Postgres: %w", err)
	}

	// Connection pool settings
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	DB = db
	log.Println("connected to Postgres successfully")
	return nil
}
