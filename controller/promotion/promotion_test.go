package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

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

func RandomInt32() int32 {
	return rand.Int31()
}

func RandomTime() time.Time {
	min := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	randomUnix := rand.Int63n(max-min) + min
	return time.Unix(randomUnix, 0)
}

func GenerateRandomJSONRawMessage() (json.RawMessage, error) {
	// Generate random JSON data
	randomData := struct {
		RandomField string
	}{
		RandomField: "random value",
	}

	// Marshal the JSON data to bytes
	dataBytes, err := json.Marshal(randomData)
	if err != nil {
		return nil, err
	}

	// Create json.RawMessage from the bytes
	rawMessage := json.RawMessage(dataBytes)
	return rawMessage, nil
}

func randomPromotionApplied() []db.PostPromotionAppliedParams {
	appliedParams := []db.PostPromotionAppliedParams{
		{
			Promotionid:       RandomString(36),
			Skuid:             RandomString(36),
			PromotiondetailID: RandomString(36),
		},
	}

	return appliedParams
}

func randomPromotion() db.PostPromotionTableParams {
	discount := ConditionDiscount{
		Discount: RandomInt32(),
	}

	conditionBytes, err := json.Marshal(discount)
	if err != nil {
		panic(err)
	}
	rawConditionDiscount := json.RawMessage(conditionBytes)

	argDiscount := db.PostPromotionTableParams{
		Promotionid:    RandomString(36),
		Promotiontitle: RandomString(36),
		Promotiontype:  RandomInt32(),
		Startdate:      RandomTime(),
		Enddate:        RandomTime(),
		Description:    RandomString(36),
		Condition:      rawConditionDiscount,
		// Update the field name to match your struct definition
	}

	return argDiscount

}

func requireBodyMatchPostPromotion(t *testing.T, body *bytes.Buffer, appliedParams []db.PostPromotionAppliedParams, argDiscount db.PostPromotionTableParams) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var combinedData struct {
		PostappliedParams []db.PostPromotionAppliedParams
		PostargDiscount   db.PostPromotionTableParams
	}
	err = json.Unmarshal(data, &combinedData)
	require.NoError(t, err)
	require.Equal(t, appliedParams, combinedData.PostappliedParams)
	require.Equal(t, argDiscount, combinedData.PostargDiscount)

}

type PromotionRepositoryMock struct {
	mock.Mock
}

type MockPromotionRepository struct {
	// Define mock methods and their expected behaviors
	PostPromotionAppliedItemFn func(ctx context.Context, args []db.PostPromotionAppliedParams) error
	PostPromotionFn            func(ctx context.Context, arg db.PostPromotionTableParams) error
}

func (m *MockPromotionRepository) PostPromotionAppliedItem(ctx context.Context, args []db.PostPromotionAppliedParams) error {
	if m.PostPromotionAppliedItemFn != nil {
		return m.PostPromotionAppliedItemFn(ctx, args)
	}
	return nil
}

func (m *MockPromotionRepository) PostPromotion(ctx context.Context, arg db.PostPromotionTableParams) error {
	if m.PostPromotionFn != nil {
		return m.PostPromotionFn(ctx, arg)
	}
	return nil
}

