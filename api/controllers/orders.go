// controllers/orders.go

package controllers

type CreateOrderInput struct {
	ID uint   `json:"id" binding:"required"`
	Items []OrderItem `json:"items" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	// Validate input
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Post order
	order := models.Order{Items: input.Items}
	models.DB.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}