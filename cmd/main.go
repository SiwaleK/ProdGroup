package main

import (
	"fmt"
	"os"

	"example.com/go-crud-api/db/database"
	"example.com/go-crud-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")

	database.DatabaseConnection()

	r := gin.Default()

	ROUTER_CONFIX := os.Getenv("ROUTER_CONFIX")
	router.PaymentConfig(r.Group(ROUTER_CONFIX))
	router.ReceiptHistory(r.Group(ROUTER_CONFIX))
	router.PosConfig(r.Group(ROUTER_CONFIX))
	router.Promotion(r.Group(ROUTER_CONFIX))
	router.ProductGroup(r.Group(ROUTER_CONFIX))

	r.Run(":8000")
}
