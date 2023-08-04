package database

import (
	"log"
	"spotbuzz/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection initializes the database connection and performs auto-migration for the Player model.
func Connection() {
	var err error
	// Construct the DSN (Data Source Name) for the PostgreSQL connection
	dsn := "host=database user=shyam password=password@123 dbname=new_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Perform auto-migration for the Player model
	err = DB.AutoMigrate(&model.Player{})
	if err != nil {
		panic(err)
	}
}
