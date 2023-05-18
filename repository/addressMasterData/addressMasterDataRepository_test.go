package repo_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	mocks "example.com/go-crud-api/db/mock"
	repo "example.com/go-crud-api/repository/addressMasterData"
)

func TestGetAllZipcodeRepository(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		//รอปิด mock database connection หลังจากจบ function ด้วย
		defer sqlDB.Close()

		//สร้างตัว expected result
		expectedResults := []int32{11111, 22222, 33333}
		//สร้าง rows ที่มี struct ตามที่ต้องการ
		rows := sqlmock.NewRows([]string{"id", "zipcode"})
		//insert expected result เข้าไปใน mock database
		for i, result := range expectedResults {
			rows.AddRow(i+1, result)
		}
		//บอกว่าถ้าเจอ query นี้ จะ return rows ออกมา (เป็น regex อันนี้คือให้มัน accept ทุก query ไปก่อน)
		mock.ExpectQuery(".+").WillReturnRows(rows)

		//แล้วก็เช็ค
		addressRepo := repo.NewAddressMasterDataRepository()
		result, err := addressRepo.GetAllZipcodeRepository()

		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		//สร้าง errors ขึ้นมาใหม่
		expectedError := errors.New("some database error")
		//เปลี่ยนตรงนี้แทนว่าถ้าเจอ query ยิงมาให้คืน error
		mock.ExpectQuery(".+").WillReturnError(expectedError)

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetAllZipcodeRepository()

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetAllProvinceRepository(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		expectedResults := []repo.ReadAllProvinceResp{
			{ProvinceId: 1, ProvinceName: "Province1"},
			{ProvinceId: 2, ProvinceName: "Province2"},
			{ProvinceId: 3, ProvinceName: "Province3"},
		}
		rows := sqlmock.NewRows([]string{"provinceid", "provincename"})
		for _, result := range expectedResults {
			rows.AddRow(result.ProvinceId, result.ProvinceName)
		}
		mock.ExpectQuery(".+").WillReturnRows(rows)

		addressRepo := repo.NewAddressMasterDataRepository()
		result, err := addressRepo.GetAllProvinceRepository()

		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		expectedError := errors.New("some database error")
		mock.ExpectQuery(".+").WillReturnError(expectedError)

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetAllProvinceRepository()

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetAddressByZipcodeRepository(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		zipCode := "12345"

		expectedResult := repo.GetAddressByZipcodeResponse{
			SubDistrictItems: []repo.SubDistrictItem{
				{SubDistrictId: "1", SubDistrictName: "SubDistrict1"},
				{SubDistrictId: "2", SubDistrictName: "SubDistrict2"},
				{SubDistrictId: "3", SubDistrictName: "SubDistrict3"},
			},
			Districtid:   1,
			Districtname: "District1",
			Provinceid:   1,
			Provincename: "Province1",
		}

		rows := sqlmock.NewRows([]string{"subdistrictid", "subdistrictname", "districtid", "districtname", "provinceid", "provincename"})
		for _, subDistrict := range expectedResult.SubDistrictItems {
			rows.AddRow(subDistrict.SubDistrictId, subDistrict.SubDistrictName, expectedResult.Districtid, expectedResult.Districtname, expectedResult.Provinceid, expectedResult.Provincename)
		}

		mock.ExpectQuery(".+").WillReturnRows(rows)

		addressRepo := repo.NewAddressMasterDataRepository()
		result, err := addressRepo.GetAddressByZipcodeRepository(zipCode)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("invalid zip code", func(t *testing.T) {
		db, _ := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		zipCode := "invalid"

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetAddressByZipcodeRepository(zipCode)

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("invalid id: %s", zipCode), err)
	})

	t.Run("no results", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		zipCode := "12345"

		mock.ExpectQuery(".+").WillReturnRows(sqlmock.NewRows([]string{"subdistrictid", "subdistrictname", "districtid", "districtname", "provinceid", "provincename"}))

		addressRepo := repo.NewAddressMasterDataRepository()
		result, err := addressRepo.GetAddressByZipcodeRepository(zipCode)

		assert.Nil(t, err)
		assert.Equal(t, repo.GetAddressByZipcodeResponse{}, result)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		zipCode := "12345"

		expectedError := errors.New("some database error")
		mock.ExpectQuery(".+").WillReturnError(expectedError)

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetAddressByZipcodeRepository(zipCode)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetDistrictByProvinceNameRepository(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceName := "Province1"

		expectedResults := []repo.GetDistrictByProvinceNameResult{
			{Districtid: 1, Districtname: "District1"},
			{Districtid: 2, Districtname: "District2"},
			{Districtid: 3, Districtname: "District3"},
		}

		rows := sqlmock.NewRows([]string{"districtid", "districtname"})
		for _, result := range expectedResults {
			rows.AddRow(result.Districtid, result.Districtname)
		}

		mock.ExpectQuery(".+").WillReturnRows(rows)

		addressRepo := repo.NewAddressMasterDataRepository()
		result, err := addressRepo.GetDistrictByProvinceNameRepository(provinceName)

		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceName := "Province1"

		expectedError := errors.New("some database error")
		mock.ExpectQuery(".+").WillReturnError(expectedError)

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetDistrictByProvinceNameRepository(provinceName)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetSubDistrictByDistrictIdRepository(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		districtId := "1"

		expectedResults := []repo.GetSubDistrictByDistrictIdResult{
			{Subdistrictid: "1", Subdistrictname: "SubDistrict1"},
			{Subdistrictid: "2", Subdistrictname: "SubDistrict2"},
			{Subdistrictid: "3", Subdistrictname: "SubDistrict3"},
		}

		rows := sqlmock.NewRows([]string{"subdistrictid", "subdistrictname"})
		for _, result := range expectedResults {
			rows.AddRow(result.Subdistrictid, result.Subdistrictname)
		}

		mock.ExpectQuery(".+").WillReturnRows(rows)

		addressRepo := repo.NewAddressMasterDataRepository()
		results, err := addressRepo.GetSubDistrictByDistrictIdRepository(districtId)

		assert.Nil(t, err)
		assert.Equal(t, expectedResults, results)
	})

	t.Run("invalid district id", func(t *testing.T) {
		db, _ := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		districtId := "invalid"

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetSubDistrictByDistrictIdRepository(districtId)

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("invalid id: %s", districtId), err)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		districtId := "1"
		expectedError := errors.New("some database error")

		mock.ExpectQuery(".+").WillReturnError(expectedError)

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetSubDistrictByDistrictIdRepository(districtId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetDistrictByProvinceIdRepository(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceId := "1"
		expectedResults := []repo.GetDistrictByProvinceIdResult{
			{Districtid: 1, Districtname: "District1"},
			{Districtid: 2, Districtname: "District2"},
			{Districtid: 3, Districtname: "District3"},
		}

		rows := sqlmock.NewRows([]string{"districtid", "districtname"})
		for _, result := range expectedResults {
			rows.AddRow(result.Districtid, result.Districtname)
		}

		mock.ExpectQuery(".+").WillReturnRows(rows)

		addressRepo := repo.NewAddressMasterDataRepository()
		result, err := addressRepo.GetDistrictByProvinceIdRepository(provinceId)

		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)
	})

	t.Run("invalid province id", func(t *testing.T) {
		db, _ := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceId := "invalid"

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetDistrictByProvinceIdRepository(provinceId)

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("invalid id: %s", provinceId), err)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceId := "1"
		expectedError := errors.New("some database error")
		mock.ExpectQuery(".+").WillReturnError(expectedError)

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetDistrictByProvinceIdRepository(provinceId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetProvinceByProvinceIdRepository(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceId := "1"
		expectedResult := repo.GetProvinceByProvinceIdResult{
			Provinceid:   1,
			Provincename: "Province1",
		}

		rows := sqlmock.NewRows([]string{"provinceid", "provincename"})
		rows.AddRow(expectedResult.Provinceid, expectedResult.Provincename)

		mock.ExpectQuery(".+").WillReturnRows(rows)

		addressRepo := repo.NewAddressMasterDataRepository()
		result, err := addressRepo.GetProvinceByProvinceIdRepository(provinceId)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("invalid province id", func(t *testing.T) {
		db, _ := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceId := "invalid"

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetProvinceByProvinceIdRepository(provinceId)

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("invalid id: %s", provinceId), err)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		provinceId := "1"
		expectedError := errors.New("some database error")
		mock.ExpectQuery(".+").WillReturnError(expectedError)

		addressRepo := repo.NewAddressMasterDataRepository()
		_, err := addressRepo.GetProvinceByProvinceIdRepository(provinceId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
