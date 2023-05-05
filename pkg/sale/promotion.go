package sale

import (
	"net/http"

	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PromotionHandler struct {
	repo repository.PromotionRepository
}

func NewPromotionHandler(repo repository.PromotionRepository) *PromotionHandler {
	return &PromotionHandler{repo: repo}
}

func (h *PromotionHandler) GetPromotionByID(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID123"})
		return
	}

	promotion, err := h.repo.GetPromotionByID(c.Request.Context(), parsedID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get promotion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"promotion": promotion})
}
