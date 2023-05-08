package config

import (
	"database/sql"
	"log"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn   *sql.DB
	Logger zerolog.Logger
}

type Config struct {
	Logger zerolog.Logger
}

var DB *gorm.DB

func Connect(cfg Config) (Database, error) {
	db := Database{}
	var err error

	DB, err = gorm.Open(postgres.Open("postgresql://root:secret@localhost:5432/mos_sku0?sslmode=disable"), &gorm.Config{})
	//DB, err = gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.Logger = cfg.Logger
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}
