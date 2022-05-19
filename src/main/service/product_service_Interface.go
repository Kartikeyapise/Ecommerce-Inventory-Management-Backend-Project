package service

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
)

type ProductServiceInterface interface {
	AddProducts(userEmail string, products []model.Product) error
	GetProductById(s string) (*model.Product, error)
	GetProducts() ([]model.Product, error)
	BuyProduct(purchaseInfo model.Sales) (*model.Product, error)
	GetRecommendedProducts(topNProducts string) ([]model.Product, error)
}
