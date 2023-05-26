package controller

import (
	"net/http"

	errortype "example.com/go-crud-api/common/errorType"
	repo "example.com/go-crud-api/repository/posConfig"
	"github.com/gin-gonic/gin"
)

type PosConfigController struct {
	posConfigRepo repo.PosConfigRepository
}

func NewPosConfigController(posConfigRepo repo.PosConfigRepository) *PosConfigController {
	return &PosConfigController{
		posConfigRepo: posConfigRepo,
	}
}

func (r *PosConfigController) GetPosConfig(c *gin.Context) {
	var req repo.GetPosClientReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorCode": errortype.BadRequestPayload,
			"errorDetail": errortype.ErrorMessage(errortype.BadRequestPayload)})
		return
	}
	posClientID := req.PosClientID
	posConfig, err := r.posConfigRepo.GetPosConfig(posClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode":   errortype.ServerError,
			"errorDetail": errortype.ErrorMessage(errortype.ServerError),
		})
		return
	}

	c.JSON(http.StatusOK, posConfig)
}
