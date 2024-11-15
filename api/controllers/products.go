// controllers/products.go

package controllers

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	PLU         int    `json:"plu" bindig:"required"`
}

// Get all products
func FindProducts(c *gin.Context) {
	prd := new(models.ProductRequestData)
	if c.Query("name") != "" {
		models.DB.Where("name = ?", c.Query("name")).Find(&prd.Products)
	} else {
		models.DB.Find(&prd.Products)
	}
	prd.Retrieved = time.Now()
	c.JSON(http.StatusOK, prd)
}

func CreateProduct(c *gin.Context) {
	// Validate input
	input := new(CreateProductInput)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create product
	product := models.Product{Name: input.Name, Type: input.Type, Description: input.Description,
		Image: input.Image, Quantity: input.Quantity, Price: input.Price,
		PLU: input.PLU}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func FindProduct(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
