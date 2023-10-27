package model

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	//Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
