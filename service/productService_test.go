package service

import (
	"github.com/kartikeya/product_catalog_DIY/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) FindById(id string) (*entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Product), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Product), args.Error(1)
}

func (mock *MockRepository) Create(products []entity.Product) ([]entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Product), args.Error(1)
}

func (mock *MockRepository) Update(product *entity.Product) (*entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Product), args.Error(1)
}

func (mock *MockRepository) GetProductById(id string) (*entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Product), args.Error(1)
}

func TestGetProductById(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectations
	product := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	}

	mockRepo.On("FindById").Return(product, nil).Once()

	testService := NewProductService(mockRepo)

	result, err := testService.GetProductById("1")
	//Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "N", result.Name)
	assert.Equal(t, "D", result.Description)
	assert.Equal(t, "P", result.Price)
	assert.Equal(t, "Q", result.Quantity)
	assert.Nil(t, err)
}

func TestAddProducts(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectations
	products := []entity.Product{entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}

	mockRepo.On("Create").Return(products, nil).Once()

	testService := NewProductService(mockRepo)

	result, err := testService.AddProducts(products)

	//Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.Equal(t, uint(1), result[0].ID)
	assert.Equal(t, "N", result[0].Name)
	assert.Equal(t, "D", result[0].Description)
	assert.Equal(t, "P", result[0].Price)
	assert.Equal(t, "Q", result[0].Quantity)
	assert.Nil(t, err)
}

func TestBuyProductWhenQuantityNotAvailable(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectations
	product := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "100",
	}

	mockRepo.On("FindById").Return(product, nil).Once()
	//mockRepo.On("Update").Return(product, nil).Once()

	testService := NewProductService(mockRepo)

	result, err := testService.BuyProduct("1", "1000")

	//Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestBuyProduct(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectations
	product := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "100",
	}

	mockRepo.On("FindById").Return(product, nil).Once()
	mockRepo.On("Update").Return(product, nil).Once()

	testService := NewProductService(mockRepo)

	result, err := testService.BuyProduct("1", "10")

	//Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "N", result.Name)
	assert.Equal(t, "D", result.Description)
	assert.Equal(t, "P", result.Price)
	assert.Equal(t, "90", result.Quantity)
	assert.Nil(t, err)
}

func TestGetTop5Products(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectations
	products := []entity.Product{
		entity.Product{
			Model:       gorm.Model{ID: 1, UpdatedAt: time.Now()},
			Name:        "N",
			Description: "D",
			Price:       "P",
			Quantity:    "Q",
		},
	}

	mockRepo.On("FindAll").Return(products, nil).Once()

	testService := NewProductService(mockRepo)

	result, err := testService.GetTop5Products()

	//Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.Equal(t, uint(1), result[0].ID)
	assert.Equal(t, "N", result[0].Name)
	assert.Equal(t, "D", result[0].Description)
	assert.Equal(t, "P", result[0].Price)
	assert.Equal(t, "Q", result[0].Quantity)
	assert.Nil(t, err)
}