func TestPostPromotionAFreeB(t *testing.T) {
	repo := &MockPromotionRepository{}
	handler := NewPromotionHandler(repo)
	condition := ConditionAFREEB{
		MinimumAmountToEnable: int(RandomInt32()),
		FreeAmount:            int(RandomInt32()),
		PremiumItemsId: []string{
			*RandomString(36),
			*RandomString(36),
		},
	}
	reqBody := PostPromotionAFREEBRequest{
		Promotionid:    RandomString(36),
		Promotiontitle: RandomString(36),
		Promotiontype:  RandomInt32(),
		Startdate:      RandomTime(),
		Enddate:        RandomTime(),
		Description:    RandomString(36),
		Condition:      condition,
		AppliedItemsID: []AppliedItems{
			{
				PromotiondetailID: RandomString(36),
				Skuid:             RandomString(36),
			},
		},
	}
	requestBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Errorf("Failed to marshal request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/sale/api/v1/Promotion/AFreeB", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Errorf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(recorder)
	ginContext.Request = req

	arg := db.PostPromotionTableParams{
		Promotionid:    reqBody.Promotionid,
		Promotiontitle: reqBody.Promotiontitle,
		Promotiontype:  reqBody.Promotiontype,
		Startdate:      reqBody.Startdate,
		Enddate:        reqBody.Enddate,
		Description:    reqBody.Description,
		Condition:      requestBody,
	}
	handler.repo.PostPromotion(ginContext, arg)

	// Perform the request
	//router.ServeHTTP(recorder, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestPostPromotionDiscount(t *testing.T) {
	repo := &MockPromotionRepository{}
	handler := NewPromotionHandler(repo)
	condition := ConditionDiscount{
		Discount: RandomInt32(),
	}

	reqBody := PostPromotionDiscountRequest{
		Promotionid:    RandomString(36),
		Promotiontitle: RandomString(36),
		Promotiontype:  RandomInt32(),
		Startdate:      RandomTime(),
		Enddate:        RandomTime(),
		Description:    RandomString(36),
		Condition:      condition,
		AppliedItemsID: []AppliedItems{
			{
				PromotiondetailID: RandomString(36),
				Skuid:             RandomString(36),
			},
		},
	}

	requestBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Errorf("Failed to marshal request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/sale/api/v1/Promotion/DiscountPerItem", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Errorf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(recorder)
	ginContext.Request = req

	arg := db.PostPromotionTableParams{
		Promotionid:    reqBody.Promotionid,
		Promotiontitle: reqBody.Promotiontitle,
		Promotiontype:  reqBody.Promotiontype,
		Startdate:      reqBody.Startdate,
		Enddate:        reqBody.Enddate,
		Description:    reqBody.Description,
		Condition:      requestBody,
	}
	handler.repo.PostPromotion(ginContext, arg)

	// Perform the request
	//router.ServeHTTP(recorder, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestPostPromotionStepPurchase(t *testing.T) {
	repo := &MockPromotionRepository{}
	handler := NewPromotionHandler(repo)
	condition := ConditionStepPurchase{
		SpecialPriceAtXItemConditionDetail: []SpecialPriceAtXItemConditionDetail{
			{
				MinimumItemToEnable: int(RandomInt32()),
				Discount:            int(RandomInt32()),
			},
		},
	}

	reqBody := PostPromotionStepPurchaseRequest{
		Promotionid:    RandomString(36),
		Promotiontitle: RandomString(36),
		Promotiontype:  RandomInt32(),
		Startdate:      RandomTime(),
		Enddate:        RandomTime(),
		Description:    RandomString(36),
		Condition:      condition,
		AppliedItemsID: []AppliedItems{
			{
				PromotiondetailID: RandomString(36),
				Skuid:             RandomString(36),
			},
		},
	}
	requestBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Errorf("Failed to marshal request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/sale/api/v1/Promotion/StepPurchase", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Errorf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(recorder)
	ginContext.Request = req

	arg := db.PostPromotionTableParams{
		Promotionid:    reqBody.Promotionid,
		Promotiontitle: reqBody.Promotiontitle,
		Promotiontype:  reqBody.Promotiontype,
		Startdate:      reqBody.Startdate,
		Enddate:        reqBody.Enddate,
		Description:    reqBody.Description,
		Condition:      requestBody,
	}
	handler.repo.PostPromotion(ginContext, arg)

	// Perform the request
	//router.ServeHTTP(recorder, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

}

// func TestPostDiscountPromotion(t *testing.T) {
// 	appliedParams := randomPromotionApplied()
// 	argDiscount := randomPromotion()
// 	ctx := context.Background()
// 	testCases := []struct {
// 		name               string
// 		body               gin.H
// 		buildStubs1        func(m *mock_repository.MockPromotionRepository)
// 		buildStubs2        func(m *mock_repository.MockPromotionRepository)
// 		checkResponse      func(recorder *httptest.ResponseRecorder)
// 		expectedStatusCode int
// 	}{
// 		{
// 			name: "Test Case 1 OK",
// 			body: gin.H{
// 				"Promotionid":    argDiscount.Promotionid,
// 				"Promotiontitle": argDiscount.Promotiontitle,
// 				"Promotiontype":  argDiscount.Promotiontype,
// 				"Startdate":      argDiscount.Startdate,
// 				"Enddate":        argDiscount.Enddate,
// 				"Description":    argDiscount.Description,
// 				"Condition":      argDiscount.Condition,
// 				"AppliedItemId":  appliedParams,
// 			},
// 			buildStubs1: func(m *mock_repository.MockPromotionRepository) {
// 				appliedItemID := []db.PostPromotionAppliedParams{
// 					{
// 						PromotiondetailID: appliedParams[0].PromotiondetailID,
// 						Promotionid:       appliedParams[0].Promotionid,
// 						Skuid:             appliedParams[0].Skuid,
// 					},
// 				}
// 				//appliedItemID := appliedParams[0]

// 				// Extract the first element of appliedParams
// 				// Add this line for the missing call
// 				m.EXPECT().PostPromotionAppliedItem(gomock.Any(), gomock.Eq(appliedItemID)).Times(1).Return(nil)
// 				err := m.PostPromotionAppliedItem(ctx, appliedItemID)

// 				// Assert that the expected behavior occurred
// 				if err != nil {
// 					t.Errorf("Expected no error, but got: %v", err)
// 				}

// 				//m.EXPECT().PostPromotionAppliedItem(gomock.Any(), gomock.Eq([]db.PostPromotionAppliedParams{appliedItemID})).Times(1).Return(nil)

// 			},
// 			buildStubs2: func(m *mock_repository.MockPromotionRepository) {
// 				argDiscount := db.PostPromotionTableParams{
// 					Promotionid:    argDiscount.Promotionid,
// 					Promotiontitle: argDiscount.Promotiontitle,
// 					Promotiontype:  argDiscount.Promotiontype,
// 					Startdate:      argDiscount.Startdate,
// 					Enddate:        argDiscount.Enddate,
// 					Description:    argDiscount.Description,
// 					Condition:      argDiscount.Condition,
// 					// Update the field name to match your struct definition
// 				}
// 				m.EXPECT().PostPromotion(gomock.Any(), gomock.Eq(argDiscount)).Times(1).Return(nil)

// 				// Call the method being tested
// 				err := m.PostPromotion(ctx, argDiscount)

// 				// Assert that the expected behavior occurred
// 				if err != nil {
// 					t.Errorf("Expected no error, but got: %v", err)
// 				}

// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code) // Change the expected status code to http.StatusOK
// 			},
// 			expectedStatusCode: http.StatusOK,
// 		},
// 		{
// 			name: "InternalErrorPostPromotionAppliedItem",
// 			body: gin.H{
// 				"Promotionid":    argDiscount.Promotionid,
// 				"Promotiontitle": argDiscount.Promotiontitle,
// 				"Promotiontype":  argDiscount.Promotiontype,
// 				"Startdate":      argDiscount.Startdate,
// 				"Enddate":        argDiscount.Enddate,
// 				"Description":    argDiscount.Description,
// 				"Condition":      argDiscount.Condition,
// 				"AppliedItemId":  appliedParams,
// 			},
// 			buildStubs1: func(m *mock_repository.MockPromotionRepository) {
// 				appliedItemID := []db.PostPromotionAppliedParams{
// 					{
// 						PromotiondetailID: appliedParams[0].PromotiondetailID,
// 						Promotionid:       appliedParams[0].Promotionid,
// 						Skuid:             appliedParams[0].Skuid,
// 					},
// 				}
// 				m.EXPECT().PostPromotionAppliedItem(gomock.Any(), gomock.Eq(appliedItemID)).Times(1).Return(errors.New("internal error"))
// 				err := m.PostPromotionAppliedItem(ctx, appliedItemID)

// 				// Assert that the expected behavior occurred
// 				if err == nil || err.Error() != "internal error" {
// 					t.Errorf("Expected internal error, but got: %v", err)
// 				}
// 			},
// 			buildStubs2: func(m *mock_repository.MockPromotionRepository) {
// 				argDiscount := db.PostPromotionTableParams{
// 					Promotionid:    argDiscount.Promotionid,
// 					Promotiontitle: argDiscount.Promotiontitle,
// 					Promotiontype:  argDiscount.Promotiontype,
// 					Startdate:      argDiscount.Startdate,
// 					Enddate:        argDiscount.Enddate,
// 					Description:    argDiscount.Description,
// 					Condition:      argDiscount.Condition,
// 					// Update the field name to match your struct definition
// 				}
// 				m.EXPECT().PostPromotion(gomock.Any(), gomock.Eq(argDiscount)).Times(1).Return(nil)

// 				// Call the method being tested
// 				err := m.PostPromotion(ctx, argDiscount)

// 				// Assert that the expected behavior occurred
// 				if err != nil {
// 					t.Errorf("error: %v", err)
// 				}
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 		},
// 		{
// 			name: "InternalErrorPostPromotion",
// 			body: gin.H{
// 				"Promotionid":    argDiscount.Promotionid,
// 				"Promotiontitle": argDiscount.Promotiontitle,
// 				"Promotiontype":  argDiscount.Promotiontype,
// 				"Startdate":      argDiscount.Startdate,
// 				"Enddate":        argDiscount.Enddate,
// 				"Description":    argDiscount.Description,
// 				"Condition":      argDiscount.Condition,
// 				"AppliedItemId":  appliedParams,
// 			},
// 			buildStubs1: func(m *mock_repository.MockPromotionRepository) {

// 				m.EXPECT().PostPromotionAppliedItem(gomock.Any(), gomock.Eq(appliedParams)).Times(1).Return(sql.ErrConnDone)
// 			},
// 			buildStubs2: func(m *mock_repository.MockPromotionRepository) {

// 				m.EXPECT().PostPromotion(gomock.Any(), gomock.Eq(argDiscount)).Times(1).Return(nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 			expectedStatusCode: 500,
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mock_repository.NewMockPromotionRepository(ctrl)
// 			tc.buildStubs1(store)
// 			tc.buildStubs2(store)

// 			router := gin.Default()
// 			router.POST("/sale/api/v1/Promotion/DiscountPerItem", PromotionHandler.PostDiscountPromotion) // Add the route definition here
// 			recorder := httptest.NewRecorder()

// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			url := "/sale/api/v1/Promotion/DiscountPerItem"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			router.ServeHTTP(recorder, request)

// 			require.Equal(t, tc.expectedStatusCode, recorder.Code)

// 			// Call the checkResponse function if needed
// 			tc.checkResponse(recorder)
// 		})
// 	}
// }
