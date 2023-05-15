package repository

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
