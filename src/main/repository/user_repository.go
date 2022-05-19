package repository

import (
	"errors"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) Create(user model.User) (*model.User, error) {
	err := u.DB.Create(&user).Error
	if err != nil {
		return nil, errors.New("user Already Exixts")
	}
	return &user, nil
}

func (u UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, errors.New("user Does Not Exist")
	}
	return &user, nil
}
