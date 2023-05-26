package database

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	//ซ่อนenv
	loadEnv()
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	SSL_MODE := os.Getenv("SSL_MODE")

	fmt.Println("------------------------------")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, SSL_MODE)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("DSN:", dsn)

	fmt.Println("Database connection successful...")
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + "MPOS_backend" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
