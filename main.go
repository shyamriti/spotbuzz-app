package main

import (
	"spotbuzz/pkg/database"
	"spotbuzz/pkg/router"
)

func main() {
	// Initialize the database connection and perform auto-migration for the models
	database.Connection()

	// Initialize the Gin router and define the routes
	r := router.Router()

	// Start the server and listen on port 8080
	r.Run(":8080")
}
