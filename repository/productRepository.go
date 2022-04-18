package repository

import (
	"github.com/kartikeya/product_catalog_DIY/entity"
	"gorm.io/gorm"
	"strconv"
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
	for i := 0; i < len(products); i++ {
		var p entity.Product
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

func (r repository) Update(product *entity.Product) (*entity.Product, error) {
	err := r.DB.Save(&product).Error
	return product, err
}

func addQuantity(q1, q2 string) string {
	q1int, _ := strconv.Atoi(q1)
	q2int, _ := strconv.Atoi(q2)
	return strconv.Itoa(q1int + q2int)
}
