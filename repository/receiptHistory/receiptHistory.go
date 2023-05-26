package receipthistory

import (
	"errors"
	"time"

	"example.com/go-crud-api/db/database"
)

type ReceiptHistoryRepository interface {
	GetReceiptHistoryByID(saleOrderID string) ([]ReceiptHistory, error)
	GetReceiptHistoryByDate(startDate, endDate time.Time) ([]ReceiptHistory, error)
}

type receiptHistoryRepository struct{}

func NewReceiptHistoryRepository() ReceiptHistoryRepository {
	return &receiptHistoryRepository{}
}

type ReceiptHistory struct {
	Saleorderid string    `gorm:"column:saleorderid" json:"saleOrderID"`
	Totalsale   float64   `gorm:"column:totalsale" json:"TotalSale"`
	Createdate  time.Time `gorm:"column:createdate;not null" json:"CreateDate"`
}

func (r *receiptHistoryRepository) GetReceiptHistoryByID(saleOrderID string) ([]ReceiptHistory, error) {
	var receiptHistories []ReceiptHistory
	if saleOrderID == "" {
		return nil, errors.New("saleOrderID parameters are required")
	}

	if err := database.DB.Table("saleorder").
		Select("saleorder.saleorderid, saleorder.totalsale, saleorder.createdate").
		Where("saleorder.saleorderid = ?", saleOrderID).
		Find(&receiptHistories).
		Error; err != nil {
		return nil, err
	}

	return receiptHistories, nil
}

func (r *receiptHistoryRepository) GetReceiptHistoryByDate(startDate, endDate time.Time) ([]ReceiptHistory, error) {
	var receiptHistories []ReceiptHistory

	if startDate.IsZero() || endDate.IsZero() {
		return nil, errors.New("startDate and endDate parameters are required")
	}
	if err := database.DB.Table("saleorder").
		Select("saleorder.saleorderid, saleorder.totalsale, saleorder.createdate").
		Where("saleorder.createdate >= ? AND saleorder.createdate <= ?", startDate, endDate).
		Find(&receiptHistories).
		Error; err != nil {
		return nil, err
	}

	return receiptHistories, nil
}
