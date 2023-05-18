package repo

import (
	"fmt"
	"strconv"

	"example.com/go-crud-api/db/database"
	db "example.com/go-crud-api/model/db"
)

// สร้าง signature ให้ interface ของ repository ว่าจะมี method อะไร รับค่าอะไร return อะไร
type AddressMasterDataRepository interface {
	GetAllZipcodeRepository() ([]int32, error)
	GetAllProvinceRepository() ([]ReadAllProvinceResp, error)
	GetAddressByZipcodeRepository(zipcode string) (GetAddressByZipcodeResponse, error)
	GetDistrictByProvinceNameRepository(provinceName string) ([]GetDistrictByProvinceNameResult, error)
	GetSubDistrictByDistrictIdRepository(districtId string) ([]GetSubDistrictByDistrictIdResult, error)
	GetDistrictByProvinceIdRepository(provinceId string) ([]GetDistrictByProvinceIdResult, error)
	GetProvinceByProvinceIdRepository(provinceId string) (GetProvinceByProvinceIdResult, error)
}

// กำหนด struct สำหรับ repo มา
type addressMasterDataRepository struct{}

// function เอาไว้สร้าง repo ใหม่
func NewAddressMasterDataRepository() AddressMasterDataRepository {
	return &addressMasterDataRepository{}
}

// implements ไอตัว methods ตาม interface ที่เรากำหนดไว้ข้างบน
func (r *addressMasterDataRepository) GetAllZipcodeRepository() ([]int32, error) {
	var response []int32
	var results []db.Subdistrict

	if err := database.DB.Table("subdistricts").
		Select("zipcode").
		Find(&results).Error; err != nil {
		return []int32{}, err
	}

	for _, subdistrict := range results {
		response = append(response, subdistrict.Zipcode)
	}

	return response, nil
}

