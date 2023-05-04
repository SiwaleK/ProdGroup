package main

import (
	"database/sql"
	"log"

	"github.com/SiwaleK/ProdGroup/pkg/sale"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *sql.DB
}

var DB *gorm.DB

func main() {
	var err error

	DB, err = gorm.Open(postgres.Open("postgresql://root:secret@localhost:5432/mos_sku0?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	prodgroupRepo := sale.NewProdgroupRepository(sqlDB)
	prodgroupHandler := sale.NewProdgroupHandler(prodgroupRepo)

	//promotionRepo := repository.NewDBPromotionRepository(db)
	//promotionHandler := sale.NewPromotionHandler(promotionRepo)

	r := gin.Default()

	r.GET("/prodgroups", prodgroupHandler.GetProdgroup)
	//r.GET("/promotions", promotionHandler.GetPromotions)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
