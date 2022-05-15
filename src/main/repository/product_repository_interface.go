package repository

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/entity"
)

type ProductRepositoryInterface interface {
	FindById(id string) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
	Create([]entity.Product) ([]entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
}
