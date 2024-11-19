// controllers/orders.go

package controllers

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateOrderItem struct {
	Quantity  int `json:"quantity" binding:"required"`
	ProductID int `json:"product" binding:"required"`
}

type CreateOrderInput struct {
	Items []CreateOrderItem `json:"items" binding:"required"`
      Name string `json:"name" binding:"required"`
      Email string `json:"email" binding:"required"`
      Address1 string `json:"address_one" binding:"required"`
      Address2 string `json:"address_two" binding:"required"`
      City string `json:"city" binding:"required"`
      State string `json:"state" binding:"required"`
      Zip string `json:"zip" binding:"required"`
}

func FindOrders(c *gin.Context) {
	ordersData := new(models.OrderRequestData)
	if c.Query("orderID") != "" {
		models.DB.Where("orderID = ?", c.Query("OrderID")).Find(&ordersData.Orders)
	} else {
		models.DB.Find(&ordersData.Orders)
	}
	ordersData.Retrieved = time.Now()
	c.JSON(http.StatusOK, ordersData)
}

func CreateOrder(c *gin.Context) {
	// Validate input
	input := new(CreateOrderInput)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Order items, store in Order
	order := new(models.Order)
	for _, line := range input.Items {
		product := new(models.Product)
		models.DB.Where("id = ?", line.ProductID).First(product)
		item := models.OrderItem{}
		item.Quantity = line.Quantity
		item.Price = product.Price //Price at time of sale
		item.Product = product
		item.ProductID = product.ID
		order.Items = append(order.Items, item)
	}
	//Insert Order
	models.DB.Create(order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}
