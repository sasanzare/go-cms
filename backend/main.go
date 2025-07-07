package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sasanzare/go-cms/config"
	"github.com/sasanzare/go-cms/migrations"
	"github.com/sasanzare/go-cms/routes"
)

func main() {
	// Load configuration
	dbConfig := config.LoadDBConfig()

	// Connect to database
	db, err := config.ConnectDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run auto migrations
	if err := migrations.InitAutoMigrations(db); err != nil {
		log.Fatalf("Failed to run auto migrations: %v", err)
	}

	// Create Gin router
	r := gin.Default()

	// Setup main routes
	routes.SetupRouter(r, db)

	// Start server
	log.Println("Starting server on :8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}