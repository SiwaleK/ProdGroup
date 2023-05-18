package controllers

import (
	"fmt"
	"net/http"

	repo "example.com/go-crud-api/repository/addressMasterData"
	"github.com/gin-gonic/gin"
)

// กำหนด struct ของ controller โดยจะมี Method signature ตาม interface ของ repo ที่เราสร้าง
type AddressMasterDataController struct {
	addressRepo repo.AddressMasterDataRepository
}

// constructtor เอาไว้สร้างก้อน controller โดยจะรับตัว interface ที่ใช้สร้าง
func NewAddressMasterDataController(addressRepo repo.AddressMasterDataRepository) *AddressMasterDataController {
	return &AddressMasterDataController{
		addressRepo: addressRepo,
	}
}

// ใส่ method ของ controllers ไป
func (ac *AddressMasterDataController) GetAllZipcode(c *gin.Context) {
	//addressRepo := repo.NewAddressMasterDataRepository()
	res, err := ac.addressRepo.GetAllZipcodeRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": res,
	})
	return
}

func (ac *AddressMasterDataController) GetAllProvince(c *gin.Context) {
	res, err := ac.addressRepo.GetAllProvinceRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": res,
	})
	return
}

func (ac *AddressMasterDataController) GetAddressByZipcode(c *gin.Context) {
	var zipCode string
	zipCode = c.Query("zipcode")
	res, err := ac.addressRepo.GetAddressByZipcodeRepository(zipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	if len(res.SubDistrictItems) == 0 {
		var err = fmt.Errorf("Zipcode not found")
		c.JSON(http.StatusNotFound, gin.H{
			"errorCode":   http.StatusNotFound,
			"errorDetail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (ac *AddressMasterDataController) GetDistrictByProvinceName(c *gin.Context) {
	var provinceName string
	provinceName = c.Query("provinceName")
	res, err := ac.addressRepo.GetDistrictByProvinceNameRepository(provinceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}
	fmt.Println(err)
	if len(res) == 0 {
		fmt.Println("res == 0")
		var err = fmt.Errorf("Province name not found")
		c.JSON(http.StatusNotFound, gin.H{
			"errorCode":   http.StatusNotFound,
			"errorDetail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": res,
	})
	return
}

func (ac *AddressMasterDataController) GetSubDistrictByDistrictId(c *gin.Context) {
	var districtId string
	districtId = c.Param("districtId")
	res, err := ac.addressRepo.GetSubDistrictByDistrictIdRepository(districtId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	if len(res) == 0 {
		var err = fmt.Errorf("District ID not found")
		c.JSON(http.StatusNotFound, gin.H{
			"errorCode":   http.StatusNotFound,
			"errorDetail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": res,
	})
	return
}

func (ac *AddressMasterDataController) GetDistrictByProvinceId(c *gin.Context) {
	var provinceId string
	provinceId = c.Param("provinceId")
	res, err := ac.addressRepo.GetDistrictByProvinceIdRepository(provinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	if len(res) == 0 {
		var err = fmt.Errorf("Province ID not found")
		c.JSON(http.StatusNotFound, gin.H{
			"errorCode":   http.StatusNotFound,
			"errorDetail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": res,
	})
	return
}

func (ac *AddressMasterDataController) GetProvinceByProvinceId(c *gin.Context) {
	var provinceId string
	provinceId = c.Param("provinceId")
	res, err := ac.addressRepo.GetProvinceByProvinceIdRepository(provinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   http.StatusInternalServerError,
			"errorDetail": err.Error(),
		})
		return
	}

	if res == (repo.GetProvinceByProvinceIdResult{}) {
		var err = fmt.Errorf("Province ID not found")
		c.JSON(http.StatusNotFound, gin.H{
			"errorCode":   http.StatusNotFound,
			"errorDetail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
