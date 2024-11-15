package main

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/controllers"
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// For Jess; from Austin: excuse me?
func poop(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("💩"))
}

func main() {
	// Initialize the router and static routes
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/", poop)
	r.GET("/poop", poop)

	//Product endpoints
	r.GET("/products", controllers.FindProducts) //All products
	r.POST("/products/new", controllers.CreateProduct)
	r.GET("/product/:id", controllers.FindProduct)
	r.PATCH("/product/:id/update", controllers.UpdateProduct)
	r.DELETE("/product/:id/delete", controllers.DeleteProduct)

	//Order endpoints
	r.POST("/orders/new", controllers.CreateOrder)

	err := r.Run("127.0.0.1:8080")
	models.Check(err, "Error starting API")
}
