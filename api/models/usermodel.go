// models/usermodel.go

package models

type UserDB struct {
	ID          		uint   `json:"ID" gorm:"primary_key"`
	users_name  		string `json:"users_name"`
	permissions_level 	uint   `json:"permissions_level"`
	password_hash		string `json:"password_hash"`
	email				string `json:"email"`
}

type UserFrontEnd struct {
	ID          		uint   `json:"ID" gorm:"primary_key"`
	users_name  		string `json:"users_name"`
	permissions_level 	uint   `json:"permissions_level"`
	password			string `json:"password"`
	email				string `json:"email"`
}