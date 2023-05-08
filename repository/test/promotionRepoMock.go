package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/stretchr/testify/mock"
)

type PromotionRepositoryMock struct {
	mock.Mock
}

func (r *PromotionRepositoryMock) GetPromotionByID(ctx context.Context, promotionID string) (*db.Promotion, error) {
	args := r.Called(ctx)

	if promotion, ok := args.Get(0).(*db.Promotion); ok {
		return promotion, args.Error(1)
	}

	return nil, args.Error(1)
}
