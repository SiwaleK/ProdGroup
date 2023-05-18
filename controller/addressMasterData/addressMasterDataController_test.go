package controllers_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "example.com/go-crud-api/controller/addressMasterData"
	mocks "example.com/go-crud-api/db/mock/addressMasterData"
	repo "example.com/go-crud-api/repository/addressMasterData"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

// สร้าง struct สำหรับ error
type CustomError struct {
	ErrorCode   int
	ErrorDetail string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("errorCode: %d, errorDetail: %s", e.ErrorCode, e.ErrorDetail)
}

func TestGetAllZipcode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns 200 on success", func(t *testing.T) {
		//ทำคล้าย ๆ ใน main เลย พวก set repo, controller, router แต่เราจะไปเรียก mockRepo แทน
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Zippcode", controller.GetAllZipcode)

		//	สร้างตัว recorder แล้วก้ยิง request ไปตาม path ที่ต้องการ
		w := httptest.NewRecorder()
		//ใส่ตัว Body ที่เราจะส่งเป็น parameter ตัวที่ 2 ได้ (ตรง nil เปลี่ยนเป็น body ที่ต้องการส่งไปแทน)
		req, _ := http.NewRequest("GET", "/sale/api/v1/AddressMasterData/Zippcode", nil)
		router.ServeHTTP(w, req)

		// Check status code ได้เช่นกัน
		assert.Equal(t, 200, w.Code)

		// Check the response body ได้ กรณีอยากจะเช็คตัว resonse ด้วย
		expectedBody := `{"items":[10000,10001,10002]}`
		assert.Equal(t, expectedBody, w.Body.String())

		//Check ต่าง ๆ ก็ได้ ก็ดูไปว่าอยากจะเช็คอะไร
	})

	t.Run("returns 500 on error", func(t *testing.T) {
		//ทำเหมือนเดิม
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		//คราวนี้ override ตัว method MockGetAllZipcodeRepository เดิม ให้มันออก error ตามที่เราจะ test
		mockRepo.MockGetAllZipcodeRepository = func() ([]int32, error) {
			return nil, errors.New("test error")
		}

		// Set up the controller using the mock repository
		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Zippcode", controller.GetAllZipcode)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/AddressMasterData/Zippcode", nil)
		router.ServeHTTP(w, req)

		// Check status code
		assert.Equal(t, 500, w.Code)

		// // Check the response body
		// expectedBody := `{"errorCode":500,"errorDetail":"test error"}`
		// assert.JSONEq(t, expectedBody, w.Body.String())
	})
}

func TestGetAllProvince(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Province", controller.GetAllProvince)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/AddressMasterData/Province", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("returns 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()

		mockRepo.MockGetAllProvinceRepository = func() ([]repo.ReadAllProvinceResp, error) {
			return []repo.ReadAllProvinceResp{}, errors.New("test error")
		}

		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Province", controller.GetAllProvince)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/AddressMasterData/Province", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetAddressByZipcode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Address", controller.GetAddressByZipcode)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/AddressMasterData/Address?zipcode=10210", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("returns 404 not found", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()

		mockRepo.MockGetAddressByZipcodeRepository = func() (repo.GetAddressByZipcodeResponse, error) {
			return repo.GetAddressByZipcodeResponse{}, nil
		}

		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Address", controller.GetAddressByZipcode)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/AddressMasterData/Address?=1021", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("returns 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()

		mockRepo.MockGetAddressByZipcodeRepository = func() (repo.GetAddressByZipcodeResponse, error) {
			return repo.GetAddressByZipcodeResponse{}, errors.New("test error")
		}

		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Address", controller.GetAddressByZipcode)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/AddressMasterData/Address?zipcode=10210", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetDistrictByProvinceName(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceName := "Mock"
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/District", controller.GetDistrictByProvinceName)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/District?provinceName=%s", mockProvinceName), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("returns 404 not found", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceName := "Unknown"
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/District", controller.GetDistrictByProvinceName)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/District?provinceName=%s", mockProvinceName), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("returns 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceName := "Mock"

		mockRepo.MockGetDistrictByProvinceNameRepository = func(provinceName string) ([]repo.GetDistrictByProvinceNameResult, error) {
			return []repo.GetDistrictByProvinceNameResult{}, errors.New("test error")
		}

		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/District", controller.GetDistrictByProvinceName)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/District?provinceName=%s", mockProvinceName), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetSubDistrictByDistrictId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockDistrictId := "1"
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/SubDistrict/:districtId", controller.GetSubDistrictByDistrictId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/SubDistrict/%s", mockDistrictId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("returns 404 not found", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockDistrictId := "0"
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/SubDistrict/:districtId", controller.GetSubDistrictByDistrictId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/SubDistrict/%s", mockDistrictId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("returns 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockDistrictId := "1"

		mockRepo.MockGetSubDistrictByDistrictIdRepository = func(districtId string) ([]repo.GetSubDistrictByDistrictIdResult, error) {
			return []repo.GetSubDistrictByDistrictIdResult{}, errors.New("test error")
		}

		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/SubDistrict/:districtId", controller.GetSubDistrictByDistrictId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/SubDistrict/%s", mockDistrictId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetDistrictByProvinceId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceId := "1"
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/District/:provinceId", controller.GetDistrictByProvinceId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/District/%s", mockProvinceId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("returns 404 not found", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceId := "0"
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/District/:provinceId", controller.GetDistrictByProvinceId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/District/%s", mockProvinceId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("returns 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceId := "1"

		mockRepo.MockGetDistrictByProvinceIdRepository = func(provinceId string) ([]repo.GetDistrictByProvinceIdResult, error) {
			return []repo.GetDistrictByProvinceIdResult{}, errors.New("test error")
		}

		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/District/:provinceId", controller.GetDistrictByProvinceId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/District/%s", mockProvinceId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}

func TestGetProvinceByProvinceId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceId := "1"
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Province/:provinceId", controller.GetProvinceByProvinceId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/Province/%s", mockProvinceId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("returns 404 not found", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceId := "0" // Invalid province ID
		controller := controllers.NewAddressMasterDataController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Province/:provinceId", controller.GetProvinceByProvinceId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/Province/%s", mockProvinceId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("returns 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockAddressMasterDataRepository()
		mockProvinceId := "1"

		mockRepo.MockGetProvinceByProvinceIdRepository = func(provinceId string) (repo.GetProvinceByProvinceIdResult, error) {
			return repo.GetProvinceByProvinceIdResult{}, errors.New("test error")
		}

		controller := controllers.NewAddressMasterDataController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/AddressMasterData/Province/:provinceId", controller.GetProvinceByProvinceId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/sale/api/v1/AddressMasterData/Province/%s", mockProvinceId), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}
