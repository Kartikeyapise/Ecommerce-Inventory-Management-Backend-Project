package repository

import "github.com/kartikeya/product_catalog_DIY/src/main/model"

type UserRepositoryInterface interface {
	Create(user model.User) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
}
