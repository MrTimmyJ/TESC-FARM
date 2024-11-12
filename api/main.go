package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Product struct {
	Name     string
	Quantity int
}

type ProduceData struct {
	Retrieved time.Time
	Products  []Product
}

func getProduce(c *gin.Context) {
	data := ProduceData{}
    data.Retrieved = time.Now()
	c.JSON(http.StatusOK, data)
}

func main() {
    // Instantiate the router
    r := gin.Default()
    // Define root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})
    // Define produce endpoint
	r.GET("/produce.json", getProduce)
    // Run router on port 8080
	r.Run("127.0.0.1:8080")
}
