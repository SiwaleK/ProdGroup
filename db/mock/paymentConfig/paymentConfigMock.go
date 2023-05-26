package mocks

import (
	"errors"

	repo "example.com/go-crud-api/repository/paymentConfig"
)

type MockPaymentConfigRepository struct {
	MockGetPosClientConfig func(posClientID string) ([]repo.PosclientConfig, error)
	MockGetPaymentMethod   func() ([]repo.PaymentMethod, error)
}

func (m *MockPaymentConfigRepository) GetPosClientConfig(posClientID string) ([]repo.PosclientConfig, error) {
	if m.MockGetPosClientConfig != nil {
		return m.MockGetPosClientConfig(posClientID)
	}
	return nil, errors.New("MockGetPosClientConfig is not implemented")
}

func (m *MockPaymentConfigRepository) GetPaymentMethod() ([]repo.PaymentMethod, error) {
	if m.MockGetPaymentMethod != nil {
		return m.MockGetPaymentMethod()
	}
	return nil, errors.New("MockGetPaymentMethod is not implemented")
}

func NewMockPaymentConfigRepository() *MockPaymentConfigRepository {
	return &MockPaymentConfigRepository{
		MockGetPosClientConfig: func(posClientID string) ([]repo.PosclientConfig, error) {
			return []repo.PosclientConfig{
				{
					Iscash:      true,
					Isqrcode:    true,
					Ispaotang:   true,
					Istongfah:   true,
					Iscoupon:    true,
					Accountname: "dev2",
					Accountcode: "789",
				},
			}, nil
		},
		MockGetPaymentMethod: func() ([]repo.PaymentMethod, error) {
			return []repo.PaymentMethod{
				{
					Paymentmethodid: 1,
					Paymentname:     "cash",
				},
				{
					Paymentmethodid: 2,
					Paymentname:     "qrcode",
				},
				{
					Paymentmethodid: 3,
					Paymentname:     "paotung",
				},
				{
					Paymentmethodid: 4,
					Paymentname:     "tongfah",
				},
				{
					Paymentmethodid: 5,
					Paymentname:     "Coupon",
				},
			}, nil
		},
	}
}
