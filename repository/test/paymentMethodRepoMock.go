package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/stretchr/testify/mock"
)

type PaymentMethodRepositoryMock struct {
	mock.Mock
}

func (r *PaymentMethodRepositoryMock) GetPaymentMethod(ctx context.Context) (*db.PaymentMethod, error) {
	args := r.Called(ctx)

	if payment_method, ok := args.Get(0).(*db.PaymentMethod); ok {
		return payment_method, args.Error(1)
	}
	return nil, args.Error(1)

}
