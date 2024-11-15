//models/setup.go

package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Create a shorthand function to check for errors
func check(e error, m string) {
	if e != nil {
		log.Panic(m+": ", e.Error())
	}
}

func ConnectDatabase() {

	// Connect to the database
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	check(err, "Database connection error")

	// Migrate product definition to database
	err = database.AutoMigrate(&Product{})
	check(err, "Product migration error")

	// Migrate order definition to database
	err = database.AutoMigrate(&Order{})
	check(err, "Order migration error")

	// Set global database variable
	DB = database
}
