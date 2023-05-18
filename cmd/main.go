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

	//loadPrefix จาก env
	ADDRESSMASTERDATA_PREFIX := os.Getenv("ADDRESSMASTERDATA_PREFIX")
	//สร้าง router
	router.AddressMasterDataRoutes(r.Group(ADDRESSMASTERDATA_PREFIX))

	r.Run(":8000")
}
