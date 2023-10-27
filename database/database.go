package database

import (
	"fmt"
	"itmx-test/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error

	// Open a SQLite database connection
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Auto-migrate the model
	err = DB.AutoMigrate(&model.Customers{})
	if err != nil {
		panic("Failed to auto-migrate the model")
	}

	// Create sample customer records
	DB.Create(&model.Customers{Name: "James", Age: 21})
	DB.Create(&model.Customers{Name: "Ben", Age: 20})
	DB.Create(&model.Customers{Name: "Jack", Age: 22})

	fmt.Println("Database initialization complete!")
}
