package main

import (
      "github.com/gin-gonic/gin"
      "gorm.io/driver/sqlite"
      "gorm.io/gorm"
      "net/http"
      "time"
)

// GLOBAL VARIABLES
var db *gorm.DB

// STRUCTS & METHODS


type ProduceData struct {
      Retrieved time.Time
      Products  []Product
}

// ENDPOINT FUNCTIONS
func getProduce(c *gin.Context) {
      // Instantiate data
      data := ProduceData{}
      // Get data from DB
      db.Find(&data.Products)
      // Set retrieval time
      data.Retrieved = time.Now()
      // Return JSON
      c.JSON(http.StatusOK, data)
}

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
	r.GET("/produce", controllers.FindProduce) //Products, type: produce
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.FindProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct) 

	err := r.Run("127.0.0.1:8080")
	if err != nil {
			panic("Could not run the database.")
	}
}
