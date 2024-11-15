// models/order.go

package models

type OrderItem struct {
	ID        int `json:"id" gorm:"primary_key"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"` //in pennies
	ProductID int `json:"product" gorm:"foreignKey:ID"`
	OrderID   int `json:"order_id"`
}

// Method for getting subtotal
func (s OrderItem) SubTotal() int {
	return s.Price * s.Quantity
}

type Order struct {
	ID    uint        `json:"id" gorm:"primary_key"`
	Items []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}
