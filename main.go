package main

import (
	"log"
	"taskflow-samrat/resources"
	"taskflow-samrat/routers"

	"github.com/joho/godotenv"

	_ "taskflow-samrat/docs"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables: %v", err)
	}

	resources.InitConfig()

	// Connect to database
	if err := resources.ConnectPostgres(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer resources.DB.Close()

	r := routers.InitRouters()

	r.Run(":8080")
}
