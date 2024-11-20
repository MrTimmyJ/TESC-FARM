// models/usermodel.go

package models

type UserDB struct {
	ID          		uint   `json:"ID" gorm:"primary_key"`
	users_name  		string `json:"users_name"`
	permissions_level 	uint   `json:"permissions_level"`
	password_hash		string `json:"password_hash"`
	email				string `json:"email"`
}

type UserFE struct {
	users_name  		string `json:"users_name" binding:"required"`
	password			string `json:"password" binding: "required"`
	email				string `json:"email" binding: "required"`
}