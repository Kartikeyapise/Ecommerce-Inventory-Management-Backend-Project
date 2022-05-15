package service

import (
	"errors"
	"github.com/kartikeya/product_catalog_DIY/src/main/entity"
	"github.com/kartikeya/product_catalog_DIY/src/main/repository"
	"sort"
	"strconv"
	"time"
)

type Service struct {
	ProductRepository repository.ProductRepositoryInterface
}

func (s Service) AddProducts(products []entity.Product) ([]entity.Product, error) {
	return s.ProductRepository.Create(products)
}

func (s Service) GetProductById(id string) (*entity.Product, error) {
	return s.ProductRepository.FindById(id)
}

func (s Service) GetProducts() ([]entity.Product, error) {
	return s.ProductRepository.FindAll()
}

func (s Service) BuyProduct(id string, quantity string) (*entity.Product, error) {
	product, err := s.ProductRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	numberOfProductsAvailable, _ := strconv.Atoi(product.Quantity)
	numberOfProductsRequired, _ := strconv.Atoi(quantity)
	if numberOfProductsAvailable < numberOfProductsRequired {
		//return nil, "Max Quantity available is "+ strconv.Itoa(numberOfProductsAvailable)
		return nil, errors.New("Max Quantity exceeded")
	}
	product.Quantity = strconv.Itoa(numberOfProductsAvailable - numberOfProductsRequired)
	return s.ProductRepository.Update(product)
}

func (s Service) GetTop5Products() ([]entity.Product, error) {
	products, err := s.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	sort.Slice(products, func(i, j int) bool {
		return products[i].UpdatedAt.After(products[j].UpdatedAt)
	})
	i := 0
	for _, p := range products {
		if p.UpdatedAt.After(time.Now().Add(-1*time.Hour)) && i < 5 {
			i++
		} else {
			break
		}
	}
	return products[0:i], nil
}
