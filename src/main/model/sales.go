package model

import "gorm.io/gorm"

//Sales contains : user email and the ProductId, quantity which the user has bought.
type Sales struct {
	gorm.Model
	User      User    `gorm:"foreignkey:UserEmail"`
	Product   Product `gorm:"foreignkey:ProductId"`
	UserEmail string  `json:"user_email"`
	ProductId string  `json:"product_id"`
	Quantity  string  `json:"quantity"`
}
