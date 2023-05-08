package controller

import (
	"net/http"

	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
)

type PaymentconfigHandlerImpl struct {
	repo repository.PaymentConfigRepository
}

func NewPaymentConfighandler(repo repository.PaymentConfigRepository) *PaymentconfigHandlerImpl {
	return &PaymentconfigHandlerImpl{
		repo: repo,
	}
}

func (h *PaymentconfigHandlerImpl) GetPaymentConfig(c *gin.Context) {
	paymentconfigs, err := h.repo.GetPaymentConfig(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get paymentconfigs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"paymentconfigs": paymentconfigs})
}
