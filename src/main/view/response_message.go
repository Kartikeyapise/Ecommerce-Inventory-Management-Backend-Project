package view

import "github.com/kartikeya/product_catalog_DIY/src/main/model"

type ResponseMessage struct {
	Message string `json:"message"`
}

type UserProducts struct {
	Email    string          `json:"user_email"`
	Products []model.Product `json:"products"`
}

type PurchaseInfo struct {
	Email     string `json:"email"`
	ProductId string `json:"product_id"`
	Quantity  string `json:"quantity"`
}
