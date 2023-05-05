package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SiwaleK/ProdGroup/router"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *sql.DB
}

var DB *gorm.DB

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	// Get DB source from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to database
	DB, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{
		Logger: nil,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	routers := router.RegisterRoute(sqlDB)
	if err := routers.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
