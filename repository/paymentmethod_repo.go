package repository

import (
	"context"
	"database/sql"

	"github.com/SiwaleK/ProdGroup/model"
)

type PaymentMethodRepository interface {
	GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error)
}

type DBPaymentMethodRepository struct {
	db *sql.DB
}

func NewProdgroupRepositroy(db *sql.DB) *DBPaymentMethodRepository {
	return &DBPaymentMethodRepository{
		db: db,
	}
}
