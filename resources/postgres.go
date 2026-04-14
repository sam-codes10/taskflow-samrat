package resources

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectPostgres() error {
	
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		AppConfig.Config.Postgres.Host, AppConfig.Config.Postgres.Port, AppConfig.Config.Postgres.User, AppConfig.Config.Postgres.Password, AppConfig.Config.Postgres.DbName)

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