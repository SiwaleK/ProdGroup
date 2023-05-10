package controller

import (
	"encoding/json"
	"net/http"
	"time"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
)

type PromotionHandler struct {
	repo repository.PromotionRepository
}

func NewPromotionHandler(repo repository.PromotionRepository) *PromotionHandler {
	return &PromotionHandler{repo: repo}
}

// func (h *PromotionHandler) GetPromotionByID(c *gin.Context) {
// 	id := c.Param("id")
// 	parsedID, err := uuid.Parse(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
// 		return
// 	}

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID123"})
// 		return
// 	}

// 	promotion, err := h.repo.GetPromotionByID(c.Request.Context(), parsedID.String())
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get promotion"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"promotion": promotion})
// }

// type PostPromotionResponse struct {
// 	ID             string    `json:"id"`
// 	Text           string    `json:"text"`
// 	StartDate      time.Time `json:"startDate"`
// 	EndDate        time.Time `json:"endDate"`
// 	Condition      string    `json:"condition"`
// 	AppliedItemsID []string  `json:"appliedItemsId"`
// }

// type Condition struct {
// 	Discount              int      `json:"discount"`
// 	MinimumAmountToEnable int      `json:"minimumAmountToEnable"`
// 	FreeAmount            int      `json:"freeAmount"`
// 	PremiumItemsID        []string `json:"premiumItemsId"`
// }

// func (h *PromotionHandler) PostPromotion(c *gin.Context) {
// 	var req db.PostPromotionParams
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
// 		return
// 	}

// 	var condition json.RawMessage
// 	if err := json.Unmarshal([]byte(req.Condition), &condition); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid condition format",
// 		})
// 		return
// 	}

// 	arg := db.PostPromotionParams{
// 		Promotionid:       req.Promotionid,
// 		Promotiontype:     req.Promotiontype,
// 		Startdate:         req.Startdate,
// 		Enddate:           req.Enddate,
// 		Description:       req.Description,
// 		Condition:         condition,
// 		Skuid:             req.Skuid,
// 		PromotiondetailID: req.PromotiondetailID,
// 	}

// 	err := h.repo.PostPromotion(c.Request.Context(), arg)
// 	if err != nil {
// 		// Handle specific errors
// 		if err.Error() == "validation_error" {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Validation Error",
// 			})
// 			return
// 		} else if err.Error() == "database_error" {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Database Error",
// 			})
// 			return
// 		}

// 		// Handle other errors
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error":  "Internal Server Error",
// 			"detail": err.Error(),
// 		})
// 		return
// 	}

// 	// Construct the response
// 	response := PostPromotionResponse{
// 		ID: arg.Promotionid.String,
// 		//Text:      req.Text,
// 		StartDate: req.Startdate,
// 		EndDate:   req.Enddate,
// 		Condition: string(req.Condition),
// 		AppliedItemsID: []string{
// 			req.Promotionid.String,
// 			req.PromotiondetailID.String,
// 			req.Skuid.String,
// 		},
// 	}

// 	// Return the JSON response
// 	c.JSON(http.StatusOK, response)
// }

type PostPromotionDiscountResponse struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Condition struct {
		Discount int
	}
	AppliedItemsID AppliedItems `json:"appliedItemsId"`
}

type Condition struct {
	Discount int `json:"discount"`
}

type AppliedItems struct {
	ID                string `json:"id"`
	PromotiondetailID string `json:"promotiondetailID"`
	SKUID             string `json:"skuid"`
}

func (h PromotionHandler) PostDiscountPromotion(c *gin.Context) {
	//var input Input
	var req db.PostPromotionParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
		return
	}
	// // Marshal the Condition field to a JSON string
	condition := Condition{}
	if err := json.Unmarshal(req.Condition, &condition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid condition"})
		return
	}
	arg := db.PostPromotionParams{
		Promotionid:       req.Promotionid,
		Promotiontype:     req.Promotiontype,
		Startdate:         req.Startdate,
		Enddate:           req.Enddate,
		Description:       req.Description,
		Condition:         req.Condition,
		Skuid:             req.Skuid,
		PromotiondetailID: req.PromotiondetailID,
	}

	err := h.repo.PostPromotion(c.Request.Context(), arg)
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

	// Construct the response
	response := PostPromotionDiscountResponse{
		ID:        req.Promotionid.String,
		Text:      req.Description.String,
		StartDate: req.Startdate,
		EndDate:   req.Enddate,
		Condition: struct{ Discount int }{
			Discount: condition.Discount,
		},
		AppliedItemsID: AppliedItems{
			ID:                req.Promotionid.String,
			PromotiondetailID: req.PromotiondetailID.String,
			SKUID:             req.Skuid.String,
		},
	}

	// Return the JSON response
	c.JSON(http.StatusOK, response)
}
