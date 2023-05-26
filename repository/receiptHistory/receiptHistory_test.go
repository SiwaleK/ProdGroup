package receipthistory_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	mocks "example.com/go-crud-api/db/mock"
	repo "example.com/go-crud-api/repository/receiptHistory"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetRecieptHistoryRepositoryByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		saleOrderID := "123"

		expectedResults := []repo.ReceiptHistory{
			{
				Saleorderid: "123",
				Totalsale:   100,
				Createdate:  time.Date(2023, 5, 11, 7, 0, 0, 0, time.FixedZone("GMT+7", 7*60*60)),
			},

			{
				Saleorderid: "123",
				Totalsale:   100,
				Createdate:  time.Date(2023, 5, 11, 7, 0, 0, 0, time.FixedZone("GMT+7", 7*60*60)),
			},
		}
		rows := sqlmock.NewRows([]string{"saleorderid", "totalsale", "createdate"})
		for _, receiptHistory := range expectedResults {
			rows.AddRow(receiptHistory.Saleorderid, receiptHistory.Totalsale, receiptHistory.Createdate)
		}
		mock.ExpectQuery(".+").WillReturnRows(rows)
		receipthistory := repo.NewReceiptHistoryRepository()
		result, err := receipthistory.GetReceiptHistoryByID(saleOrderID)
		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)

	})
	t.Run("invalid saleOrderid", func(t *testing.T) {
		db, _ := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		saleOrderID := ""

		receipthistory := repo.NewReceiptHistoryRepository()
		_, err := receipthistory.GetReceiptHistoryByID(saleOrderID)

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("saleOrderID parameters are required"), err)
	})
	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
		saleOrderID := "123"
		expectedError := errors.New("some database err")
		mock.ExpectQuery(".+").WillReturnError(expectedError)
		receipthistory := repo.NewReceiptHistoryRepository()
		_, err := receipthistory.GetReceiptHistoryByID(saleOrderID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

}

func TestGetRecieptHistoryRepositoryByDate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		startDate := "2023-05-11"
		endDate := "2023-06-11"
		startDateTime, _ := time.Parse("2006-01-02", startDate)
		endDateTime, _ := time.Parse("2006-01-02", endDate)

		expectedResults := []repo.ReceiptHistory{
			{
				Saleorderid: "123",
				Totalsale:   100,
				Createdate:  time.Date(2023, time.May, 22, 0, 0, 0, 0, time.UTC),
			},

			{
				Saleorderid: "123",
				Totalsale:   100,
				Createdate:  time.Date(2023, time.May, 22, 0, 0, 0, 0, time.UTC),
			},
		}
		rows := sqlmock.NewRows([]string{"saleorderid", "totalsale", "createdate"})
		for _, receiptHistory := range expectedResults {
			rows.AddRow(receiptHistory.Saleorderid, receiptHistory.Totalsale, receiptHistory.Createdate)
		}
		mock.ExpectQuery(".+").WillReturnRows(rows)
		receipthistory := repo.NewReceiptHistoryRepository()
		result, err := receipthistory.GetReceiptHistoryByDate(startDateTime, endDateTime)
		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)

	})
	t.Run("invalid start date and end date ", func(t *testing.T) {
		db, _ := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		startDate := time.Time{}
		endDate := time.Time{}

		receipthistory := repo.NewReceiptHistoryRepository()
		_, err := receipthistory.GetReceiptHistoryByDate(startDate, endDate)

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("startDate and endDate parameters are required"), err)
	})
	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		startDate := "2023-05-11"
		endDate := "2023-06-11"
		startDateTime, _ := time.Parse("2006-01-02", startDate)
		endDateTime, _ := time.Parse("2006-01-02", endDate)
		expectedError := errors.New("some database err")
		mock.ExpectQuery(".+").WillReturnError(expectedError)
		receipthistory := repo.NewReceiptHistoryRepository()

		_, err := receipthistory.GetReceiptHistoryByDate(startDateTime, endDateTime)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

}
