package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

type PaymentConfigRepository interface {
	GetPaymentConfig(ctx context.Context) ([]db.GetPaymentConfigRow, error)
	GetPaymentMethod(ctx context.Context) ([]db.PaymentMethod, error)
	GetPosClientMethod(ctx context.Context, posClientID *string) ([]db.GetPosClientMethodRow, error)
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
	dbPaymentConfig, err := r.db.GetPaymentConfig(ctx)
	if err != nil {
		return nil, err
	}
	paymentConfig := make([]db.GetPaymentConfigRow, len(dbPaymentConfig))
	for i, dbPaymentConfig := range dbPaymentConfig {
		paymentConfig[i] = db.GetPaymentConfigRow{
			AccountName: dbPaymentConfig.AccountName,
			AccountCode: dbPaymentConfig.AccountCode,
		}
	}

	return paymentConfig, nil
}

func (r *DBPaymentConfigRepository) GetPaymentMethod(ctx context.Context) ([]db.PaymentMethod, error) {
	dbPaymentMethod, err := r.db.GetPaymentMethod(ctx)
	if err != nil {
		return nil, err
	}
	paymentMethod := make([]db.PaymentMethod, len(dbPaymentMethod))
	for i, dbPaymentMethod := range dbPaymentMethod {
		paymentMethod[i] = db.PaymentMethod{
			Paymentmethodid: dbPaymentMethod.Paymentmethodid,
			Paymentname:     dbPaymentMethod.Paymentname,
		}
	}

	return paymentMethod, nil
}

func (r *DBPaymentConfigRepository) GetPosClientMethod(ctx context.Context, posClientID *string) ([]db.GetPosClientMethodRow, error) {
	dbPosclient, err := r.db.GetPosClientMethod(ctx, posClientID)
	if err != nil {
		return nil, err
	}
	posclient := make([]db.GetPosClientMethodRow, len(dbPosclient))
	for i, dbPosclient := range dbPosclient {
		posclient[i] = db.GetPosClientMethodRow{
			IsCash:      dbPosclient.IsCash,
			IsPaotang:   dbPosclient.IsPaotang,
			IsQrcode:    dbPosclient.IsQrcode,
			IsTongfah:   dbPosclient.IsTongfah,
			IsCoupon:    dbPosclient.IsCoupon,
			AccountName: dbPosclient.AccountName,
			AccountCode: dbPosclient.AccountCode,
		}
	}

	return dbPosclient, nil
}
