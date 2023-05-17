package controller

import (
	"encoding/json"
	"fmt"
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

type PromptPayData struct {
	AccountName *string `json:"promptPayAccountName"`
	AccountCode *string `json:"promptPayNo"`
}

type PaymentConfig struct {
	PromptPayData PromptPayData `json:"promptPayData"`
	Items         []Item        `json:"items"`
}

type Item struct {
	PaymentMethodID int32   `json:"paymentMethodID"`
	PaymentName     *string `json:"paymentName"`
	IsEnable        *bool   `json:"isEnable"`
}

type GetPosClientMethodReq struct {
	PosClientID string `json:"posClientID"`
}

func (h *PaymentconfigHandlerImpl) GetPosClientMethod(c *gin.Context) {
	var req GetPosClientMethodReq

	// Bind the JSON data to the req struct
	if err := c.ShouldBindJSON(&req); err != nil {
		// Handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	posClientID := &req.PosClientID
	//fmt.Printf(posClientID)
	posClients, err := h.repo.GetPosClientMethod(c.Request.Context(), posClientID)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get payment configuration", "detail": err.Error()})
		return
	}

	paymentMethod, err := h.repo.GetPaymentMethod(c.Request.Context())
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get payment configuration GetpaymentMethod", "detail": err.Error()})
		return
	}

	promptPayData := PromptPayData{
		AccountName: posClients[0].AccountName,
		AccountCode: posClients[0].AccountCode,
	}

	paymentConfig := PaymentConfig{
		PromptPayData: promptPayData,
		Items:         make([]Item, len(paymentMethod)),
	}

	// Iterate over the items slice
	for j, item := range paymentMethod {
		isEnable := false

		// Check if paymentMethodID and corresponding is_* field match in posclient table
		for _, posClient := range posClients {
			switch item.Paymentmethodid {
			case 1:
				if item.Paymentmethodid == 1 && *posClient.IsCash == 1 {
					isEnable = true

				}
				fmt.Printf("paymentid %d", item.Paymentmethodid)
				fmt.Printf("cash %d", &posClient.IsCash)

			case 2:
				if item.Paymentmethodid == 2 && *posClient.IsQrcode == 1 {
					isEnable = true
				}
			case 3:
				if item.Paymentmethodid == 3 && *posClient.IsPaotang == 1 {
					isEnable = true
				}
			case 4:
				if item.Paymentmethodid == 4 && *posClient.IsTongfah == 1 {
					isEnable = true
				}
			case 5:
				if item.Paymentmethodid == 5 && *posClient.IsCoupon == 1 {
					isEnable = true
				}
			}
		}

		paymentConfig.Items[j] = Item{
			PaymentMethodID: item.Paymentmethodid,
			PaymentName:     item.Paymentname,
			IsEnable:        &isEnable,
		}
	}

	// Marshal the paymentConfig to JSON
	jsonData, err := json.Marshal(paymentConfig)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal response"})
		return
	}

	// Set the response headers
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)

	// Write the JSON response
	_, err = c.Writer.Write(jsonData)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write response"})
		return
	}
}