type ReadAllProvinceResp struct {
	ProvinceId   int32  `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
}

func (r *addressMasterDataRepository) GetAllProvinceRepository() ([]ReadAllProvinceResp, error) {
	var response []ReadAllProvinceResp
	var results []db.Province
	if err := database.DB.Table("provinces").
		Select("provinceid", "provincename").
		Find(&results).Error; err != nil {
		return []ReadAllProvinceResp{}, err
	}
	for _, province := range results {
		tmp := ReadAllProvinceResp{
			ProvinceId:   province.Provinceid,
			ProvinceName: province.Provincename,
		}
		response = append(response, tmp)
	}
	return response, nil
}

type SubDistrictItem struct {
	SubDistrictId   string `json:"subDistrictId"`
	SubDistrictName string `json:"subDistrictName"`
}

type GetAddressByZipcodeResponse struct {
	SubDistrictItems []SubDistrictItem
	Districtid       int32  `json:"districtId"`
	Districtname     string `json:"districtName"`
	Provinceid       int32  `json:"provinceId"`
	Provincename     string `json:"provinceName"`
}

type GetAddressByZipcodeResult struct {
	Subdistrictid   string `gorm:"subdistrictid"`
	Subdistrictname string `gorm:"subdistrictname"`
	Districtid      int32  `gorm:"districtid"`
	Districtname    string `gorm:"districtname"`
	Provinceid      int32  `gorm:"provinceid"`
	Provincename    string `gorm:"provincename"`
}

func (r *addressMasterDataRepository) GetAddressByZipcodeRepository(zipCodeString string) (GetAddressByZipcodeResponse, error) {
	var results []GetAddressByZipcodeResult
	var subDistrictItems []SubDistrictItem

	zipCode, err := strconv.Atoi(zipCodeString)
	if err != nil {
		return GetAddressByZipcodeResponse{}, fmt.Errorf("invalid id: %s", zipCodeString)
	}

	var response GetAddressByZipcodeResponse
	if err := database.DB.Table("subdistricts").
		Select("DISTINCT subdistricts.subdistrictid, subdistricts.subdistrictname, districts.districtid, districts.districtname, provinces.provinceid, provinces.provincename").
		Joins("INNER JOIN districts ON districts.districtId = subdistricts.districtId").
		Joins("INNER JOIN provinces ON provinces.provinceId = districts.provinceId").
		Where("subdistricts.zipCode = ?", zipCode).
		Find(&results).
		Error; err != nil {
		return GetAddressByZipcodeResponse{}, err
	}

	if len(results) == 0 {
		return GetAddressByZipcodeResponse{}, nil
	}

	for _, subDistrict := range results {
		subDistrictItem := SubDistrictItem{
			SubDistrictId:   subDistrict.Subdistrictid,
			SubDistrictName: subDistrict.Subdistrictname,
		}
		subDistrictItems = append(subDistrictItems, subDistrictItem)
	}
	response.SubDistrictItems = subDistrictItems
	response.Districtid = results[0].Districtid
	response.Districtname = results[0].Districtname
	response.Provinceid = results[0].Provinceid
	response.Provincename = results[0].Provincename

	return response, nil
}

type GetDistrictByProvinceNameResult struct {
	Districtid   int32  `json:"districtId"`
	Districtname string `json:"districtName"`
}

func (r *addressMasterDataRepository) GetDistrictByProvinceNameRepository(provinceName string) ([]GetDistrictByProvinceNameResult, error) {
	var results []GetDistrictByProvinceNameResult
	if err := database.DB.Table("provinces").
		Select("districts.districtid, districts.districtname").
		Joins("INNER JOIN districts ON districts.provinceId = provinces.provinceId").
		Where("provinces.provinceName = ?", provinceName).
		Find(&results).
		Error; err != nil {
		return []GetDistrictByProvinceNameResult{}, err
	}

	return results, nil
}

type GetSubDistrictByDistrictIdResult struct {
	Subdistrictid   string `json:"id"`
	Subdistrictname string `json:"subDistrictName"`
}

func (r *addressMasterDataRepository) GetSubDistrictByDistrictIdRepository(districtIdString string) ([]GetSubDistrictByDistrictIdResult, error) {
	var results []GetSubDistrictByDistrictIdResult

	districtId, err := strconv.Atoi(districtIdString)
	if err != nil {
		return []GetSubDistrictByDistrictIdResult{}, fmt.Errorf("invalid id: %s", districtIdString)
	}

	if err := database.DB.Table("districts").
		Select("subdistricts.subdistrictid, subdistricts.subdistrictname").
		Joins("INNER JOIN subdistricts ON districts.districtId = subdistricts.districtId").
		Where("districts.districtId = ?", districtId).
		Find(&results).
		Error; err != nil {
		return []GetSubDistrictByDistrictIdResult{}, err
	}

	return results, nil
}

type GetDistrictByProvinceIdResult struct {
	Districtid   int32  `json:"id"`
	Districtname string `json:"districtName"`
}

func (r *addressMasterDataRepository) GetDistrictByProvinceIdRepository(provinceIdString string) ([]GetDistrictByProvinceIdResult, error) {
	var results []GetDistrictByProvinceIdResult

	provinceId, err := strconv.Atoi(provinceIdString)
	if err != nil {
		return []GetDistrictByProvinceIdResult{}, fmt.Errorf("invalid id: %s", provinceIdString)
	}

	if err := database.DB.Table("provinces").
		Select("districts.districtid, districts.districtname").
		Joins("INNER JOIN districts ON districts.provinceId = provinces.provinceId").
		Where("provinces.provinceId = ?", provinceId).
		Find(&results).
		Error; err != nil {
		return []GetDistrictByProvinceIdResult{}, err
	}

	return results, nil
}

type GetProvinceByProvinceIdResult struct {
	Provinceid   int32  `json:"id"`
	Provincename string `json:"provinceName"`
}

func (r *addressMasterDataRepository) GetProvinceByProvinceIdRepository(provinceIdString string) (GetProvinceByProvinceIdResult, error) {
	var results GetProvinceByProvinceIdResult

	provinceId, err := strconv.Atoi(provinceIdString)
	if err != nil {
		return GetProvinceByProvinceIdResult{}, fmt.Errorf("invalid id: %s", provinceIdString)
	}

	if err := database.DB.Table("provinces").
		Select("provinces.provinceid, provinces.provincename").
		Where("provinces.provinceId = ?", provinceId).
		Find(&results).Error; err != nil {

		return GetProvinceByProvinceIdResult{}, err

	}

	return results, nil
}
