package router

import (
	controllers "example.com/go-crud-api/controller/addressMasterData"
	repo "example.com/go-crud-api/repository/addressMasterData"
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
