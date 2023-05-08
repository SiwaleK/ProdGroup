package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/stretchr/testify/mock"
)

func TestGetPromotionByID(t *testing.T) {
	// Mock data

	repo := &PromotionRepositoryMock{}
	mockPromotion := &db.Promotion{
		Promotionid:   sql.NullString{String: "123", Valid: true},
		Promotiontype: 1,
		Startdate:     time.Date(2023, time.May, 8, 12, 0, 0, 0, time.UTC),
		Enddate:       time.Date(2023, time.May, 20, 12, 0, 0, 0, time.UTC),
		Description:   sql.NullString{String: "Sample promotion", Valid: true},
		Conditions:    json.RawMessage(`{"key": "value"}`),
	}
	repo.On("GetPromotionByID", mock.Anything).Return(mockPromotion, nil)

	// Call the GetPromotionByID method
	ctx := context.Background()
	promotion, err := repo.GetPromotionByID(ctx, "123")

	expectedPromotion := &db.Promotion{
		Promotionid:   sql.NullString{String: "123", Valid: true},
		Promotiontype: 1,
		Startdate:     time.Date(2023, time.May, 8, 12, 0, 0, 0, time.UTC),
		Enddate:       time.Date(2023, time.May, 20, 12, 0, 0, 0, time.UTC),
		Description:   sql.NullString{String: "Sample promotion", Valid: true},
		Conditions:    json.RawMessage(`{"key": "value"}`),
	}

	// Assert the expected values
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	// if !reflect.DeepEqual(promotion, expectedPromotion) {
	// 	t.Errorf("Expected promotion %+v, but got %+v", mockPromotion, promotion)
	// }

	if sql.NullString(promotion.Promotionid) != sql.NullString(expectedPromotion.Promotionid) {
		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Promotionid, promotion.Promotionid)
	}
	if time.Time(promotion.Startdate) != time.Time(expectedPromotion.Startdate) {
		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Startdate, promotion.Startdate)
	}
	if time.Time(promotion.Enddate) != time.Time(expectedPromotion.Enddate) {
		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Enddate, promotion.Enddate)
	}

	if string(promotion.Conditions) != string(expectedPromotion.Conditions) {
		t.Errorf("Expected promotion %+v, but got %+v", expectedPromotion.Conditions, promotion.Conditions)
	}

}
