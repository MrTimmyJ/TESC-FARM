// controllers/orders.go

package controllers

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateOrderInput struct {
	Items    []models.OrderItem `json:"items" binding:"required"`
	Name     string             `json:"name" binding:"required"`
	Email    string             `json:"email" binding:"required"`
	Address1 string             `json:"address_one" binding:"required"`
	Address2 string             `json:"address_two" binding:"required"`
	City     string             `json:"city" binding:"required"`
	State    string             `json:"state" binding:"required"`
	Zip      string             `json:"zip" binding:"required"`
}

func GetOrders(c *gin.Context) {
	ord := new(models.OrderRequestData)
	if c.Query("orderID") != "" {
		models.DB.Where("orderID = ?", c.Query("OrderID")).Find(&ord.Orders)
	} else {
		models.DB.Preload("Items").Preload("Items.Product").Find(&ord.Orders)
	}
  
	ord.Retrieved = time.Now()
	c.JSON(http.StatusOK, ord)
}

func SearchOrders(c *gin.Context) {
	ord := new(models.OrderRequestData)
	query := c.Query("query")

	//Magic gorm search
	if err := models.DB.Preload("Items").Preload("Items.Product").Where("id LIKE ? OR name LIKE ? OR email LIKE ? OR address1 LIKE ? OR address2 LIKE ? OR city LIKE ? OR state LIKE ? OR zip LIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&ord.Orders).Error; err != nil {
		//Database error, return
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//Query completed succesfully, return json
	ord.Retrieved = time.Now()
	c.JSON(http.StatusOK, ord)
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
