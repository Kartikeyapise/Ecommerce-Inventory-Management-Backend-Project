package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kartikeya/product_catalog_DIY/src/main/entity"
	"github.com/kartikeya/product_catalog_DIY/src/main/service"
	"github.com/kartikeya/product_catalog_DIY/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestGetProductById(t *testing.T) {
	//mockRepo := new(mocks.MockRepository)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	//setup expectations
	expectedProduct := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	}

	mockProductRepository.EXPECT().FindById("1").Return(expectedProduct, nil).Times(1)

	product, err := productService.GetProductById("1")

	//data Assertion
	assert.Nil(t, err)
	assert.Equal(t, uint(1), product.ID)
	assert.Equal(t, expectedProduct.Name, product.Name)
	assert.Equal(t, expectedProduct.Description, product.Description)
	assert.Equal(t, expectedProduct.Price, product.Price)
	assert.Equal(t, expectedProduct.Quantity, product.Quantity)

}

func TestGetProductByIdWhenIdNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	mockProductRepository.EXPECT().FindById("1").Return(nil, errors.New("error")).Times(1)

	product, err := productService.GetProductById("1")

	//data Assertion
	assert.Nil(t, product)
	assert.NotNil(t, err)

}

func TestAddProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	//setup expectations
	expectedProducts := []entity.Product{entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}

	mockProductRepository.EXPECT().Create(gomock.Any()).Return(expectedProducts, nil).Times(1)

	products, err := productService.AddProducts(expectedProducts)

	//data Assertion
	assert.Nil(t, err)
	assert.Equal(t, expectedProducts[0].ID, products[0].ID)
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)
	assert.Equal(t, expectedProducts[0].Description, products[0].Description)
	assert.Equal(t, expectedProducts[0].Price, products[0].Price)
	assert.Equal(t, expectedProducts[0].Quantity, products[0].Quantity)
}

func TestBuyProductWhenQuantityNotAvailable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	//setup expectations
	expectedProduct := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "100",
	}

	mockProductRepository.EXPECT().FindById("1").Return(expectedProduct, nil).Times(1)

	result, err := productService.BuyProduct("1", "1000")

	//data Assertion
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "Max Quantity exceeded", err.Error())
}

func TestBuyProductWhenIdNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	mockProductRepository.EXPECT().FindById("1").Return(nil, errors.New("record not found")).Times(1)

	result, err := productService.BuyProduct("1", "1000")

	//data Assertion
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
}

func TestBuyProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	//setup expectations
	expectedProduct := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "100",
	}

	mockProductRepository.EXPECT().FindById("1").Return(expectedProduct, nil).Times(1)
	mockProductRepository.EXPECT().Update(gomock.Any()).Return(expectedProduct, nil).Times(1)
	product, err := productService.BuyProduct("1", "10")

	//data Assertion
	assert.Equal(t, expectedProduct.ID, product.ID)
	assert.Equal(t, expectedProduct.Name, product.Name)
	assert.Equal(t, expectedProduct.Description, product.Description)
	assert.Equal(t, expectedProduct.Price, product.Price)
	assert.Equal(t, expectedProduct.Quantity, product.Quantity)
	assert.Nil(t, err)
}

func TestGetTop5Products(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	//setup expectations
	expectedProducts := []entity.Product{
		entity.Product{
			Model:       gorm.Model{ID: 1, UpdatedAt: time.Now()},
			Name:        "N",
			Description: "D",
			Price:       "P",
			Quantity:    "Q",
		},
	}

	mockProductRepository.EXPECT().FindAll().Return(expectedProducts, nil).Times(1)

	products, err := productService.GetTop5Products()

	//data Assertion
	assert.Nil(t, err)
	assert.Equal(t, expectedProducts[0].ID, products[0].ID)
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)
	assert.Equal(t, expectedProducts[0].Description, products[0].Description)
	assert.Equal(t, expectedProducts[0].Price, products[0].Price)
	assert.Equal(t, expectedProducts[0].Quantity, products[0].Quantity)
}

func TestGetTop5ProductsWhenRepoThrowsAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	mockProductRepository.EXPECT().FindAll().Return(nil, errors.New("something went wrong")).Times(1)

	products, err := productService.GetTop5Products()

	//data Assertion
	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "something went wrong", err.Error())
}
