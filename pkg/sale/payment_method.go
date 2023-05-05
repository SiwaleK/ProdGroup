package sale

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/SiwaleK/ProdGroup/model"
	"github.com/gin-gonic/gin"
)

// func GetPaymentMethod(c *gin.Context) {
// 	var items []model.PaymentMethod
// 	result := config.DB.Find(&items)
// 	if result.Error != nil {
// 		c.JSON(500, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(200, items)
// }

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

type PaymentMethodHandler struct {
	repo PaymentMethodRepository
}

func NewPaymentMethodHandler(repo PaymentMethodRepository) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		repo: repo,
	}
}

func (h *PaymentMethodHandler) GetPaymentMethod(c *gin.Context) {
	paymentmethod, err := h.repo.GetPaymentMethods(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get payment method"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"paymentmethod": paymentmethod})
}
