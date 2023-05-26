package mocks

import (
	"errors"

	"example.com/go-crud-api/db/db"
)

type MockPromotionRepository struct {
	MockCreatePromotion            func(arg []db.Promotion) error
	MockCreatePromotionAppliedItem func(arg []db.PromotionAppliedItemsID) error
}

func (m *MockPromotionRepository) CreatePromotion(arg []db.Promotion) error {
	if m.MockCreatePromotion != nil {
		return m.MockCreatePromotion(arg)
	}
	return errors.New("MockCreatePromotion is not implemented")
}

func (m *MockPromotionRepository) CreatePromotionAppliedItem(arg []db.PromotionAppliedItemsID) error {
	if m.MockCreatePromotion != nil {
		return m.MockCreatePromotionAppliedItem(arg)
	}
	return errors.New("MockCreatePromotion is not implemented")
}


func NewMockPromotionRepository() *MockPromotionRepository{
	return &MockPromotionRepository{
		MockCreatePromotion:   func()(arg []db.PromotionAppliedItemsID)error{
			return nil
		}
}