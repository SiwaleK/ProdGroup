package controller

import (
	"net/http"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
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

type InternalServerError struct {
	Message string
}

// Implement the error interface for InternalServerError
func (e InternalServerError) Error() string {
	return e.Message
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
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

type createUserRequest struct {
	IsCash      interface{} `json:"is_cash"`
	IsQRCode    interface{} `json:"is_qrcode"`
	IsPaoTang   interface{} `json:"is_paotang"`
	IsTongFah   interface{} `json:"is_tongfah"`
	IsCoupon    interface{} `json:"is_coupon"`
	PrinterType interface{} `json:"printer_type"`
	AccountName string      `json:"account_name"`
	AccountCode string      `json:"account_code"`
}

func (h *PaymentconfigHandlerImpl) UpsertPaymentConfigHandler(c *gin.Context) {
	// Parse request parameters from the Gin context
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create UpsertPaymentConfigParams
	params := db.UpsertPaymentConfigParams{
		IsCash:      req.IsCash,
		IsQrcode:    req.IsQRCode,
		IsPaotang:   req.IsPaoTang,
		IsTongfah:   req.IsTongFah,
		IsCoupon:    req.IsCoupon,
		PrinterType: req.PrinterType,
		//AccountName: sql.NullString{String: req.AccountName, Valid: req.AccountName != ""},
		//AccountCode: sql.NullString{String: req.AccountCode, Valid: req.AccountCode != ""},
	}

	// Call UpsertPaymentConfig method
	err := h.repo.UpsertPaymentConfig(c.Request.Context(), params)
	if err != nil {
		// Handle specific errors
		if err.Error() == "validation_error" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Validation Error",
			})
			return
		} else if err.Error() == "database_error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database Error",
			})
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Internal Server Error",
			"detail": err.Error(),
		})
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment configuration updated successfully",
	})
}

// func (repo *PaymentconfigHandlerImpl) UpsertPaymentConfig(ctx context.Context, arg db.UpsertPaymentConfigParams) error {
// 	// Convert the payload to JSON
// 	payload, err := json.Marshal(arg)
// 	if err != nil {
// 		return err
// 	}

// 	// Create a new request with the payload
// 	req, err := http.NewRequest("POST", "localhost:8080/sale/api/v1/PaymentConfig", bytes.NewBuffer(payload))
// 	if err != nil {
// 		return err
// 	}

// 	// Set the appropriate headers
// 	req.Header.Set("Content-Type", "application/json")

// 	// Send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Check the response status code
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("request failed with status code: %d", resp.StatusCode)
// 	}

// 	// Handle the response
// 	// ...

// 	return nil
// }
