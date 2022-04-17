package repository

import (
	"github.com/kartikeya/product_catalog_DIY/Database"
	"github.com/kartikeya/product_catalog_DIY/entity"
)

type repository struct{}

func NewProductRepository() ProductRepository {
	return &repository{}
}

func (r repository) GetRecordById(id string) (entity.Product, error) {
	var product entity.Product
	err := Database.DB.First(&product, id).Error
	return product, err
}

func (r repository) GetAllRecords() ([]entity.Product, error) {
	var products []entity.Product
	err := Database.DB.Find(&products).Error
	return products, err
}

func (r repository) AddRecords(products []entity.Product) ([]entity.Product, error) {
	err := Database.DB.Create(products).Error
	return products, err
}

func (r repository) UpdateRecord(product entity.Product) (*entity.Product, error) {
	err := Database.DB.Save(&product).Error
	return &product, err
}
