package sale

// func GetPaymentConfig(c *gin.Context) {
// 	var paymentconfig []model.Paymentconfig
// 	result := config.DB.Find(&paymentconfig)
// 	if result.Error != nil {
// 		c.JSON(500, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(200, paymentconfig)
// }

// func PostPaymentConfig(c *gin.Context) {
// 	var paymentconfig model.Paymentconfig
// 	id := c.Param("paymentid")
// 	if err := config.DB.Where("paymentid = ?", id).First(&paymentconfig).Error; err != nil {
// 		c.JSON(404, gin.H{"error": "paymentconfig not found"})
// 		return
// 	}
// 	if err := c.BindJSON(&paymentconfig); err != nil {
// 		c.JSON(500, gin.H{"error": "Invalid request payload"})
// 		return
// 	}
// 	if err := config.DB.Save(&paymentconfig).Error; err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to update paymentconfig"})
// 		return
// 	}
// 	c.JSON(200, paymentconfig)
// }
