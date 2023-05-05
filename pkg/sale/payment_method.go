package sale

import (
	"net/http"

	"github.com/SiwaleK/ProdGroup/repository"
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

type PaymentMethodHandler struct {
	repo repository.PaymentMethodRepository
}

func NewPaymentMethodHandler(repo repository.PaymentMethodRepository) *PaymentMethodHandler {
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
