package repository

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
)

type ProductRepositoryInterface interface {
	FindById(id string) (*model.Product, error)
	FindAll() ([]model.Product, error)
	Create([]model.Product) ([]model.Product, error)
	Update(product *model.Product) (*model.Product, error)
}
