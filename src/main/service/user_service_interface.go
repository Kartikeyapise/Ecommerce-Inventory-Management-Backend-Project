package service

import "github.com/kartikeya/product_catalog_DIY/src/main/model"

type UserServiceInterface interface {
	CreateUser(user model.User) (*model.User, error)
	IsUserValid(email string) (bool, error)
	IsUserTypeValid(user model.User) bool
	IsNormalUser(user model.User) bool
	IsMerchant(user model.User) bool
	IsMerchantEmail(email string) (bool, error)
}
