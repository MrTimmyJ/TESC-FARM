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
type Product struct {
	gorm.Model
	Name     string
	Quantity int // In inventory
	Price    int // In cents
}

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

// For Jess
func poop(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("💩"))
}

// MAIN FUNCTION
func main() {
    // Connect to database
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
    if err != nil {
        panic("Database Connection Error: " + err.Error())
    }
    // Migrate ORM
    db.AutoMigrate(&Product{})

	// Instantiate the router
	r := gin.Default()
	// Define root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})
    // Define poop endpoint
	r.GET("/poop", poop)
	// Define produce endpoint
	r.GET("/produce.json", getProduce)
	// Run router on localhost, port 8080
	r.Run("127.0.0.1:8080")
}
