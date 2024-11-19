// models/order.go

package models

import (
	"time"
)

type OrderItem struct {
	ID        int      `json:"id" gorm:"primary_key"`
	Quantity  int      `json:"quantity"`
	Price     int      `json:"price"` //in pennies
	Product   *Product `json:"product"`
	ProductID uint     `json:"product_id"`
	Order     *Order   `json:"order"`
	OrderID   uint     `json:"order_id"`
}

// Method for getting subtotal
func (s OrderItem) SubTotal() int {
	return s.Price * s.Quantity
}

type Order struct {
	ID    uint        `json:"id" gorm:"primary_key"`
	Items []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
      Name string `json:"name"`
      Email string `json:"email"`
      Address1 string `json:"address_one"`
      Address2 string `json:"address_two"`
      City string `json:"city"`
      State string `json:"state"`
      Zip string `json:"zip"`
}

type OrderRequestData struct {
	Orders    []Order
	Retrieved time.Time
}
