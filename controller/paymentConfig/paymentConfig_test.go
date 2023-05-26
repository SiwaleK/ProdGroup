package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controllers "example.com/go-crud-api/controller/paymentConfig"
	mocks "example.com/go-crud-api/db/mock/paymentConfig"
	repo "example.com/go-crud-api/repository/paymentConfig"

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

// version อยู่ใน config
func TestGetPaymentConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("return 200 on success", func(t *testing.T) {
		mockRepo := mocks.NewMockPaymentConfigRepository()

		controller := controllers.NewPaymentConfigController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/PaymentConfig", controller.GetPaymentConfig)
		requestBody := []byte(`{"posClientID": "value"}`)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfig", bytes.NewBuffer(requestBody))
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
	t.Run("returns 400 invalid json", func(t *testing.T) {
		mockRepo := mocks.NewMockPaymentConfigRepository()

		controller := controllers.NewPaymentConfigController(mockRepo)
		router := gin.Default()
		router.GET("/sale/api/v1/PaymentConfig", controller.GetPaymentConfig)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfig", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
	t.Run("returns 500 on Posclient error", func(t *testing.T) {
		mockRepo := mocks.NewMockPaymentConfigRepository()

		mockRepo.MockGetPosClientConfig = func(posClientID string) ([]repo.PosclientConfig, error) {
			// Simulate an error
			return []repo.PosclientConfig{}, errors.New("test error")
		}
		controller := controllers.NewPaymentConfigController(mockRepo)
		requestBody := []byte(`{"posClientID": "value"}`)
		// Create a new Gin router and register the controller's handler
		router := gin.Default()
		router.GET("/sale/api/v1/PaymentConfig", controller.GetPaymentConfig)

		// Perform the HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfig", bytes.NewBuffer(requestBody))
		router.ServeHTTP(w, req)

		// Assert the response status code and body
		assert.Equal(t, 500, w.Code)

	})
	t.Run("returns 500 on PaymentMethod error", func(t *testing.T) {
		mockRepo := mocks.NewMockPaymentConfigRepository()

		mockRepo.MockGetPaymentMethod = func() ([]repo.PaymentMethod, error) {
			// Simulate an error
			return []repo.PaymentMethod{}, errors.New("test error")
		}
		controller := controllers.NewPaymentConfigController(mockRepo)
		requestBody := []byte(`{"posClientID": "value"}`)
		// Create a new Gin router and register the controller's handler
		router := gin.Default()
		router.GET("/sale/api/v1/PaymentConfig", controller.GetPaymentConfig)
		// version อยู่ใน config

		// Perform the HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfig", bytes.NewBuffer(requestBody))
		router.ServeHTTP(w, req)

		// Assert the response status code and body
		assert.Equal(t, 500, w.Code)

	})

	t.Run("returns 500", func(t *testing.T) {
		mockRepo := mocks.NewMockPaymentConfigRepository()

		mockRepo.MockGetPaymentMethod = func() ([]repo.PaymentMethod, error) {
			// Simulate an error
			return []repo.PaymentMethod{}, errors.New("test error")
		}
		controller := controllers.NewPaymentConfigController(mockRepo)
		requestBody := []byte(`{"posClientID": "value"}`)

		// Create a new Gin router and register the controller's handler
		router := gin.Default()
		router.GET("/sale/api/v1/PaymentConfig", controller.GetPaymentConfig)

		// Perform the HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfig", bytes.NewBuffer(requestBody))
		router.ServeHTTP(w, req)

		// Check the response status code
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, w.Code)
		}

		// Check the response headers
		expectedContentType := "application/json"
		actualContentType := w.Header().Get("Content-Type")
		if !strings.Contains(actualContentType, expectedContentType) {
			t.Errorf("Expected Content-Type header %s, but got %s", expectedContentType, actualContentType)
		}
		// Check the response body
		expectedBody := `{"errorCode":500,"errorDetail":"test error"}`
		if w.Body.String() != expectedBody {
			t.Errorf("Expected response body %s, but got %s", expectedBody, w.Body.String())
		}
	})

	t.Run("returns 500", func(t *testing.T) {
		mockRepo := mocks.NewMockPaymentConfigRepository()

		mockRepo.MockGetPaymentMethod = func() ([]repo.PaymentMethod, error) {
			// Simulate an error
			return []repo.PaymentMethod{}, errors.New("test error")
		}
		controller := controllers.NewPaymentConfigController(mockRepo)
		requestBody := []byte(`{"posClientID": "value"}`)

		// Create a new Gin router and register the controller's handler
		router := gin.Default()
		router.GET("/sale/api/v1/PaymentConfig", controller.GetPaymentConfig)

		// Perform the HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfig", bytes.NewBuffer(requestBody))
		router.ServeHTTP(w, req)

		// Assertions
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, w.Code)
		}

		expectedContentType := "application/json"
		actualContentType := w.Header().Get("Content-Type")
		if !strings.Contains(actualContentType, expectedContentType) {
			t.Errorf("Expected Content-Type header %s, but got %s", expectedContentType, actualContentType)
		}

		expectedBody := `{"errorCode":500,"errorDetail":"test error"}`
		if w.Body.String() != expectedBody {
			t.Errorf("Expected response body %s, but got %s", expectedBody, w.Body.String())
		}

		// Mock the gin.Context and call the handler directly
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		controller.GetPaymentConfig(ctx)

		actualContentType = ctx.Writer.Header().Get("Content-Type")
		if !strings.Contains(actualContentType, expectedContentType) {
			t.Errorf("Expected Content-Type header %s, but got %s", expectedContentType, actualContentType)
		}

		// Unmarshal the response body to check the error
		var response struct {
			ErrorCode   int    `json:"errorCode"`
			ErrorDetail string `json:"errorDetail"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Failed to unmarshal response body: %s", err.Error())
		}

		if response.ErrorCode != http.StatusInternalServerError {
			t.Errorf("Expected error code %d, but got %d", http.StatusInternalServerError, response.ErrorCode)
		}

		if response.ErrorDetail != "test error" {
			t.Errorf("Expected error detail %s, but got %s", "test error", response.ErrorDetail)
		}
	})

}
