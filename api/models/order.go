// models/order.go

package models

import (
    "time"
	"gorm.io/gorm"
)

type OrderItem struct {
    gorm.Model
    Quantity  int `json:"quantity"`
    Price     int `json:"price"` //in pennies
    Product   Product `json:"product"`
    ProductID uint `json:"product_id"`
    OrderID   uint `json:"order_id"`
}

type OrderRequestData struct {
    Orders []Order
    Retrieved time.Time
}

type Order struct {
    gorm.Model
    Items    []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Address1 string      `json:"address_one"`
	Address2 string      `json:"address_two"`
	City     string      `json:"city"`
	State    string      `json:"state"`
	Zip      string      `json:"zip"`
}
