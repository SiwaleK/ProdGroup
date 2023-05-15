package repository

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"testing"
	"time"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/go-playground/assert"
)

// func TestGetPromotionByID(t *testing.T) {
// 	// Mock data

// 	repo := &PromotionRepositoryMock{}
// 	mockPromotion := &db.Promotion{
// 		Promotionid:   sql.NullString{String: "123", Valid: true},
// 		Promotiontype: 1,
// 		Startdate:     time.Date(2023, time.May, 8, 12, 0, 0, 0, time.UTC),
// 		Enddate:       time.Date(2023, time.May, 20, 12, 0, 0, 0, time.UTC),
// 		Description:   sql.NullString{String: "Sample promotion", Valid: true},
// 		Conditions:    json.RawMessage(`{"key": "value"}`),
// 	}
// 	repo.On("GetPromotionByID", mock.Anything).Return(mockPromotion, nil)

// 	// Call the GetPromotionByID method
// 	ctx := context.Background()
// 	promotion, err := repo.GetPromotionByID(ctx, "123")

// 	expectedPromotion := &db.Promotion{
// 		Promotionid:   sql.NullString{String: "123", Valid: true},
// 		Promotiontype: 1,
// 		Startdate:     time.Date(2023, time.May, 8, 12, 0, 0, 0, time.UTC),
// 		Enddate:       time.Date(2023, time.May, 20, 12, 0, 0, 0, time.UTC),
// 		Description:   sql.NullString{String: "Sample promotion", Valid: true},
// 		Conditions:    json.RawMessage(`{"key": "value"}`),
// 	}

// 	// Assert the expected values
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}
// 	// if !reflect.DeepEqual(promotion, expectedPromotion) {
// 	// 	t.Errorf("Expected promotion %+v, but got %+v", mockPromotion, promotion)
// 	// }

// 	if sql.NullString(promotion.Promotionid) != sql.NullString(expectedPromotion.Promotionid) {
// 		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Promotionid, promotion.Promotionid)
// 	}
// 	if time.Time(promotion.Startdate) != time.Time(expectedPromotion.Startdate) {
// 		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Startdate, promotion.Startdate)
// 	}
// 	if time.Time(promotion.Enddate) != time.Time(expectedPromotion.Enddate) {
// 		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Enddate, promotion.Enddate)
// 	}

// 	if string(promotion.Conditions) != string(expectedPromotion.Conditions) {
// 		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Conditions, promotion.Conditions)
// 	}

// }
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

func TestDBPromotionRepository_PostPromotionTable(t *testing.T) {

	randomJSON, err := GenerateRandomJSONRawMessage()

	// Mock dependencies
	mockDB := &PromotionRepositoryMock{} // Replace with your mock implementation
	//repo := repository.NewPromotionRepository(mockDB)

	// Define test case input
	ctx := context.Background()
	arg := db.PostPromotionTableParams{
		Promotionid:    RandomString(36),
		Promotiontitle: RandomString(36),
		Promotiontype:  RandomInt32(),
		Startdate:      RandomTime(),
		Enddate:        RandomTime(),
		Description:    RandomString(36),
		Condition:      randomJSON,
	}

	// Define expected result
	expectedErr := errors.New("expected error")

	// Set up mock expectations and return values
	mockDB.On("PostPromotionTable", ctx, arg).Return(expectedErr)

	// Call the method being tested
	err = mockDB.PostPromotionTable(ctx, arg)

	// Assert the result
	assert.Equal(t, expectedErr, err)

}

func TestDBPromotionRepository_PostPromotionAppliedItem(t *testing.T) {

	// Mock dependencies
	mockDB := &PromotionRepositoryMock{} // Replace with your mock implementation
	//repo := repository.NewPromotionRepository(mockDB)

	// Define test case input
	ctx := context.Background()
	appliedParams := []db.PostPromotionAppliedParams{
		{
			Promotionid:       RandomString(36),
			Skuid:             RandomString(36),
			PromotiondetailID: RandomString(36),
		},
	}

	// Define expected result
	expectedErr := errors.New("expected error")

	// Set up mock expectations and return values
	mockDB.On("PostPromotionAppliedItem", ctx, appliedParams).Return(expectedErr)

	// Call the method being tested
	err := mockDB.PostPromotionAppliedItem(ctx, appliedParams)

	// Assert the result
	assert.Equal(t, expectedErr, err)

}
