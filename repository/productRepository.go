package repository

import (
	"github.com/kartikeya/product_catalog_DIY/entity"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewProductRepository(database *gorm.DB) ProductRepository {
	return &repository{DB: database}
}

func (r repository) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

func (r repository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r repository) Create(products []entity.Product) ([]entity.Product, error) {
	err := r.DB.Create(products).Error
	return products, err
}

func (r repository) Update(product *entity.Product) (*entity.Product, error) {
	err := r.DB.Save(&product).Error
	return product, err
}
