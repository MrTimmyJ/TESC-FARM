// controllers/ordercontroller.go

//Will create possible token creation and validation

package controllers

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//Hashes and salts password for DB storage
func HashPassword(pass string) (string, error) {
	bytepass := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(bytepass, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

//User creation endpoint
func CreateUser(c *gin.Context) {
	input := new(models.UserFE)
	//checks value extraction from json
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed_password, err := HashPassword(input.Password)
	//checks whether hash and salt failed
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}
	//creates DB model object
	user := models.UserDB{Users_name: input.Users_name,
		Permissions_level: 0,
		Password_hash:     hashed_password,
		Email:             input.Email}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
