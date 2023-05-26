package controller

import (
	"encoding/json"
	"net/http"

	errortype "example.com/go-crud-api/common/errorType"
	repo "example.com/go-crud-api/repository/paymentConfig"

	"github.com/gin-gonic/gin"
)

type PaymentConfigController struct {
	paymentConfigRepo repo.PaymentConfigRepository
}

func NewPaymentConfigController(paymentConfigRepo repo.PaymentConfigRepository) *PaymentConfigController {
	return &PaymentConfigController{
		paymentConfigRepo: paymentConfigRepo,
	}
}

func (r *PaymentConfigController) GetPaymentConfig(c *gin.Context) {
	var req repo.GetPosClientMethodReq
	if err := c.ShouldBindJSON(&req); err != nil {
		// Handle the error
		c.JSON(http.StatusBadRequest, gin.H{"errorCode": errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload)})
		return
	}

	posClientID := req.PosClientID

	posClients, err := r.paymentConfigRepo.GetPosClientConfig(posClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	paymentMethod, err := r.paymentConfigRepo.GetPaymentMethod()
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	PromptPayData := repo.PromptPayData{
		AccountName: posClients[0].Accountname,
		AccountCode: posClients[0].Accountcode,
	}
	paymentConfig := repo.PaymentConfig{
		PromptPayData: PromptPayData,
		Items:         make([]repo.Item, len(paymentMethod)),
	}

	for j, item := range paymentMethod {
		isEnable := false

		// Check if paymentMethodID and corresponding is_* field match in posclient table
		for _, posClient := range posClients {
			switch item.Paymentmethodid {
			case 1:
				if posClient.Iscash {
					isEnable = true
				}

			case 2:
				if posClient.Isqrcode {
					isEnable = true
				}
			case 3:
				if posClient.Ispaotang {
					isEnable = true
				}
			case 4:
				if posClient.Istongfah {
					isEnable = true
				}
			case 5:
				if posClient.Iscoupon {
					isEnable = true
				}
			}
		}

		paymentConfig.Items[j] = repo.Item{
			PaymentMethodID: item.Paymentmethodid,
			PaymentName:     item.Paymentname,
			IsEnable:        &isEnable,
		}
	}

	// Marshal the paymentConfig to JSON
	jsonData, err := json.Marshal(paymentConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}
	// Set the response headers
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)

	// Write the JSON response
	_, err = c.Writer.Write(jsonData)
	if err != nil {
		c.Error(err) // Gin will handle the error and respond accordingly
		return
	}

}
