// controllers/products.go

package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/rahmanfadhil/gin-bookstore/models"
)

// GET /products
// Get all products
func FindProducts(c *gin.Context) {
var products []models.Product
models.DB.Find(&products)

c.JSON(http.StatusOK, gin.H{"data": products})
}