// controllers/ordercontroller.go

//Will create possible token creation and validation

package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Hashes and salts password for DB storage
func HashPassword(string pass) {
	bytepass := byte[](pass)
	hash, err := bcrypt.GenerateFromPassword(bytepass, bcrypt.MinCost)
	if err != nil {
		log.Println(err);
		return nil, err
	}
	return string(hash)
}

//User creation endpoint
func CreateUser(c *gin.Context) {
	input := new(models.UserFE)
	//checks value extraction from json
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed_password, err := HashPassword(input.password)
	//checks whether hash and salt failed
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error:", err.Error()})
		return
	}
	//creates DB model object
	user := models.UserDB{ID: nil, users_name: input.users_name,
							permissions_level: 0, 
							password_hash: hashed_password,
							email: input.email,}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin,H{"data": user})
}
