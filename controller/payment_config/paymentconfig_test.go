package controller

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/SiwaleK/ProdGroup/db/mock"
	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
)

func RandomInt0() *int16 {

	random := int16(rand.Intn(2))
	return &random
}

func RandomInt32() int32 {
	return rand.Int31()
}

func RandomString(n int) *string {
	str := RandomStringWithCharset(n, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return &str
}

func RandomStringWithCharset(n int, charset string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	str := string(b)
	return str
}

// func RandomPosClient() []db.GetPosClientMethodRow {

// 	posClientID := RandomString(36)
// 	PaymentConfig := []db.GetPosClientMethodRow{
// 		{
// 			IsCash:      RandomInt0(),
// 			IsPaotang:   RandomInt0(),
// 			IsQrcode:    RandomInt0(),
// 			IsTongfah:   RandomInt0(),
// 			IsCoupon:    RandomInt0(),
// 			AccountName: RandomString(36),
// 			AccountCode: RandomString(36),
// 		},
// 	}

// 	return PaymentConfig
// }

// func TestGetPosClientMethod(t *testing.T) {
// 	// Prepare the test cases
// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(m *mockdb.MockPaymentConfigRepository)
// 		checkResponse func(recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"posClientID": "examplePosClientID",
// 			},
// 			buildStubs: func(m *mockdb.MockPaymentConfigRepository) {
// 				// Set up the expected calls on the mock repository
// 				m.EXPECT().GetPosClientMethod(gomock.Any(), gomock.Eq("examplePosClientID"))
// 					.Returns([]db.GetPosClientMethodRow{
// 						{
// 							IsCash:      RandomInt0(),
// 							IsPaotang:   RandomInt0(),
// 							IsQrcode:    RandomInt0(),
// 							IsTongfah:   RandomInt0(),
// 							IsCoupon:    RandomInt0(),
// 							AccountName: RandomString(36),
// 							AccountCode: RandomString(36),
// 						},
// 					}, nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				// Assert the response status code, headers, and body
// 				assert.Equal(t, http.StatusOK, recorder.Code)
// 				assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
// 				// TODO: Assert the response body content
// 			},
// 		},
// 		// Add more test cases as needed
// 	}

// 	// Iterate over the test cases
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create a new controller for the test case
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			// Create a mock repository and build the stubs
// 			mockRepo := mockdb.NewMockPaymentConfigRepository(ctrl)
// 			tc.buildStubs(mockRepo)

// 			// Create a new instance of the handler with the mock repository
// 			handler := &PaymentconfigHandlerImpl{
// 				repo: mockRepo,
// 			}

// 			// Create a new Gin router and register the handler function
// 			router := gin.Default()
// 			router.GET("/sale/api/v1/PaymentConfigs", handler.GetPosClientMethod)

// 			// Create a new HTTP request for the endpoint
// 			req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfig", nil)
// 			// TODO: Set the request body and headers if needed

// 			// Create a response recorder to capture the response
// 			recorder := httptest.NewRecorder()

// 			// Serve the HTTP request using the router and recorder
// 			router.ServeHTTP(recorder, req)

// 			// Check the response against the expected result
// 			tc.checkResponse(recorder)
// 		})
// 	}
// }

// func TestGetPosClientMethod(t *testing.T) {
// 	// Create a new instance of the handler with a mock repository
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	var req GetPosClientMethodReq
// 	posClientID := "77"

// 	// Assign the address of posClientID to the req struct field
// 	req.PosClientID = posClientID

// 	mockRepo := mockdb.NewMockPaymentConfigRepository(ctrl)
// 	mockRepo.EXPECT().GetPaymentMethod(gomock.Any()).
// 		Return([]db.PaymentMethod{
// 			{
// 				Paymentmethodid: RandomInt32(),
// 				Paymentname:     RandomString(36),
// 			},
// 		}, nil).Times(1)
// 	// Set up the expectation for the GetPosClientMethod method
// 	mockRepo.EXPECT().GetPosClientMethod(gomock.Any(), gomock.Eq(&req.PosClientID)).
// 		Return([]db.GetPosClientMethodRow{
// 			{
// 				IsCash:      RandomInt0(),
// 				IsPaotang:   RandomInt0(),
// 				IsQrcode:    RandomInt0(),
// 				IsTongfah:   RandomInt0(),
// 				IsCoupon:    RandomInt0(),
// 				AccountName: RandomString(36),
// 				AccountCode: RandomString(36),
// 			},
// 		}, nil).Times(2)

// 	handler := &PaymentconfigHandlerImpl{
// 		repo: mockRepo,
// 	}

// 	// Test case 1: Successful request
// 	t.Run("SuccessfulRequest", func(t *testing.T) {
// 		// Create a new Gin router and register the handler function
// 		router := gin.Default()
// 		router.GET("/sale/api/v1/PaymentConfigs", handler.GetPosClientMethod)

// 		// Create a new HTTP request with JSON data
// 		jsonData := `{"posClientID": "examplePosClientID"}`
// 		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfigs", strings.NewReader(jsonData))

// 		reqBody := `{"posClientID": "77"}`
// 		req.Body = ioutil.NopCloser(strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")

// 		// Create a response recorder to capture the response
// 		recorder := httptest.NewRecorder()

// 		// Serve the HTTP request using the router and recorder
// 		router.ServeHTTP(recorder, req)

// 		// Assert the response status code, headers, and body
// 		assert.Equal(t, http.StatusOK, recorder.Code)
// 		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
// 		// TODO: Assert the response body content
// 	})

// 	// Test case 2: Invalid JSON request
// 	t.Run("InvalidJSONRequest", func(t *testing.T) {
// 		// Create a new Gin router and register the handler function
// 		router := gin.Default()
// 		router.GET("/sale/api/v1/PaymentConfigs", handler.GetPosClientMethod)

// 		// Create a new HTTP request with invalid JSON data
// 		req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfigs", strings.NewReader("invalidJSON"))
// 		reqBody := `{"posClientID": "77"}`
// 		req.Body = ioutil.NopCloser(strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")

// 		// Create a response recorder to capture the response
// 		recorder := httptest.NewRecorder()

// 		// Serve the HTTP request using the router and recorder
// 		router.ServeHTTP(recorder, req)

// 		// Assert the response status code, headers, and body
// 		assert.Equal(t, http.StatusBadRequest, recorder.Code)
// 		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
// 		// TODO: Assert the response body content
// 	})

// 	// TODO: Add more test cases to cover different scenarios

//		// TODO: Test error cases and handle the assertions accordingly
//	}
func TestPaymentConfig(t *testing.T) {
	var req GetPosClientMethodReq
	posClientID := "77"

	// Assign the address of posClientID to the req struct field
	req.PosClientID = posClientID
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs1   func(m *mockdb.MockPaymentConfigRepository)
		buildStubs2   func(m *mockdb.MockPaymentConfigRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"posClientID": req.PosClientID,
			},
			buildStubs1: func(m *mockdb.MockPaymentConfigRepository) {
				m.EXPECT().GetPaymentMethod(gomock.Any()).
					Return([]db.PaymentMethod{
						{
							Paymentmethodid: RandomInt32(),
							Paymentname:     RandomString(36),
						},
					}, nil).Times(1)
			},
			buildStubs2: func(m *mockdb.MockPaymentConfigRepository) {
				m.EXPECT().GetPosClientMethod(gomock.Any(), gomock.Eq(&req.PosClientID)).
					Return([]db.GetPosClientMethodRow{
						{
							IsCash:      RandomInt0(),
							IsPaotang:   RandomInt0(),
							IsQrcode:    RandomInt0(),
							IsTongfah:   RandomInt0(),
							IsCoupon:    RandomInt0(),
							AccountName: RandomString(36),
							AccountCode: RandomString(36),
						},
					}, nil).Times(1)
			}, checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockPaymentConfigRepository(ctrl)
			tc.buildStubs1(store)
			tc.buildStubs2(store)
			handler := &PaymentconfigHandlerImpl{
				repo: store,
			}

			router := gin.Default()
			router.GET("/sale/api/v1/PaymentConfigs", handler.GetPosClientMethod)
			recorder := httptest.NewRecorder()
			tc.checkResponse(t, recorder)

			bodyJSON, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("GET", "/sale/api/v1/PaymentConfigs", bytes.NewBuffer(bodyJSON))
			req.Header.Set("Content-Type", "application/json")

			// Serve the HTTP request using the router and recorder
			router.ServeHTTP(recorder, req)

			// Check the response status code, headers, and body
			// Example assertions:
			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
		})
	}
}
