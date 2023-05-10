package router

import (
	"github.com/SiwaleK/ProdGroup/controller"
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
	prodgroupHandler := controller.NewProdgroupHandler(prodgroupRepo)
	paymentmethodHandler := controller.NewPaymentMethodHandler(paymentmethodRepo)
	promotionHandler := controller.NewPromotionHandler(promotionRepo)
	paymentconfigHandler := controller.NewPaymentConfighandler(paymentconfigRepo)

	// Initialize router
	r := gin.Default()
	r.GET("/sale/api/v1/ProductGroup", prodgroupHandler.GetProdgroup)
	r.GET("/sale/api/v1/ProductGroup/:id", prodgroupHandler.GetProdgroupByID)
	r.GET("/sale/api/v1/PaymentMethod", paymentmethodHandler.GetPaymentMethod)
	//r.GET("/sale/api/v1/Promotions/:id", promotionHandler.GetPromotionByID)

	//Promotion
	r.POST("/sale/api/v1/Promotion/DiscountPerItem", promotionHandler.PostDiscountPromotion)

	// PaymentConfig
	r.GET("/sale/api/v1/PaymentConfig", paymentconfigHandler.GetPaymentConfig)
	r.POST("/sale/api/v1/PaymentConfig", paymentconfigHandler.UpsertPaymentConfigHandler)

	return r
}
