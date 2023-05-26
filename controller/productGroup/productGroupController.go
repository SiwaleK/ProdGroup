package controller

import (
	"net/http"
	"strconv"

	errortype "example.com/go-crud-api/common/errorType"
	repo "example.com/go-crud-api/repository/productGroup"

	"github.com/gin-gonic/gin"
)

type ProductGroupController struct {
	productGroupRepo repo.ProductGroupRepository
}

func NewProductGroupController(productGroupRepo repo.ProductGroupRepository) *ProductGroupController {
	return &ProductGroupController{
		productGroupRepo: productGroupRepo,
	}
}

func (r *ProductGroupController) GetProductGroup(c *gin.Context) {
	res, err := r.productGroupRepo.GetProductGroup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}

func (r *ProductGroupController) GetProductGroupByID(c *gin.Context) {
	prodGroupIDStr := c.Param("prodgroupID")
	prodGroupID, err := strconv.Atoi(prodGroupIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorCode":   errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload),
		})
		return
	}

	res, err := r.productGroupRepo.GetProductGroupByID(prodGroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{

			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
