package controller

import (
	"net/http"
	"time"

	errortype "example.com/go-crud-api/common/errorType"
	repo "example.com/go-crud-api/repository/receiptHistory"

	"github.com/gin-gonic/gin"
)

type ReceiptHistoryController struct {
	receiptRepo repo.ReceiptHistoryRepository
}

func NewReceiptHistoryController(receiptRepo repo.ReceiptHistoryRepository) *ReceiptHistoryController {
	return &ReceiptHistoryController{
		receiptRepo: receiptRepo,
	}
}

func (r *ReceiptHistoryController) GetReceiptHistoryByID(c *gin.Context) {
	saleOrderID := c.Query("saleOrderID")
	if saleOrderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
		})
		return
	}

	res, err := r.receiptRepo.GetReceiptHistoryByID(saleOrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ReceiptHistory": res,
	})
}

func (r *ReceiptHistoryController) GetReceiptHistoryByDate(c *gin.Context) {
	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
		})
		return
	}

	res, err := r.receiptRepo.GetReceiptHistoryByDate(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
