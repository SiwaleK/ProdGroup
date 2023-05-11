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

type ConditionAFREEB struct {
	MinimumAmountToEnable int
	FreeAmount            int
	PremiumItemsId        []string
}

type ConditionDiscount struct {
	Discount int
}

type AppliedItems struct {
	ID                string `json:"id"`
	PromotiondetailID string `json:"promotiondetailID"`
	SKUID             string `json:"skuid"`
}

type PostPromotionDiscountRequest struct {
	Promotionid    *string   `json:"id"`
	Promotiontitle *string   `json:"title"`
	Promotiontype  int32     `json:"promotiontype"`
	Startdate      time.Time `json:"startDate"`
	Enddate        time.Time `json:"endDate"`
	Description    *string   `json:"text"`
	Condition      struct {
		Discount int
	}
	AppliedItemsID []struct {
		PromotiondetailID *string `json:"promotiondetail_id"`
		Skuid             *string `json:"skuid"`
	} `json:"appliedItemsId"`
}

type PostPromotionAFREEBRequest struct {
	Promotionid    *string   `json:"id"`
	Promotiontitle *string   `json:"title"`
	Promotiontype  int32     `json:"promotiontype"`
	Startdate      time.Time `json:"startdate"`
	Enddate        time.Time `json:"enddate"`
	Description    *string   `json:"text"`
	Condition      struct {
		MinimumAmountToEnable int
		FreeAmount            int
		PremiumItemsId        []string
	} `json:"condition"`
	AppliedItemsID []struct {
		PromotiondetailID *string `json:"promotiondetail_id"`
		Skuid             *string `json:"skuid"`
	} `json:"appliedItemsId"`
}

type ConditionStepPurchase struct {
	SpecialPriceAtXItemConditionDetail []SpecialPriceAtXItemConditionDetail `json:"specialPriceAtXItemConditionDetail"`
}

type SpecialPriceAtXItemConditionDetail struct {
	MinimumItemToEnable int `json:"minimumItemToEnable"`
	Discount            int `json:"discount"`
}

type PostPromotionStepPurchaseRequest struct {
	Promotionid    *string   `json:"id"`
	Promotiontitle *string   `json:"title"`
	Promotiontype  int32     `json:"promotiontype"`
	Startdate      time.Time `json:"startdate"`
	Enddate        time.Time `json:"enddate"`
	Description    *string   `json:"text"`
	Condition      struct {
		SpecialPriceAtXItemConditionDetail []struct {
			MinimumItemToEnable int `json:"minimumItemToEnable"`
			Discount            int `json:"discount"`
		} `json:"specialPriceAtXItemConditionDetail"`
	} `json:"condition"`
	AppliedItemsID []struct {
		PromotiondetailID *string `json:"promotiondetail_id"`
		Skuid             *string `json:"skuid"`
	} `json:"appliedItemsId"`
}

func (h PromotionHandler) PostDiscountPromotion(c *gin.Context) {

	var req PostPromotionDiscountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
		return
	}
	condition := ConditionDiscount{
		Discount: req.Condition.Discount,
	}

	conditionBytes, err := json.Marshal(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal condition"})
		return
	}

	appliedItemsIDs := make([]struct {
		PromotiondetailID *string `json:"promotiondetail_id"`
		Skuid             *string `json:"skuid"`
	}, len(req.AppliedItemsID))

	for i, item := range req.AppliedItemsID {
		appliedItemsIDs[i].PromotiondetailID = item.PromotiondetailID
		appliedItemsIDs[i].Skuid = item.Skuid
	}
	rawCondition := json.RawMessage(conditionBytes)
	arg := db.PostPromotionParams{
		Promotionid:       req.Promotionid,
		Promotiontitle:    req.Promotiontitle,
		Promotiontype:     req.Promotiontype,
		Startdate:         req.Startdate,
		Enddate:           req.Enddate,
		Description:       req.Description,
		Condition:         rawCondition,
		Skuid:             appliedItemsIDs[0].Skuid,
		PromotiondetailID: appliedItemsIDs[0].PromotiondetailID,
	}

	err = h.repo.PostPromotion(c.Request.Context(), arg)
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

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{"message": "Promotion created successfully"})
}

func (h PromotionHandler) PostPromotionAFREEB(c *gin.Context) {
	//var input Input
	var req PostPromotionAFREEBRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
		return
	}

	condition := ConditionAFREEB{
		MinimumAmountToEnable: req.Condition.MinimumAmountToEnable,
		FreeAmount:            req.Condition.FreeAmount,
		PremiumItemsId:        req.Condition.PremiumItemsId,
	}

	conditionBytes, err := json.Marshal(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal condition"})
		return
	}

	appliedItemsIDs := make([]struct {
		PromotiondetailID *string `json:"promotiondetail_id"`
		Skuid             *string `json:"skuid"`
	}, len(req.AppliedItemsID))

	for i, item := range req.AppliedItemsID {
		appliedItemsIDs[i].PromotiondetailID = item.PromotiondetailID
		appliedItemsIDs[i].Skuid = item.Skuid
	}

	rawCondition := json.RawMessage(conditionBytes)
	arg := db.PostPromotionParams{
		Promotionid:       req.Promotionid,
		Promotiontitle:    req.Promotiontitle,
		Promotiontype:     req.Promotiontype,
		Startdate:         req.Startdate,
		Enddate:           req.Enddate,
		Description:       req.Description,
		Condition:         rawCondition,
		Skuid:             appliedItemsIDs[0].Skuid,
		PromotiondetailID: appliedItemsIDs[0].PromotiondetailID,
	}

	err = h.repo.PostPromotion(c.Request.Context(), arg)
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

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{"message": "Promotion created successfully"})

}

func (h PromotionHandler) PostPromotionStepPurchase(c *gin.Context) {

	var req PostPromotionStepPurchaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error", "detail": err.Error()})
		return

	}

	specialPriceConditions := ConditionStepPurchase{
		SpecialPriceAtXItemConditionDetail: make([]SpecialPriceAtXItemConditionDetail, len(req.Condition.SpecialPriceAtXItemConditionDetail)),
	}

	for i, condition := range req.Condition.SpecialPriceAtXItemConditionDetail {
		specialPriceConditions.SpecialPriceAtXItemConditionDetail[i].MinimumItemToEnable = condition.MinimumItemToEnable
		specialPriceConditions.SpecialPriceAtXItemConditionDetail[i].Discount = condition.Discount
	}

	conditionBytes, err := json.Marshal(specialPriceConditions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal condition"})
		return
	}

	rawCondition := json.RawMessage(conditionBytes)
	appliedItemsIDs := make([]struct {
		PromotiondetailID *string `json:"promotiondetail_id"`
		Skuid             *string `json:"skuid"`
	}, len(req.AppliedItemsID))

	for i, item := range req.AppliedItemsID {
		appliedItemsIDs[i].PromotiondetailID = item.PromotiondetailID
		appliedItemsIDs[i].Skuid = item.Skuid
	}

	arg := db.PostPromotionParams{
		Promotionid:       req.Promotionid,
		Promotiontitle:    req.Promotiontitle,
		Promotiontype:     req.Promotiontype,
		Startdate:         req.Startdate,
		Enddate:           req.Enddate,
		Description:       req.Description,
		Condition:         rawCondition,
		Skuid:             appliedItemsIDs[0].Skuid,
		PromotiondetailID: appliedItemsIDs[0].PromotiondetailID,
	}

	err = h.repo.PostPromotion(c.Request.Context(), arg)
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

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{"message": "Promotion created successfully"})

}
