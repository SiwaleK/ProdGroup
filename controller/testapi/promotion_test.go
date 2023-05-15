package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/stretchr/testify/require"

	"net/http"
	"net/http/httptest"

	mock_repository "github.com/SiwaleK/ProdGroup/repository/Mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
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

type ConditionDiscount struct {
	Discount int32
}
type ConditionAFREEB struct {
	MinimumAmountToEnable int
	FreeAmount            int
	PremiumItemsId        []string
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

	PremiumItemsId := make([]string, 0) // Initialize an empty slice

	for {
		PremiumItemsId = append(PremiumItemsId, *RandomString(36))
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

func TestPostDiscountPromotion(t *testing.T) {
	appliedParams := randomPromotionApplied()
	argDiscount := randomPromotion()

	testCases := []struct {
		name        string
		body        gin.H
		buildStubs1 func(*mock_repository.MockPromotionRepository)
		buildStubs2 func(m *mock_repository.MockPromotionRepository)
	}{
		{
			name: "Test Case 1",
			body: gin.H{
				// ...
			},
			buildStubs1: func(m *mock_repository.MockPromotionRepository) {
				appliedItemID := []db.PostPromotionAppliedParams{
					{
						PromotiondetailID: appliedParams[0].PromotiondetailID,
						Promotionid:       appliedParams[0].Promotionid,
						Skuid:             appliedParams[0].Skuid,
					},
				}
				m.EXPECT().PostPromotionAppliedItem(gomock.Any(), gomock.Eq(appliedItemID)).Times(1).Return(nil)
			},
			buildStubs2: func(m *mock_repository.MockPromotionRepository) {
				argDiscount := db.PostPromotionTableParams{
					Promotionid:    argDiscount.Promotionid,
					Promotiontitle: argDiscount.Promotiontitle,
					Promotiontype:  argDiscount.Promotiontype,
					Startdate:      argDiscount.Startdate,
					Enddate:        argDiscount.Enddate,
					Description:    argDiscount.Description,
					Condition:      argDiscount.Condition,
					// Update the field name to match your struct definition
				}
				m.EXPECT().PostPromotion(gomock.Any(), gomock.Eq(argDiscount)).Times(1).Return(nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPostPromotion(t, recorder.Body, appliedParams, argDiscount)
			},
		},
		{
			name: "InternalErrorPostPromotionAppliedItem",
			body: gin.H{
				"Promotionid": argDiscount.Promotionid,
			},
			buildStubs1: func(m *mock_repository.MockPromotionRepository) {

				m.EXPECT().PostPromotionAppliedItem(gomock.Any(), gomock.Eq(appliedParams)).Times(1).Return(db.PostPromotionAppliedParams{}, sql.ErrConnDone)
			},
			buildStubs2: func(m *mock_repository.MockPromotionRepository) {

				m.EXPECT().PostPromotion(gomock.Any(), gomock.Eq(argDiscount)).Times(1).Return(db.PostPromotionTableParams{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalErrorPostPromotion",
			body: gin.H{
				"Promotionid": argDiscount.Promotionid,
			},
			buildStubs1: func(m *mock_repository.MockPromotionRepository) {

				m.EXPECT().PostPromotionAppliedItem(gomock.Any(), gomock.Eq(appliedParams)).Times(1).Return(db.PostPromotionAppliedParams{}, sql.ErrConnDone)
			},
			buildStubs2: func(m *mock_repository.MockPromotionRepository) {

				m.EXPECT().PostPromotion(gomock.Any(), gomock.Eq(argDiscount)).Times(1).Return(db.PostPromotionTableParams{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_repository.NewMockStore(ctrl)
			tc.buildStubs1(store)
			tc.buildStubs2(store)

			router := gin.New()
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/sale/v1/SaleOrder"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

	// Rest of the test code
}
