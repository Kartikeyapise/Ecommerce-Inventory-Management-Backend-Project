package repository

import "github.com/kartikeya/product_catalog_DIY/src/main/model"

type SalesRepositoryInterface interface {
	Create(purchaseInfo model.Sales) (*model.Sales, error)
}
