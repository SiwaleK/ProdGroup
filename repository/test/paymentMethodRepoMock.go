package repository

import (
	"context"
	"database/sql"
	"fmt"

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

func PaymentMethodMock() {
	repo := &PaymentMethodRepositoryMock{}

	// Set up the mock response
	mockProdgroup := &db.PaymentMethod{
		Paymentmethodid: 1,
		Paymentname:     sql.NullString{String: "QRcode", Valid: true},
	}
	repo.On("GetPaymentMethod", mock.Anything).Return(mockProdgroup, nil)

	// Call the GetProdgroup method
	ctx := context.Background()
	paymentMethod, err := repo.GetPaymentMethod(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PaymentMethod:", paymentMethod.Paymentmethodid, paymentMethod.Paymentname.String)
}
