package repository

import "github.com/kartikeya/product_catalog_DIY/entity"

type ProductRepository interface {
	GetRecordById(id string) (entity.Product, error)
	GetAllRecords() ([]entity.Product, error)
	AddRecords([]entity.Product) ([]entity.Product, error)
	UpdateRecord(product entity.Product) (*entity.Product, error)
}
