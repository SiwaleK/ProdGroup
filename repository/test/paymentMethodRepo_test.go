package repository

import (
	"context"
	"errors"
	"testing"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestGetPaymentMethod(t *testing.T) {
// 	repo := &PaymentMethodRepositoryMock{}

// 	// Set up the mock response
// 	mockPaymentMethod := &db.PaymentMethod{
// 		Paymentmethodid: 1,
// 		Paymentname:     sql.NullString{String: "QRcode", Valid: true},
// 	}
// 	repo.On("GetPaymentMethod", mock.Anything).Return(mockPaymentMethod, nil)

// 	// Call the GetPaymentMethod method
// 	ctx := context.Background()
// 	paymentMethod, err := repo.GetPaymentMethod(ctx)

// 	// Assert the expected values
// 	expectedPaymentMethod := &db.PaymentMethod{
// 		Paymentmethodid: 1,
// 		Paymentname:     sql.NullString{String: "QRcode", Valid: true},
// 	}
// 	if !reflect.DeepEqual(paymentMethod, expectedPaymentMethod) {
// 		t.Errorf("Expected PaymentMethod %+v, but got %+v", expectedPaymentMethod, paymentMethod)
// 	}
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}
// }

func TestGetPaymentMethod(t *testing.T) {
	repo := &PaymentMethodRepositoryMock{}
	expectedPaymentMethod := &db.PaymentMethod{
		Paymentmethodid: RandomInt32(),
		Paymentname:     RandomString(36),
	}
	expectedError := errors.New("some error")

	// Set up the mock behavior for the case where PaymentMethod is found
	repo.On("GetPaymentMethod", mock.Anything).Return(expectedPaymentMethod, expectedError)

	// Call the function being tested
	paymentMethod, err := repo.GetPaymentMethod(context.Background())

	// Assert that the returned PaymentMethod matches the expectedPaymentMethod
	assert.Equal(t, expectedPaymentMethod, paymentMethod)

	// Assert that the returned error matches the expected error
	assert.Equal(t, expectedError, err)

	// Assert that the mock repository's Call method was called with the correct arguments
	repo.AssertCalled(t, "GetPaymentMethod", context.Background())

	// Set up the mock behavior for the case where PaymentMethod is not found
	repo.On("GetPaymentMethod", mock.Anything).Return((*db.PaymentMethod)(nil), expectedError)

	// Call the function being tested again
	paymentMethod, err = repo.GetPaymentMethod(context.Background())

	// Assert that the returned PaymentMethod is nil
	assert.Nil(t, paymentMethod)

	// Assert that the returned error matches the expected error
	assert.Equal(t, expectedError, err)

	// Assert that the mock repository's Call method was called with the correct arguments
	repo.AssertCalled(t, "GetPaymentMethod", context.Background())
}
