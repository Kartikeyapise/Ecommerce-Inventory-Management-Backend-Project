package service

import (
	"errors"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"github.com/kartikeya/product_catalog_DIY/src/main/repository"
	"strings"
)

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func (u UserService) CreateUser(user model.User) (*model.User, error) {
	if !u.IsUserTypeValid(user) {
		return nil, errors.New("cannot Create User. Invalid User Type")
	}
	return u.UserRepository.Create(user)
}

func (u UserService) IsUserTypeValid(user model.User) bool {
	for _, value := range model.ValidUserTypes {
		if strings.Compare(value, user.Type) == 0 {
			return true
		}
	}
	return false
}

func (u UserService) IsMerchant(user model.User) bool {
	if user.Type == "MERCHANT" {
		return true
	}
	return false
}

func (u UserService) IsMerchantEmail(email string) (bool, error) {
	user, err := u.UserRepository.FindUserByEmail(email)
	if err != nil {
		return false, err
	}
	if u.IsMerchant(*user) {
		return true, nil
	}
	return false, errors.New("this is not a Merchant Email")
}

func (u UserService) IsUserValid(email string) (bool, error) {
	_, err := u.UserRepository.FindUserByEmail(email)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u UserService) IsNormalUser(user model.User) bool {
	if user.Type == "USER" {
		return true
	}
	return false
}
