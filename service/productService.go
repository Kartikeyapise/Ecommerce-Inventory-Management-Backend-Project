package service

import (
	"errors"
	"github.com/kartikeya/product_catalog_DIY/entity"
	"github.com/kartikeya/product_catalog_DIY/repository"
	"sort"
	"strconv"
	"time"
)

type service struct{}

var (
	productRepository repository.ProductRepository
)

func NewProductService(repo repository.ProductRepository) ProductService {
	productRepository = repo
	return &service{}
}

func (s service) AddProducts(products []entity.Product) error {
	_, err := productRepository.AddRecords(products)
	return err
}

func (s service) GetProductById(id string) (entity.Product, error) {
	return productRepository.GetRecordById(id)
}

func (s service) GetProducts() ([]entity.Product, error) {
	return productRepository.GetAllRecords()
}

func (s service) BuyProduct(id string, quantity string) (*entity.Product, error) {
	product, err := productRepository.GetRecordById(id)
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
	return productRepository.UpdateRecord(product)
}

func (s service) GetTop5Products() ([]entity.Product, error) {
	products, err := productRepository.GetAllRecords()
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
