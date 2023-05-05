package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SiwaleK/ProdGroup/model"
)

type PaymentMethodRepository interface {
	GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error)
}

type DBPaymentMethodRepository struct {
	db *sql.DB
}

func NewPaymentMethodRepository(db *sql.DB) *DBPaymentMethodRepository {
	return &DBPaymentMethodRepository{
		db: db,
	}
}

func (r *DBPaymentMethodRepository) GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error) {
	paymentmethods := []model.PaymentMethod{}
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM paymentmethod")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pm model.PaymentMethod
		err := rows.Scan(&pm.PaymentMethodID, &pm.PaymentName)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			return nil, err
		}
		paymentmethods = append(paymentmethods, pm)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}
	return paymentmethods, nil
}
