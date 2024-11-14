// models/order.go

package models

import(
	"time"
)

type OrderItem struct {
	ID int
	Quantity int
	Price int //in pennies
}

type Order struct {
	ID uint   `json:"id" gorm:"primary_key"`
	Items OrderItem[]
	Retrieved time.Time
}