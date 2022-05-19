package service

import (
	"errors"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"github.com/kartikeya/product_catalog_DIY/src/main/repository"
	"sort"
	"strconv"
	"time"
)

type ProductService struct {
	UserService       UserServiceInterface
	ProductRepository repository.ProductRepositoryInterface
	SalesRepository   repository.SalesRepositoryInterface
}

func (s ProductService) AddProducts(userEmail string, products []model.Product) error {
	flag, err := s.UserService.IsMerchantEmail(userEmail)
	if err != nil {
		return err
	}
	if flag {
		_, err := s.ProductRepository.Create(products)
		return err
	}
	return errors.New("only merchant can add Products")
}

func (s ProductService) GetProductById(id string) (*model.Product, error) {
	return s.ProductRepository.FindById(id)
}

func (s ProductService) GetProducts() ([]model.Product, error) {
	return s.ProductRepository.FindAll()
}

func (s ProductService) BuyProduct(purchaseInfo model.Sales) (*model.Product, error) {
	_, err := s.UserService.IsUserValid(purchaseInfo.UserEmail)
	if err != nil {
		return nil, err
	}
	product, err := s.ProductRepository.FindById(purchaseInfo.ProductId)
	if err != nil {
		return nil, err
	}
	numberOfProductsAvailable, _ := strconv.Atoi(product.Quantity)
	numberOfProductsRequired, _ := strconv.Atoi(purchaseInfo.Quantity)
	if numberOfProductsAvailable < numberOfProductsRequired {
		return nil, errors.New("max Quantity exceeded")
	}

	_, err1 := s.SalesRepository.Create(purchaseInfo)

	if err1 != nil {
		return nil, err
	}
	product.Quantity = strconv.Itoa(numberOfProductsAvailable - numberOfProductsRequired)
	return s.ProductRepository.Update(product)
}

func (s ProductService) GetRecommendedProducts(topNProducts string) ([]model.Product, error) {
	n, err := strconv.Atoi(topNProducts)
	if err != nil {
		n = 5
	}

	products, err := s.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	sort.Slice(products, func(i, j int) bool {
		return products[i].UpdatedAt.After(products[j].UpdatedAt)
	})
	i := 0
	for _, p := range products {
		if p.UpdatedAt.After(time.Now().Add(-1*time.Hour)) && i < n {
			i++
		} else {
			break
		}
	}
	return products[0:i], nil
}
