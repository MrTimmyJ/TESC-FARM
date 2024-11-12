package main

import (
	"github.com/gin-gonic/gin"
      "database/sql"
	"net/http"
	"time"
)

// GLOBAL VARIABLES (>^^)>
var db *sql.DB
var products []Product

type Product struct {
	Name     string
	Quantity int
	Price    int
}

type ProduceData struct {
	Retrieved time.Time
	Products  []Product
}

func getProduce(c *gin.Context) {
	data := ProduceData{}
      rows, err := db.QueryRows("SELECT name, quantity FROM products;")
      if err != nil {
            panic(err.Error())
      }
      for rows.Next() {
            p := Product{}
            err := rows.Scan(&p.Name, &p.Quantity)
            if err != nil {
                  panic(err.Error())
            }
            data.Products = append(data.Products, p)
      }
	data.Retrieved = time.Now()
	c.JSON(http.StatusOK, data)
}

func heart(c *gin.Context) {
      c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("❤️ "))
}

func poop(c *gin.Context) {
      c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("💩"))
}

func main() {
      products = []Product{
            Product{Name: "Apple", Quantity: 10},
            Product{Name: "Peach", Quantity: 5},
            Product{Name: "Tomato", Quantity: 8},
            Product{Name: "🍇", Quantity: 12},
      }
      var err error
      db, err = sql.Open("sqlite3", "data.db")

	// Instantiate the router
      r := gin.Default()
	// Define root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})
      r.GET("/heart", heart)
      r.GET("/poop", poop)
	// Define produce endpoint
	r.GET("/produce.json", getProduce)
	// Run router on port 8080
	r.Run("127.0.0.1:8080")
}
