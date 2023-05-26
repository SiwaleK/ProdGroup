package controller_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	errortype "example.com/go-crud-api/common/errorType"
	controllers "example.com/go-crud-api/controller/receiptHistory"
	mocks "example.com/go-crud-api/db/mock/receiptHistory"
	repo "example.com/go-crud-api/repository/receiptHistory"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type CustomError struct {
	ErrorCode   int
	ErrorDetail string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("errorCode: %d, errorDetail: %s", e.ErrorCode, e.ErrorDetail)
}

func TestGetReceiptHistoryByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("return 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		controller := controllers.NewReceiptHistoryController(mockRepo)

		// Create a new Gin router
		router := gin.Default()

		// Define the route and handler function
		router.GET("/sale/api/v1/ReceiptHistoryByID", controller.GetReceiptHistoryByID)

		// Create a test request without the query parameter
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByID", nil)

		// Add the query parameter to the request URL
		q := req.URL.Query()
		q.Add("saleOrderID", "123")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()

		// Send the request to the router
		router.ServeHTTP(w, req)

		// Assert the response status code
		assert.Equal(t, http.StatusOK, w.Code)

	})
	t.Run("returns 400 no parameter", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		controller := controllers.NewReceiptHistoryController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/ReceiptHistoryByID", controller.GetReceiptHistoryByID)
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByID", nil)
		q := req.URL.Query()
		q.Add("saleOrderID", "")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, errortype.ErrorCode(w.Code), errortype.ErrorCode(400))

	})

	t.Run("return 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		mockRepo.MockGetRecieptHistoryByID = func(saleOrderID string) ([]repo.ReceiptHistory, error) {
			return nil, errors.New("mocked error")
		}
		controller := controllers.NewReceiptHistoryController(mockRepo)

		// Create a new Gin router
		router := gin.Default()

		// Define the route and handler function
		router.GET("/sale/api/v1/ReceiptHistoryByID", controller.GetReceiptHistoryByID)

		// Create a test request with the query parameter
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByID?saleOrderID=123", nil)

		w := httptest.NewRecorder()

		// Send the request to the router
		router.ServeHTTP(w, req)

		// Assert the response status code
		assert.Equal(t, errortype.ErrorCode(w.Code), errortype.ErrorCode(500))

	})

}

func TestGetReceiptHistoryByDate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("return 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		controller := controllers.NewReceiptHistoryController(mockRepo)

		// Create a new Gin router
		router := gin.Default()

		// Define the route and handler function
		router.GET("/sale/api/v1/ReceiptHistoryByDate", controller.GetReceiptHistoryByDate)

		// Create a test request without the query parameter
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByDate", nil)

		// Add the query parameter to the request URL
		q := req.URL.Query()

		q.Set("startDate", "2023-05-10")
		q.Set("endDate", "2023-05-11")

		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()

		// Send the request to the router
		router.ServeHTTP(w, req)

		// Assert the response status code
		assert.Equal(t, http.StatusOK, w.Code)

	})
	t.Run("returns 400 no parameter", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		controller := controllers.NewReceiptHistoryController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/ReceiptHistoryByDate", controller.GetReceiptHistoryByDate)
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByDate", nil)
		q := req.URL.Query()
		q.Set("startDate", "")
		q.Set("endDate", "")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, errortype.ErrorCode(w.Code), errortype.ErrorCode(400))

	})

	t.Run("return 500 on error", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		mockRepo.MockGetRecieptHistoryByDate = func(startDate, endDate time.Time) ([]repo.ReceiptHistory, error) {
			return nil, errors.New("mocked error")
		}
		controller := controllers.NewReceiptHistoryController(mockRepo)

		// Create a new Gin router
		router := gin.Default()

		// Define the route and handler function
		router.GET("/sale/api/v1/ReceiptHistoryByDate", controller.GetReceiptHistoryByDate)

		// Create a test request with the query parameter
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByDate?startDate=2023-05-10&endDate=2023-05-11", nil)

		w := httptest.NewRecorder()

		// Send the request to the router
		router.ServeHTTP(w, req)

		// Assert the response status code
		assert.Equal(t, errortype.ErrorCode(w.Code), errortype.ErrorCode(500))

	})
	t.Run("returns 400 invalid", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		controller := controllers.NewReceiptHistoryController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/ReceiptHistoryByDate", controller.GetReceiptHistoryByDate)
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByDate", nil)
		q := req.URL.Query()
		q.Set("startDate", "2023/05/11")
		q.Set("endDate", "2023/05/11")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, errortype.ErrorCode(w.Code), errortype.ErrorCode(400))

	})
	t.Run("returns 400 invalid", func(t *testing.T) {
		mockRepo := mocks.NewMockRecieptHistoryRepository()
		controller := controllers.NewReceiptHistoryController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/ReceiptHistoryByDate", controller.GetReceiptHistoryByDate)
		req, _ := http.NewRequest(http.MethodGet, "/sale/api/v1/ReceiptHistoryByDate", nil)
		q := req.URL.Query()
		q.Set("startDate", "2023-05-10")
		q.Set("endDate", "2023/05/11")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, errortype.ErrorCode(w.Code), errortype.ErrorCode(400))

	})

}
