package router

import (
	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/SiwaleK/ProdGroup/pkg/sale"
	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(queries *db.Queries) *gin.Engine {

	// Initialize repositories
	//prodgroupRepo := repository.NewProdgroupRepository(queries)
	paymentMethodRepo := repository.NewPaymentMethodRepository(queries)
	//promotionRepo := repository.NewDBPromotionRepository(queries)

	// Initialize handlers
	//prodgroupHandler := sale.NewProdgroupHandler(prodgroupRepo)
	paymentMethodHandler := sale.NewPaymentMethodHandler(paymentMethodRepo)
	//promotionHandler := sale.NewPromotionHandler(promotionRepo)

	// Initialize router
	r := gin.Default()
	//r.GET("/sale/api/v1/ProductGroup", prodgroupHandler.GetProdgroup)
	r.GET("/sale/api/v1/PaymentMethod", paymentMethodHandler.GetPaymentMethod)
	//r.GET("/sale/api/v1/Promotions/:id", promotionHandler.GetPromotionByID)

	return r
}
