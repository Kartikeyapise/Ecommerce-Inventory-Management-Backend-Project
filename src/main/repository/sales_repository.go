package repository

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"gorm.io/gorm"
)

type SalesRepository struct {
	DB *gorm.DB
}

func (s SalesRepository) Create(purchaseInfo model.Sales) (*model.Sales, error) {
	err := s.DB.Create(&purchaseInfo).Error
	if err != nil {
		return nil, err
	}
	return &purchaseInfo, nil
}
