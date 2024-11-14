// controllers/products.go
r.GET("/products/:id", controllers.FindProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct) 
package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/Acstrayer/TESCSE-Ecom/tree/main/api/models"
)

// GET /products
// Get all products
func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProduct(c *gin.Context) {

}

func FindProduct(c *gin.Context) {

}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {
	
}