package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

type PaymentMethodRepository interface {
	GetPaymentMethods(ctx context.Context) ([]db.PaymentMethod, error)
}

type paymentMethodRepository struct {
	db *db.Queries
}

func NewPaymentMethodRepository(db *db.Queries) PaymentMethodRepository {
	return &paymentMethodRepository{
		db: db,
	}
}

func (r *paymentMethodRepository) GetPaymentMethods(ctx context.Context) ([]db.PaymentMethod, error) {
	dbPaymentMethods, err := r.db.GetPaymentMethod(ctx)
	if err != nil {
		return nil, err
	}

	// Convert []db.PaymentMethod to []model.PaymentMethod
	paymentMethods := make([]db.PaymentMethod, len(dbPaymentMethods))
	for i, dbPaymentMethod := range dbPaymentMethods {
		paymentMethods[i] = db.PaymentMethod{
			Paymentmethodid: dbPaymentMethod.Paymentmethodid,
			Paymentname:     dbPaymentMethod.Paymentname,
		}
	}

	return paymentMethods, nil
}
