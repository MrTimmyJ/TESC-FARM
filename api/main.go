package main

import (
      "github.com/gin-gonic/gin"
	  "github.com/Acstrayer/TESCSE-Ecom/api/models"
	  "github.com/Acstrayer/TESCSE-Ecom/api/controllers"
      "gorm.io/driver/sqlite"
      "gorm.io/gorm"
      "net/http"
      "time"
)

// For Jess; from austin: excuse me?
func poop(c *gin.Context) {
      c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("💩"))
}

func main() {
	// ...
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/poop", poop)
	r.GET("/products", controllers.FindProducts) //All products
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.FindProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct) 

	err := r.Run("127.0.0.1:8080")
	if err != nil {
			panic("Could not run the database.")
	}
}
