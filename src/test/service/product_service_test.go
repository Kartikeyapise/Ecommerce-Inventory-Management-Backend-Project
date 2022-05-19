package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
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
	expectedProduct := &model.Product{
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
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
		UserService:       mockUserService,
	}
	//setup expectations
	products := []model.Product{model.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}

	mockUserService.EXPECT().IsMerchantEmail(gomock.Any()).Return(true, nil)
	mockProductRepository.EXPECT().Create(gomock.Any()).Return(products, nil).Times(1)

	err := productService.AddProducts("email", products)

	//data Assertion
	assert.Nil(t, err)
}

func TestAddProductsWhenRepoThrowsAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
		UserService:       mockUserService,
	}
	//setup expectations
	products := []model.Product{model.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}

	mockUserService.EXPECT().IsMerchantEmail(gomock.Any()).Return(true, nil)
	mockProductRepository.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error")).Times(1)

	err := productService.AddProducts("email", products)

	//data Assertion
	assert.Equal(t, "error", err.Error())

}

func TestAddProductsWhenUserIsNotMerchant(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
		UserService:       mockUserService,
	}
	//setup expectations
	products := []model.Product{model.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}

	mockUserService.EXPECT().IsMerchantEmail(gomock.Any()).Return(false, errors.New("error"))

	err := productService.AddProducts("email", products)

	//data Assertion
	assert.Equal(t, "error", err.Error())

}

func TestBuyProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	mockSalesRepository := mocks.NewMockSalesRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
		UserService:       mockUserService,
		SalesRepository:   mockSalesRepository,
	}

	//setup expectations
	expectedProduct := &model.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "100",
	}
	mockUserService.EXPECT().IsUserValid(gomock.Any()).Return(true, nil)
	mockProductRepository.EXPECT().FindById(gomock.Any()).Return(expectedProduct, nil).Times(1)
	mockProductRepository.EXPECT().Update(gomock.Any()).Return(expectedProduct, nil).Times(1)
	mockSalesRepository.EXPECT().Create(gomock.Any()).Return(nil, nil).Times(1)

	purchaseInfo := model.Sales{
		Model:     gorm.Model{},
		User:      model.User{},
		Product:   model.Product{},
		UserEmail: "UserEmail",
		ProductId: "1",
		Quantity:  "10",
	}

	result, err := productService.BuyProduct(purchaseInfo)

	//data Assertion
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProduct.Name, result.Name)
	assert.Equal(t, expectedProduct.Price, result.Price)
	assert.Equal(t, expectedProduct.Quantity, result.Quantity)
	assert.Equal(t, expectedProduct.Description, result.Description)
	assert.Equal(t, expectedProduct.ID, result.ID)

}

func TestBuyProductWhenIdNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	mockSalesRepository := mocks.NewMockSalesRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
		UserService:       mockUserService,
		SalesRepository:   mockSalesRepository,
	}

	//setup expectations
	mockUserService.EXPECT().IsUserValid(gomock.Any()).Return(true, nil)
	mockProductRepository.EXPECT().FindById(gomock.Any()).Return(nil, errors.New("error")).Times(1)

	purchaseInfo := model.Sales{
		Model:     gorm.Model{},
		User:      model.User{},
		Product:   model.Product{},
		UserEmail: "UserEmail",
		ProductId: "1",
		Quantity:  "10",
	}

	result, err := productService.BuyProduct(purchaseInfo)

	//data Assertion
	assert.Nil(t, result)
	assert.Equal(t, "error", err.Error())
}

func TestBuyProductWhenQuantityNotAvailable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	mockSalesRepository := mocks.NewMockSalesRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
		UserService:       mockUserService,
		SalesRepository:   mockSalesRepository,
	}

	//setup expectations
	expectedProduct := &model.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "0",
	}
	mockUserService.EXPECT().IsUserValid(gomock.Any()).Return(true, nil)
	mockProductRepository.EXPECT().FindById(gomock.Any()).Return(expectedProduct, nil).Times(1)

	purchaseInfo := model.Sales{
		Model:     gorm.Model{},
		User:      model.User{},
		Product:   model.Product{},
		UserEmail: "UserEmail",
		ProductId: "1",
		Quantity:  "10",
	}

	result, err := productService.BuyProduct(purchaseInfo)

	//data Assertion
	assert.Nil(t, result)
	assert.Equal(t, "max Quantity exceeded", err.Error())
}

func TestGetRecommendedProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	//setup expectations
	expectedProducts := []model.Product{
		model.Product{
			Model:       gorm.Model{ID: 1, UpdatedAt: time.Now()},
			Name:        "N",
			Description: "D",
			Price:       "P",
			Quantity:    "Q",
		},
	}

	mockProductRepository.EXPECT().FindAll().Return(expectedProducts, nil).Times(1)

	products, err := productService.GetRecommendedProducts("2")

	//data Assertion
	assert.Nil(t, err)
	assert.Equal(t, expectedProducts[0].ID, products[0].ID)
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)
	assert.Equal(t, expectedProducts[0].Description, products[0].Description)
	assert.Equal(t, expectedProducts[0].Price, products[0].Price)
	assert.Equal(t, expectedProducts[0].Quantity, products[0].Quantity)
}

func TestTestGetRecommendedProductsWhenRepoThrowsAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepositoryInterface(mockCtrl)
	productService := service.ProductService{
		ProductRepository: mockProductRepository,
	}

	mockProductRepository.EXPECT().FindAll().Return(nil, errors.New("something went wrong")).Times(1)

	products, err := productService.GetRecommendedProducts("1")

	//data Assertion
	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "something went wrong", err.Error())
}
