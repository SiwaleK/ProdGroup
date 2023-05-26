package controller

import (
	"encoding/json"
	"net/http"

	errortype "example.com/go-crud-api/common/errorType"
	"example.com/go-crud-api/db/db"
	repo "example.com/go-crud-api/repository/promotion"

	"github.com/gin-gonic/gin"
)

type PromotionController struct {
	promotionRepo repo.PromotionRepository
}

func NewPromotionController(promotionRepo repo.PromotionRepository) *PromotionController {
	return &PromotionController{
		promotionRepo: promotionRepo,
	}
}

func (r *PromotionController) CreatePromotionDiscount(c *gin.Context) {
	var req repo.PostPromotionDiscountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload)})
		return
	}

	condition := repo.ConditionDiscount{
		Discount: req.Condition.Discount,
	}

	conditionBytes, err := json.Marshal(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errortype.ServerError, "errorDetail": errortype.ErrorMessage(errortype.ServerError)})
		return
	}

	conditionString := string(conditionBytes)

	var args []db.PromotionAppliedItemsID

	for _, item := range req.AppliedItemsID {
		a := db.PromotionAppliedItemsID{
			Promotionid:       req.Promotionid,
			Skuid:             item.Skuid,
			Promotiondetailid: item.PromotiondetailID,
		}

		args = append(args, a)
	}

	err = r.promotionRepo.CreatePromotionAppliedItem(args)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	// Save other columns in main table
	arg := db.Promotion{
		Promotionid:    req.Promotionid,
		Promotiontitle: req.Promotiontitle,
		Promotiontype:  req.Promotiontype,
		Startdate:      req.Startdate,
		Enddate:        req.Enddate,
		Description:    req.Description,
		Conditions:     conditionString,
	}

	err = r.promotionRepo.CreatePromotion([]db.Promotion{arg})
	if err != nil {
		// Handle specific errors
		if err.Error() == "validation_error" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       errortype.BadRequestPayload,
				"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
			})
			return
		} else if err.Error() == "database_error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       errortype.ServerError,
				"errorDetail": errortype.ErrorMessage(errortype.ServerError),
			})
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{})

}

func (r *PromotionController) CreatePromotionAFREEB(c *gin.Context) {
	//var input Input
	var req repo.PostPromotionAFREEBRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload)})
		return
	}

	condition := repo.ConditionAFREEB{
		MinimumAmountToEnable: req.Condition.MinimumAmountToEnable,
		FreeAmount:            req.Condition.FreeAmount,
		PremiumItemsId:        req.Condition.PremiumItemsId,
	}
	conditionBytes, err := json.Marshal(condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errortype.ServerError, "errorDetail": errortype.ErrorMessage(errortype.ServerError)})
		return
	}

	conditionString := string(conditionBytes)

	var args []db.PromotionAppliedItemsID

	for _, item := range req.AppliedItemsID {
		a := db.PromotionAppliedItemsID{
			Promotionid:       req.Promotionid,
			Skuid:             item.Skuid,
			Promotiondetailid: item.PromotiondetailID,
		}

		args = append(args, a)
	}

	err = r.promotionRepo.CreatePromotionAppliedItem(args)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	// Save other columns in main table
	arg := db.Promotion{
		Promotionid:    req.Promotionid,
		Promotiontitle: req.Promotiontitle,
		Promotiontype:  req.Promotiontype,
		Startdate:      req.Startdate,
		Enddate:        req.Enddate,
		Description:    req.Description,
		Conditions:     conditionString,
	}

	err = r.promotionRepo.CreatePromotion([]db.Promotion{arg})
	if err != nil {
		// Handle specific errors
		if err.Error() == "validation_error" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       errortype.BadRequestPayload,
				"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
			})
			return
		} else if err.Error() == "database_error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       errortype.ServerError,
				"errorDetail": errortype.ErrorMessage(errortype.ServerError),
			})
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{})

}

func (r *PromotionController) CreatePromotionStepPurchase(c *gin.Context) {

	var req repo.PostPromotionStepPurchaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorCode": errortype.BadRequestPayload, "errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload)})
		return
	}

	specialPriceConditions := repo.ConditionStepPurchase{
		SpecialPriceAtXItemConditionDetail: make([]repo.SpecialPriceAtXItemConditionDetail, len(req.Condition.SpecialPriceAtXItemConditionDetail)),
	}

	for i, condition := range req.Condition.SpecialPriceAtXItemConditionDetail {
		specialPriceConditions.SpecialPriceAtXItemConditionDetail[i].MinimumItemToEnable = condition.MinimumItemToEnable
		specialPriceConditions.SpecialPriceAtXItemConditionDetail[i].Discount = condition.Discount
	}

	conditionBytes, err := json.Marshal(specialPriceConditions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorCode": errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError)})
		return
	}

	conditionString := string(conditionBytes)

	var args []db.PromotionAppliedItemsID

	for _, item := range req.AppliedItemsID {
		a := db.PromotionAppliedItemsID{
			Promotionid:       req.Promotionid,
			Skuid:             item.Skuid,
			Promotiondetailid: item.PromotiondetailID,
		}

		args = append(args, a)
	}

	err = r.promotionRepo.CreatePromotionAppliedItem(args)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	// Save other columns in main table
	arg := db.Promotion{
		Promotionid:    req.Promotionid,
		Promotiontitle: req.Promotiontitle,
		Promotiontype:  req.Promotiontype,
		Startdate:      req.Startdate,
		Enddate:        req.Enddate,
		Description:    req.Description,
		Conditions:     conditionString,
	}

	err = r.promotionRepo.CreatePromotion([]db.Promotion{arg})
	if err != nil {
		// Handle specific errors
		if err.Error() == "validation_error" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       errortype.BadRequestPayload,
				"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
			})
			return
		} else if err.Error() == "database_error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errorCode":   errortype.ServerError,
				"errorDetail": errortype.ErrorMessage(errortype.ServerError),
			})
			return
		}

		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{"errorCode": errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError)})
		return
	}

	// Return the JSON response
	c.JSON(http.StatusOK, gin.H{})

}
