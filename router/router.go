package router

import (
	"database/sql"

	"github.com/SiwaleK/ProdGroup/pkg/sale"
	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(db *sql.DB) *gin.Engine {
	// Initialize repositories
	prodgroupRepo := repository.NewProdgroupRepository(db)
	paymentMethodRepo := repository.NewPaymentMethodRepository(db)
	promotionRepo := repository.NewDBPromotionRepository(db)

	// Initialize handlers
	prodgroupHandler := sale.NewProdgroupHandler(prodgroupRepo)
	paymentMethodHandler := sale.NewPaymentMethodHandler(paymentMethodRepo)
	promotionHandler := sale.NewPromotionHandler(promotionRepo)

	// Initialize router
	r := gin.Default()
	r.GET("/sale/api/v1/ProductGroup", prodgroupHandler.GetProdgroup)
	r.GET("/sale/api/v1/PaymentMethod", paymentMethodHandler.GetPaymentMethod)
	r.GET("/sale/api/v1/Promotions/:id", promotionHandler.GetPromotionByID)

	return r
}
