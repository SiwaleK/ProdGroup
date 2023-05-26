package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "example.com/go-crud-api/controller/productGroup"
	mocks "example.com/go-crud-api/db/mock/productGroup"
	repo "example.com/go-crud-api/repository/productGroup"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProductGroup(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("return 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockProductGroupRepository()
		controller := controllers.NewProductGroupController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/ProductGroup", controller.GetProductGroup)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ProductGroup", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		expectedBody :=
			`[{"prodgroupid":1,"th_name":"ไม่ได้จัดหมวดหมู่","en_name":"No Group"},{"prodgroupid":2,"th_name":"เครื่องดื่มแอลกอฮอล์","en_name":"Alcoholic Beverage"}]`
		assert.Equal(t, expectedBody, w.Body.String())
	})
	t.Run("return 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockProductGroupRepository()
		mockRepo.MockGetProductGroup = func() ([]repo.Prodgroup, error) {
			return nil, errors.New("mocked error")
		}
		controller := controllers.NewProductGroupController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/ProductGroup", controller.GetProductGroup)
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ProductGroup", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})
}

func TestGetProductGroupByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("return 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockProductGroupRepository()
		mockprodgroupid := "1"
		expectedResult := []repo.Prodgroup{
			{
				Prodgroupid: 1,
				ThName:      "ไม่ได้จัดหมวดหมู่",
				EnName:      "No Group",
			},
		}

		mockRepo.MockGetProductGroupByID = func(prodGroupID int) ([]repo.Prodgroup, error) {
			return expectedResult, nil
		}

		controller := controllers.NewProductGroupController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/ProductGroupByID/:prodgroupID", controller.GetProductGroupByID)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ProductGroupByID/"+mockprodgroupid, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response []repo.Prodgroup
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, response)
	})

	t.Run("return 400 on invalid prodgroupID", func(t *testing.T) {
		mockRepo := mocks.NewMockProductGroupRepository()
		mockprodgroupid := "invalid"
		controller := controllers.NewProductGroupController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/ProductGroupByID/:prodgroupID", controller.GetProductGroupByID)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ProductGroupByID/"+mockprodgroupid, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedBody := `{"errorCode":400,"errorDetail":"Invalid prodgroupID"}`
		assert.Equal(t, expectedBody, w.Body.String())
	})

	t.Run("return 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockProductGroupRepository()
		mockprodgroupid := "1"
		expectedError := errors.New("mocked error")

		mockRepo.MockGetProductGroupByID = func(prodGroupID int) ([]repo.Prodgroup, error) {
			return nil, expectedError
		}

		controller := controllers.NewProductGroupController(mockRepo)

		router := gin.Default()
		router.GET("/sale/api/v1/ProductGroupByID/:prodgroupID", controller.GetProductGroupByID)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ProductGroupByID/"+mockprodgroupid, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		expectedBody := `{"errorCode":500,"errorDetail":"mocked error"}`
		assert.Equal(t, expectedBody, w.Body.String())
	})
}
