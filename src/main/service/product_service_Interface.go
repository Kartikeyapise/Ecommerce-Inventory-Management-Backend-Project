package service

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/entity"
)

type ProductServiceInterface interface {
	AddProducts(products []entity.Product) ([]entity.Product, error)
	GetProductById(s string) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	BuyProduct(id string, quantity string) (*entity.Product, error)
	GetTop5Products() ([]entity.Product, error)
}