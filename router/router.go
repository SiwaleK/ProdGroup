package router

import (
	controllers "example.com/go-crud-api/controller/addressMasterData"
	paymentConfigController "example.com/go-crud-api/controller/paymentConfig"
	posConfigCon "example.com/go-crud-api/controller/posConfig"
	productGroupCon "example.com/go-crud-api/controller/productGroup"
	promotionCon "example.com/go-crud-api/controller/promotion"
	receiptCon "example.com/go-crud-api/controller/receiptHistory"
	repo "example.com/go-crud-api/repository/addressMasterData"
	paymentConfigRepo "example.com/go-crud-api/repository/paymentConfig"
	posConfigRepo "example.com/go-crud-api/repository/posConfig"
	productGroupRepo "example.com/go-crud-api/repository/productGroup"
	promotionRepo "example.com/go-crud-api/repository/promotion"
	receiptRepo "example.com/go-crud-api/repository/receiptHistory"
	"github.com/gin-gonic/gin"
)

func AddressMasterDataRoutes(addressMasterDataAPI *gin.RouterGroup) {
	//สร้าง repo แล้วก็สร้าง controller ด้วย repo ที่สร้างขึ้นมา (dependency injection)
	var addressMasterDataRepo repo.AddressMasterDataRepository
	addressMasterDataRepo = repo.NewAddressMasterDataRepository()
	addressMasterDataController := controllers.NewAddressMasterDataController(addressMasterDataRepo)

	//กำหนด endpoint แล้วก็ function ที่จะใช้
	addressMasterDataAPI.GET("/Zippcode", addressMasterDataController.GetAllZipcode)
	addressMasterDataAPI.GET("/Province", addressMasterDataController.GetAllProvince)
	addressMasterDataAPI.GET("/Address", addressMasterDataController.GetAddressByZipcode)
	addressMasterDataAPI.GET("/District", addressMasterDataController.GetDistrictByProvinceName)
	addressMasterDataAPI.GET("/Subdistrict/:districtId", addressMasterDataController.GetSubDistrictByDistrictId)
	addressMasterDataAPI.GET("/District/:provinceId", addressMasterDataController.GetDistrictByProvinceId)
	addressMasterDataAPI.GET("/Province/:provinceId", addressMasterDataController.GetProvinceByProvinceId)
}

func PaymentConfig(paymentConfig *gin.RouterGroup) {
	paymentConfigRepo := paymentConfigRepo.NewPaymentConfigRepository()
	paymentConfigController := paymentConfigController.NewPaymentConfigController(paymentConfigRepo)

	paymentConfig.GET("/PaymentConfig", paymentConfigController.GetPaymentConfig)
}

func ReceiptHistory(receipthistory *gin.RouterGroup) {
	receipthistoryRepo := receiptRepo.NewReceiptHistoryRepository()
	receipthistoryController := receiptCon.NewReceiptHistoryController(receipthistoryRepo)

	receipthistory.GET("/ReceiptHistoryByID", receipthistoryController.GetReceiptHistoryByID)
	receipthistory.GET("/ReceiptHistoryByDate", receipthistoryController.GetReceiptHistoryByDate)

}

func PosConfig(posConfig *gin.RouterGroup) {
	posConfigRepo := posConfigRepo.NewPosConfigRepository()
	posConfigController := posConfigCon.NewPosConfigController(posConfigRepo)
	posConfig.GET("/POSConfig", posConfigController.GetPosConfig)
}

func Promotion(promotion *gin.RouterGroup) {
	promotionRepo := promotionRepo.NewPromotionRepository()
	promotionController := promotionCon.NewPromotionController(promotionRepo)
	promotion.POST("/Promotion/DiscountPerItem", promotionController.CreatePromotionDiscount)
	promotion.POST("/Promotion/AFreeB", promotionController.CreatePromotionAFREEB)
	promotion.POST("/Promotion/StepPurchase", promotionController.CreatePromotionStepPurchase)
}

func ProductGroup(productGroup *gin.RouterGroup) {
	productGroupRepo := productGroupRepo.NewProductGroupRepository()
	productGroupController := productGroupCon.NewProductGroupController(productGroupRepo)
	productGroup.GET("/ProductGroup", productGroupController.GetProductGroup)
	productGroup.GET("/ProductGroup/:prodgroupID", productGroupController.GetProductGroupByID)
}
