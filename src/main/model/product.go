package model

import "gorm.io/gorm"

//Product contains : name, description, price, and quantity
type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
}
