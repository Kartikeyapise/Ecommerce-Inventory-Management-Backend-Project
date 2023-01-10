package repository

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"gorm.io/gorm"
	"strconv"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r ProductRepository) FindById(id string) (*model.Product, error) {
	var product model.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

func (r ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r ProductRepository) Create(products []model.Product) ([]model.Product, error) {
	for i := 0; i < len(products); i++ {
		var p model.Product
		err := r.DB.Where("name = ?", products[i].Name).First(&p).Error
		if err != nil {
			err = r.DB.Create(&products[i]).Error
			if err != nil {
				return nil, err
			}
		} else {
			p.Quantity = addQuantity(p.Quantity, products[i].Quantity)
			products[i].Quantity = p.Quantity
			err := r.DB.Save(&p).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return products, nil
}

func (r ProductRepository) Update(product *model.Product) (*model.Product, error) {
	err := r.DB.Save(&product).Error
	return product, err
}

func addQuantity(q1, q2 string) string {
	q1int, _ := strconv.Atoi(q1)
	q2int, _ := strconv.Atoi(q2)
	return strconv.Itoa(q1int + q2int)
}
