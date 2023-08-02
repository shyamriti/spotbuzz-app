package database

import (
	"fmt"
	"log"
	"spotbuzz/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	dbUsername = "shyam"
	dbPassword = "password@123"
	dbName     = "new_db"
)

// Connection initializes the database connection and performs auto-migration for the Player model.
func Connection() {
	var err error

	// Construct the DSN (Data Source Name) for the PostgreSQL connection
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kolkata", dbUsername, dbPassword, dbName)
	// dsn := fmt.Sprintf("host=host.docker.internal user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kolkata", dbUsername, dbPassword, dbName)
	// Open a connection to the PostgreSQL database using GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Perform auto-migration for the Player model
	err = DB.AutoMigrate(&model.Player{})
	if err != nil {
		panic(err)
	}
}
