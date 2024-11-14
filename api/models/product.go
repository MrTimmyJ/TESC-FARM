// models/product.go

package models

import (
	"time"
)

type ProductRequestData struct {
	Products  []Product
	Retrieved time.Time
}

type Product struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity uint   `json:"quantity`
	Price    int    `json: price` //In pennies
}

type CreateProductInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
