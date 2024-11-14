package main

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/controllers"
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// For Jess; from austin: excuse me?
func poop(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("💩"))
}

func main() {
	// ...
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/", poop)
	r.GET("/poop", poop)
	r.GET("/products", controllers.FindProducts) //All products
	r.POST("/products/new", controllers.CreateProduct)
	r.GET("/product/:id", controllers.FindProduct)
	r.PATCH("/product/:id/update", controllers.UpdateProduct)
	r.DELETE("/product/:id/delete", controllers.DeleteProduct)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic("Could not run the database.")
	}
}
