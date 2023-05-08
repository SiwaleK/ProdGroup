package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

type PaymentConfigRepository interface {
	GetPaymentConfig(ctx context.Context) ([]db.GetPaymentConfigRow, error)
}

type DBPaymentConfigRepository struct {
	db *db.Queries
}

func NewPaymentConfigRepository(db *db.Queries) PaymentConfigRepository {
	return &DBPaymentConfigRepository{
		db: db,
	}
}

func (r *DBPaymentConfigRepository) GetPaymentConfig(ctx context.Context) ([]db.GetPaymentConfigRow, error) {
	dbPaymentconfigs, err := r.db.GetPaymentConfig(ctx)
	if err != nil {
		return nil, err
	}
	paymentconfigs := make([]db.GetPaymentConfigRow, len(dbPaymentconfigs))
	for i, dbPaymentconfigs := range dbPaymentconfigs {
		paymentconfigs[i] = db.GetPaymentConfigRow{
			IsCash:      dbPaymentconfigs.IsCash,
			IsQrcode:    dbPaymentconfigs.IsQrcode,
			IsPaotang:   dbPaymentconfigs.IsPaotang,
			IsTongfah:   dbPaymentconfigs.IsTongfah,
			IsCoupon:    dbPaymentconfigs.IsCoupon,
			AccountName: dbPaymentconfigs.AccountName,
			AccountCode: dbPaymentconfigs.AccountCode,
		}
	}
	return paymentconfigs, nil
}
