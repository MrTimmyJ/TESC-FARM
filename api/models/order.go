// models/order.go

package models

import (
	"time"
)

type OrderItem struct {
	ID       int
	Quantity int
	Price    int //in pennies
}

// Method for getting subtotal
func (s OrderItem) SubTotal() int {
	return s.Price * s.Quantity
}

type Order struct {
	ID    uint        `json:"id" gorm:"primary_key"`
	Items []OrderItem `json:"items"`
}
