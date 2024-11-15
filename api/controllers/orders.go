// controllers/orders.go

package controllers

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateOrderItem struct {
	Quantity  int `json:"quantity" binding:"required"`
	ProductID int `json:"product" binding:"required"`
}

type CreateOrderInput struct {
	Items []CreateOrderItem `json:"items" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	// Validate input
	input := new(CreateOrderInput)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Post order
	order := models.Order{}
	for line := range input.Items {
		models.DB.Where("id = ?", line.ProductID) // Fix later
		item := models.OrderItem{}
		order.Items = append(order.Items, item)
	}
	order := models.Order{Items: input.Items}
	models.DB.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}
