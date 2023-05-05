package sale

import (
	"net/http"

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
