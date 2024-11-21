// models/usermodel.go

package models

type UserDB struct {
	ID                uint   `json:"ID" gorm:"primary_key"`
	Users_name        string `json:"users_name"`
	Permissions_level uint   `json:"permissions_level"`
	Password_hash     string `json:"password_hash"`
	Email             string `json:"email"`
}

type UserFE struct {
	Users_name string `json:"users_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required"`
}
