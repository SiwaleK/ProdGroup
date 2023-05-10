package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

type PaymentConfigRepository interface {
	GetPaymentConfig(ctx context.Context) ([]GetPaymentConfigRowWithPrinter, error)
	UpsertPaymentConfig(ctx context.Context, arg db.UpsertPaymentConfigParams) error
}

type DBPaymentConfigRepository struct {
	db *db.Queries
}

func NewPaymentConfigRepository(db *db.Queries) PaymentConfigRepository {
	return &DBPaymentConfigRepository{
		db: db,
	}
}

type Printerconfig struct {
	Value int    `json:"value"`
	Title string `json:"title"`
}

type GetPaymentConfigRowWithPrinter struct {
	db.GetPaymentConfigRow
	Printerconfig []Printerconfig `json:"printerconfig"`
}

func (r *DBPaymentConfigRepository) GetPaymentConfig(ctx context.Context) ([]GetPaymentConfigRowWithPrinter, error) {
	dbPaymentconfigs, err := r.db.GetPaymentConfig(ctx)
	if err != nil {
		return nil, err
	}

	paymentconfigs := make([]GetPaymentConfigRowWithPrinter, len(dbPaymentconfigs))
	for i, dbPaymentconfig := range dbPaymentconfigs {
		printerType, ok := dbPaymentconfig.PrinterType.(int)
		if !ok {
			// Handle the case where the PrinterType is not an int
			// You can choose to set a default value or return an error, depending on your requirements
			printerType = 0 // Set a default value of 0
		}

		printerConfig := Printerconfig{
			Value: printerType,
			Title: "test",
		}

		paymentconfig := GetPaymentConfigRowWithPrinter{
			GetPaymentConfigRow: db.GetPaymentConfigRow{
				IsCash:      dbPaymentconfig.IsCash,
				IsQrcode:    dbPaymentconfig.IsQrcode,
				IsPaotang:   dbPaymentconfig.IsPaotang,
				IsTongfah:   dbPaymentconfig.IsTongfah,
				IsCoupon:    dbPaymentconfig.IsCoupon,
				AccountName: dbPaymentconfig.AccountName,
				AccountCode: dbPaymentconfig.AccountCode,
				PrinterType: dbPaymentconfig.PrinterType,
			},
			Printerconfig: []Printerconfig{printerConfig},
		}
		paymentconfigs[i] = paymentconfig
	}

	return paymentconfigs, nil
}

func (repo *DBPaymentConfigRepository) UpsertPaymentConfig(ctx context.Context, arg db.UpsertPaymentConfigParams) error {
	err := repo.db.UpsertPaymentConfig(ctx, db.UpsertPaymentConfigParams{
		IsCash:      arg.IsCash,
		IsQrcode:    arg.IsQrcode,
		IsPaotang:   arg.IsPaotang,
		IsTongfah:   arg.IsTongfah,
		IsCoupon:    arg.IsCoupon,
		PrinterType: arg.PrinterType,
		//AccountName: arg.AccountName,
		//AccountCode: arg.AccountCode,
	})
	return err
}
