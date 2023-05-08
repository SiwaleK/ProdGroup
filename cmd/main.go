package main

import (
	"database/sql"
	"log"
	"os"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/SiwaleK/ProdGroup/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	// dbSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	// Connect to database
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	queries := db.New(dbConn)

	routers := router.RegisterRoute(queries)
	if err := routers.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
