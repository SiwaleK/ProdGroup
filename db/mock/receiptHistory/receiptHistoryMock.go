package mocks

import (
	"errors"
	"time"

	repo "example.com/go-crud-api/repository/receiptHistory"
)

type MockRecieptHistoryRepository struct {
	MockGetRecieptHistoryByID   func(saleOrderID string) ([]repo.ReceiptHistory, error)
	MockGetRecieptHistoryByDate func(startDate, endDate time.Time) ([]repo.ReceiptHistory, error)
}

func (m *MockRecieptHistoryRepository) GetReceiptHistoryByID(saleOrderID string) ([]repo.ReceiptHistory, error) {
	if m.MockGetRecieptHistoryByID != nil {
		return m.MockGetRecieptHistoryByID(saleOrderID)
	}
	return nil, errors.New("MockGetRecieptHistoryByID is not implemented")

}

func (m *MockRecieptHistoryRepository) GetReceiptHistoryByDate(startDate, endDate time.Time) ([]repo.ReceiptHistory, error) {
	if m.MockGetRecieptHistoryByID != nil {
		return m.MockGetRecieptHistoryByDate(startDate, endDate)
	}
	return nil, errors.New("MockGetRecieptHistoryByDate is not implemented")

}

func NewMockRecieptHistoryRepository() *MockRecieptHistoryRepository {
	return &MockRecieptHistoryRepository{
		MockGetRecieptHistoryByID: func(saleOrderID string) ([]repo.ReceiptHistory, error) {
			return []repo.ReceiptHistory{
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
			}, nil
		},
		MockGetRecieptHistoryByDate: func(startDate, endDate time.Time) ([]repo.ReceiptHistory, error) {
			return []repo.ReceiptHistory{
				{
					Saleorderid: "123",
					Totalsale:   100,
					Createdate:  time.Date(2023, time.May, 22, 0, 0, 0, 0, time.UTC),
				},
			}, nil
		},
	}

}
