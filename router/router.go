package router

import (
	payment_config "github.com/SiwaleK/ProdGroup/controller/payment_config"
	payment_method "github.com/SiwaleK/ProdGroup/controller/payment_method"
	prodgroup "github.com/SiwaleK/ProdGroup/controller/prodgroup"
	promotion "github.com/SiwaleK/ProdGroup/controller/promotion"
	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(queries *db.Queries) *gin.Engine {

	// Initialize repositories
	prodgroupRepo := repository.NewProdgroupRepository(queries)
	paymentmethodRepo := repository.NewPaymentMethodRepository(queries)
	promotionRepo := repository.NewPromotionRepository(queries)
	paymentconfigRepo := repository.NewPaymentConfigRepository(queries)

	// Initialize handlers
	prodgroupHandler := prodgroup.NewProdgroupHandler(prodgroupRepo)
	paymentmethodHandler := payment_method.NewPaymentMethodHandler(paymentmethodRepo)
	promotionHandler := promotion.NewPromotionHandler(promotionRepo)
	paymentconfigHandler := payment_config.NewPaymentConfighandler(paymentconfigRepo)

	// Initialize router
	r := gin.Default()
	r.GET("/sale/api/v1/ProductGroup", prodgroupHandler.GetProdgroup)
	r.GET("/sale/api/v1/ProductGroup/:id", prodgroupHandler.GetProdgroupByID)
	r.GET("/sale/api/v1/PaymentMethod", paymentmethodHandler.GetPaymentMethod)
	//r.GET("/sale/api/v1/Promotions/:id", promotionHandler.GetPromotionByID)

	//Promotion
	r.POST("/sale/api/v1/Promotion/DiscountPerItem", promotionHandler.PostDiscountPromotion)
	r.POST("/sale/api/v1/Promotion/AFreeB", promotionHandler.PostPromotionAFREEB)
	r.POST("/sale/api/v1/Promotion/StepPurchase", promotionHandler.PostPromotionStepPurchase)

	// PaymentConfig

	r.GET("/sale/api/v1/PaymentConfig", paymentconfigHandler.GetPosClientMethod)

	return r
}
