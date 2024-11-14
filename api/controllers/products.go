// controllers/products.go

package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/Acstrayer/TESCSE-Ecom/api/models"
"time"
)

type CreateProductInput struct {
	Name  string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Quantity uint `json:"quantity" binding:"required"`
	Price int `json:"price" binding:"required"`
}

// GET /products
// Get all products
func FindProducts(c *gin.Context) {
	prd := new(models.ProductRequestData)
	models.DB.Find(&prd.Products)
	prd.Retrieved = time.Now()
	c.JSON(http.StatusOK, prd)
}

func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	// Create product
	product := models.Product{Name: input.Name, Type: input.Type, Quantity: input.Quantity, Price: input.Price}
	models.DB.Create(&product)
  
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func FindProduct(c *gin.Context) {

}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}