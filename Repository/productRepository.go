package Repository

import (
	"github.com/kartikeya/product_catalog_DIY/Model"
	"gorm.io/gorm"
	"strconv"
)

func GetProductById(DB *gorm.DB, id string) Model.Product {
	var product Model.Product
	DB.First(&product, id)
	return product
}

func GetProducts(DB *gorm.DB) []Model.Product {
	var products []Model.Product
	DB.Find(&products)
	return products
}

func AddProducts(DB *gorm.DB, products []Model.Product) []Model.Product {
	DB.Create(products)
	return products
}

func UpdateProductQuantity(product Model.Product, quantity int, DB *gorm.DB) {
	product.Quantity = strconv.Itoa(quantity)
	DB.Save(&product)
}
