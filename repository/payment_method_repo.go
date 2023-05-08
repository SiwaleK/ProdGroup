package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

// type PaymentMethodRepository interface {
// 	GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error)
// }

// type DBPaymentMethodRepository struct {
// 	db *sql.DB
// }

// func NewPaymentMethodRepository(db *sql.DB) *DBPaymentMethodRepository {
// 	return &DBPaymentMethodRepository{
// 		db: db,
// 	}
// }

// func (r *DBPaymentMethodRepository) GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error) {
// 	paymentmethods := []model.PaymentMethod{}
// 	rows, err := r.db.QueryContext(ctx, "SELECT * FROM paymentmethod")
// 	if err != nil {
// 		fmt.Println("Error querying database:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var pm model.PaymentMethod
// 		err := rows.Scan(&pm.PaymentMethodID, &pm.PaymentName)
// 		if err != nil {
// 			fmt.Println("Error scanning rows:", err)
// 			return nil, err
// 		}
// 		paymentmethods = append(paymentmethods, pm)
// 	}

// 	if err = rows.Err(); err != nil {
// 		fmt.Println("Error iterating over rows:", err)
// 		return nil, err
// 	}
// 	return paymentmethods, nil
// }

///// v1 fix

// type Queries struct {
// 	db *sql.DB
// }

// func New(db *sql.DB) *Queries {
// 	return &Queries{
// 		db: db,
// 	}
// }

// type PaymentMethodRepository interface {
// 	GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error)
// }
// type DBPaymentMethodRepository struct {
// 	db      *sql.DB
// 	queries *db.Queries
// }

// func NewPaymentMethodRepository(db *sql.DB) *DBPaymentMethodRepository {
// 	return &DBPaymentMethodRepository{
// 		db:      db,
// 		queries: db.New(db),
// 	}
// }

// func (r *DBPaymentMethodRepository) GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error) {
// 	paymentmethods := []model.PaymentMethod{}
// 	paymentMethods, err := r.queries.GetPaymentMethod(ctx)
// 	if err != nil {
// 		fmt.Println("Error querying database:", err)
// 		return nil, err
// 	}

// 	for _, pm := range paymentMethods {
// 		paymentmethods = append(paymentmethods, model.PaymentMethod{
// 			PaymentMethodID: pm.Paymentmethodid,
// 			PaymentName:     pm.Paymentname,
// 		})
// 	}

// 	return paymentmethods, nil
// }

//v2

// type PaymentMethodRepository interface {
// 	GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error)
// }

// type DBPaymentMethodRepository struct {
// 	db      *sql.DB
// 	queries *sqlc.Queries
// }

// func NewPaymentMethodRepository(db *sql.DB) *DBPaymentMethodRepository {
// 	return &DBPaymentMethodRepository{
// 		db:      db,
// 		queries: sqlc.New(db),
// 	}
// }

// func (r *DBPaymentMethodRepository) GetPaymentMethods(ctx context.Context) ([]model.PaymentMethod, error) {
// 	paymentmethods := []model.PaymentMethod{}
// 	paymentMethods, err := r.queries.GetPaymentMethod(ctx)
// 	if err != nil {
// 		fmt.Println("Error querying database:", err)
// 		return nil, err
// 	}

// 	for _, pm := range paymentMethods {
// 		paymentmethods = append(paymentmethods, model.PaymentMethod{
// 			PaymentMethodID: pm.Paymentmethodid,
// 			PaymentName:     pm.Paymentname,
// 		})
// 	}

// 	return paymentmethods, nil
// }

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
