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

func produce(c *gin.Context) {
	data := ProduceData{}
	c.JSON(http.StatusOK, data)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})
	r.GET("/produce.json", produce)
	r.Run("127.0.0.1:8080")
}
