package sale

import (
	"github.com/SiwaleK/ProdGroup/db/config"
	"github.com/SiwaleK/ProdGroup/model"
	"github.com/gin-gonic/gin"
)

func GetPaymentMethod(c *gin.Context) {
	var items []model.Payment_method
	result := config.DB.Find(&items)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, items)
}
