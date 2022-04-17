package Repository1

import (
	"github.com/kartikeya/product_catalog_DIY/Model1"
	"gorm.io/gorm"
	"strconv"
)

func GetProductById(DB *gorm.DB, id string) entity.Product {
	var product entity.Product
	DB.First(&product, id)
	return product
}

func GetProducts(DB *gorm.DB) []entity.Product {
	var products []entity.Product
	DB.Find(&products)
	return products
}

func AddProducts(DB *gorm.DB, products []entity.Product) []entity.Product {
	DB.Create(products)
	return products
}

func UpdateProductQuantity(product entity.Product, quantity int, DB *gorm.DB) {
	product.Quantity = strconv.Itoa(quantity)
	DB.Save(&product)
}
