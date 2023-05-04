package router

import (
	"github.com/SiwaleK/ProdGroup/pkg/sale"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

//func RegisterRoutes(router *gin.RouterGroup) {

//saleAPI.GET("/ProductCat", sale.GetProdGroup)
//saleAPI.GET("/ProductCat/:id", sale.GetProdGroupByID)
//saleAPI.GET("/PaymentMethod", sale.GetPaymentMethod)
//saleAPI.GET("/Promotion/:id", sale.GetPromotionByID)

// Create a new instance of PromotionHandler and pass it a PromotionRepository
// 	db, err := sql.Open("driverName", "dataSourceName")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	saleAPI := router.Group("/sale/api/v1")
// 	repo := NewDBPromotionRepository(db)
// 	handler := NewPromotionHandler(repo)

// 	// Register the GetPromotionByid function with the router
// 	saleAPI.GET("/promotion/:id", handler.GetPromotionByID)

// }

func RegisterRoutes(handler *sale.PromotionHandler) *gin.Engine {
	r := gin.Default()

	promotions := r.Group("/promotions")
	{
		promotions.GET("/:id", handler.GetPromotionByID)
	}

	return r
}
