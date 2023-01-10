package model

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestProduct(t *testing.T) {
	product := model.Product{
		Model:       gorm.Model{},
		Name:        "Name",
		Description: "Description",
		Price:       "Price",
		Quantity:    "Quantity",
	}

	assert.Equal(t, uint(0x0), product.ID)
	assert.Equal(t, "Name", product.Name)
	assert.Equal(t, "Description", product.Description)
	assert.Equal(t, "Price", product.Price)
	assert.Equal(t, "Quantity", product.Quantity)

}
