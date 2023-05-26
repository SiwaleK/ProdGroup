package repo_test

import (
	"errors"
	"testing"

	repo "example.com/go-crud-api/repository/paymentConfig"

	mocks "example.com/go-crud-api/db/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPosClientConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()

		defer sqlDB.Close()
		posClientID := "778"
		expectedResults := []repo.PosclientConfig{
			{
				Iscash:      true,
				Isqrcode:    true,
				Ispaotang:   true,
				Istongfah:   true,
				Iscoupon:    true,
				Accountname: "dev2",
				Accountcode: "789",
			},
		}
		rows := sqlmock.NewRows([]string{"Iscash", "Isqrcode", "Ispaotang", "Istongfah", "Iscoupon", "Accountname", "Accountcode"})
		for _, posclient := range expectedResults {
			rows.AddRow(posclient.Iscash, posclient.Isqrcode, posclient.Ispaotang, posclient.Istongfah, posclient.Iscoupon, posclient.Accountname, posclient.Accountcode)
		}
		mock.ExpectQuery(".+").WillReturnRows(rows)
		posClient := repo.NewPaymentConfigRepository()
		result, err := posClient.GetPosClientConfig(posClientID)
		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)

	})

	t.Run("Fail", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		posClientID := "pp"

		expectedError := errors.New("some database error ")
		mock.ExpectQuery(".+").WillReturnError(expectedError)
		posClient := repo.NewPaymentConfigRepository()
		_, err := posClient.GetPosClientConfig(posClientID)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)

	})

}

func TestGetPaymentMethod(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		expectedResults := []repo.PaymentMethod{
			{Paymentmethodid: 1, Paymentname: "cash"},
			{Paymentmethodid: 2, Paymentname: "qrcode"},
			{Paymentmethodid: 3, Paymentname: "paotung"},
			{Paymentmethodid: 4, Paymentname: "tongfah"},
			{Paymentmethodid: 5, Paymentname: "Coupon"},
		}

		rows := sqlmock.NewRows([]string{"paymentmethodid", "paymentname"})
		for _, r := range expectedResults {
			rows.AddRow(r.Paymentmethodid, r.Paymentname)

		}
		mock.ExpectQuery(".+").WillReturnRows(rows)
		promotionRepo := repo.NewPaymentConfigRepository()
		result, err := promotionRepo.GetPaymentMethod()
		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)

	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		expectedError := errors.New("some database error ")
		mock.ExpectQuery(".+").WillReturnError(expectedError)
		posClient := repo.NewPaymentConfigRepository()
		_, err := posClient.GetPaymentMethod()
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)

	})
}
