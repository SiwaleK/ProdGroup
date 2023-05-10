package controller

import (
	"net/http"
	"strconv"

	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
)

type ProdgroupHandlerImpl struct {
	repo repository.ProdgroupRepository
}

func NewProdgroupHandler(repo repository.ProdgroupRepository) *ProdgroupHandlerImpl {
	return &ProdgroupHandlerImpl{
		repo: repo,
	}
}

func (h *ProdgroupHandlerImpl) GetProdgroup(c *gin.Context) {
	prodgroups, err := h.repo.GetProdgroup(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get prodgroups56"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"prodgroups": prodgroups})
}
func (h *ProdgroupHandlerImpl) GetProdgroupByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
		return
	}

	prodgroups, err := h.repo.GetProdgroupByID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get prodgroup"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"prodgroups": prodgroups})
}
