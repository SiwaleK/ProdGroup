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
	Discount int32
}

type AppliedItems struct {
	PromotiondetailID *string `json:"promotiondetail_ID"`
	Skuid             *string `json:"skuid"`
}

type PostPromotionDiscountRequest struct {
	Promotionid    *string           `json:"id"`
	Promotiontitle *string           `json:"title"`
	Promotiontype  int32             `json:"promotiontype"`
	Startdate      time.Time         `json:"startDate"`
	Enddate        time.Time         `json:"endDate"`
	Description    *string           `json:"text"`
	Condition      ConditionDiscount `json:"condition"`
	AppliedItemsID []AppliedItems    `json:"appliedItemsId"`
}

type PostPromotionAFREEBRequest struct {
	Promotionid    *string         `json:"id"`
	Promotiontitle *string         `json:"title"`
	Promotiontype  int32           `json:"promotiontype"`
	Startdate      time.Time       `json:"startdate"`
	Enddate        time.Time       `json:"enddate"`
	Description    *string         `json:"text"`
	Condition      ConditionAFREEB `json:"condition"`
	AppliedItemsID []AppliedItems  `json:"appliedItemsId"`
}

type ConditionStepPurchase struct {
	SpecialPriceAtXItemConditionDetail []SpecialPriceAtXItemConditionDetail `json:"specialPriceAtXItemConditionDetail"`
}

type SpecialPriceAtXItemConditionDetail struct {
	MinimumItemToEnable int `json:"minimumItemToEnable"`
	Discount            int `json:"discount"`
}

type PostPromotionStepPurchaseRequest struct {
	Promotionid    *string               `json:"id"`
	Promotiontitle *string               `json:"title"`
	Promotiontype  int32                 `json:"promotiontype"`
	Startdate      time.Time             `json:"startdate"`
	Enddate        time.Time             `json:"enddate"`
	Description    *string               `json:"text"`
	Condition      ConditionStepPurchase `json:"condition"`
	AppliedItemsID []AppliedItems        `json:"appliedItemsId"`
}

func (h PromotionHandler) PostDiscountPromotion(c *gin.Context) {
	var req PostPromotionDiscountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID",
			"errorDetail": err.Error()})
		return
	}

	condition := ConditionDiscount{
		Discount: req.Condition.Discount,
	}

	conditionBytes, err := json.Marshal(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal condition", "errorDetail": err.Error()})
		return
	}
	rawCondition := json.RawMessage(conditionBytes)

	var args []db.PostPromotionAppliedParams

	for _, item := range req.AppliedItemsID {
		a := db.PostPromotionAppliedParams{
			Promotionid:       req.Promotionid,
			Skuid:             item.Skuid,
			PromotiondetailID: item.PromotiondetailID,
		}

		args = append(args, a)
	}

	err = h.repo.PostPromotionAppliedItem(c.Request.Context(), args)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   "Failed to save appliedItemsId",
			"errorDetail": err.Error(),
		})
		return
	}

	// Save other columns in main table
	arg := db.PostPromotionTableParams{
		Promotionid:    req.Promotionid,
		Promotiontitle: req.Promotiontitle,
		Promotiontype:  req.Promotiontype,
		Startdate:      req.Startdate,
		Enddate:        req.Enddate,
		Description:    req.Description,
		Condition:      rawCondition,
	}

	err = h.repo.PostPromotion(c.Request.Context(), arg)
	if err != nil {
		// Handle specific errors
		if err.Error() == "validation_error" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       "Validation Error",
				"errorDetail": err.Error(),
			})
			return
		} else if err.Error() == "database_error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Database Error",
				"errorDetail": err.Error(),
			})
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   "Internal Server Error",
			"errorDetail": err.Error(),
		})
		return
	}

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{})

}

func (h PromotionHandler) PostPromotionAFREEB(c *gin.Context) {
	//var input Input
	var req PostPromotionAFREEBRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID", "errorDetail": err.Error()})
		return
	}

	condition := ConditionAFREEB{
		MinimumAmountToEnable: req.Condition.MinimumAmountToEnable,
		FreeAmount:            req.Condition.FreeAmount,
		PremiumItemsId:        req.Condition.PremiumItemsId,
	}

	conditionBytes, err := json.Marshal(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal condition", "errorDetail": err.Error()})
		return
	}
	rawCondition := json.RawMessage(conditionBytes)

	var args []db.PostPromotionAppliedParams

	for _, item := range req.AppliedItemsID {
		a := db.PostPromotionAppliedParams{
			Promotionid:       req.Promotionid,
			Skuid:             item.Skuid,
			PromotiondetailID: item.PromotiondetailID,
		}

		args = append(args, a)
	}

	err = h.repo.PostPromotionAppliedItem(c.Request.Context(), args)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   "Failed to save appliedItemsId",
			"errorDetail": err.Error(),
		})
		return
	}

	// Save other columns in main table
	arg := db.PostPromotionTableParams{
		Promotionid:    req.Promotionid,
		Promotiontitle: req.Promotiontitle,
		Promotiontype:  req.Promotiontype,
		Startdate:      req.Startdate,
		Enddate:        req.Enddate,
		Description:    req.Description,
		Condition:      rawCondition,
	}

	err = h.repo.PostPromotion(c.Request.Context(), arg)
	if err != nil {
		// Handle specific errors
		if err.Error() == "validation_error" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       "Validation Error",
				"errorDetail": err.Error(),
			})
			return
		} else if err.Error() == "database_error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Database Error",
				"errorDetail": err.Error(),
			})
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   "Internal Server Error",
			"errorDetail": err.Error(),
		})
		return
	}

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{})

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

	var args []db.PostPromotionAppliedParams

	for _, item := range req.AppliedItemsID {
		a := db.PostPromotionAppliedParams{
			Promotionid:       req.Promotionid,
			Skuid:             item.Skuid,
			PromotiondetailID: item.PromotiondetailID,
		}

		args = append(args, a)
	}

	err = h.repo.PostPromotionAppliedItem(c.Request.Context(), args)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   "Failed to save appliedItemsId",
			"errorDetail": err.Error(),
		})
		return
	}

	// Save other columns in main table
	arg := db.PostPromotionTableParams{
		Promotionid:    req.Promotionid,
		Promotiontitle: req.Promotiontitle,
		Promotiontype:  req.Promotiontype,
		Startdate:      req.Startdate,
		Enddate:        req.Enddate,
		Description:    req.Description,
		Condition:      rawCondition,
	}

	err = h.repo.PostPromotion(c.Request.Context(), arg)
	if err != nil {
		// Handle specific errors
		if err.Error() == "validation_error" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       "Validation Error",
				"errorDetail": err.Error(),
			})
			return
		} else if err.Error() == "database_error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Database Error",
				"errorDetail": err.Error(),
			})
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{})

}
