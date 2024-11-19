// controllers/ordercontroller.go
//INCOMPLETE: working on more of it later.
//Will create a put endpoint for users, and possible token
//creation and validation

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
}

func CreateUserOrders(c *gin.Context) {

}
