// models/order.go

package models

type OrderItem struct {
	ID int
	Quantity int
	Price int //in pennies
}

type Order struct {
	ID uint   `json:"id" gorm:"primary_key"`
	Items []OrderItem `json:"items"`
      Name string `json:"name"`
      Email string `json:"email"`
      Address1 string `json:"address_one"`
      Address2 string `json:"address_two"`
      City string `json:"city"`
      State string `json:"state"`
      Zip string `json:"zip"`
}
