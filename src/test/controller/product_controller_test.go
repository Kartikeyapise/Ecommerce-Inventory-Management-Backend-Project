package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/src/main/controller"
	"github.com/kartikeya/product_catalog_DIY/src/main/entity"
	"github.com/kartikeya/product_catalog_DIY/src/main/view"
	"github.com/kartikeya/product_catalog_DIY/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductById(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	expectedProduct := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	}

	mockProductService.EXPECT().GetProductById("1").Return(expectedProduct, nil).Times(1)

	req, _ := http.NewRequest("GET", "/product/id", nil)
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	handler := http.HandlerFunc(productController.GetProductById)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	var product entity.Product
	json.NewDecoder(response.Body).Decode(&product)

	assert.NotNil(t, product)
	assert.Equal(t, expectedProduct.ID, product.ID)
	assert.Equal(t, expectedProduct.Name, product.Name)
	assert.Equal(t, expectedProduct.Description, product.Description)
	assert.Equal(t, expectedProduct.Price, product.Price)
	assert.Equal(t, expectedProduct.Quantity, product.Quantity)
}

func TestGetProductByIdWhenIdNotAvailable(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	mockProductService.EXPECT().GetProductById("2").Return(nil, errors.New("record not found")).Times(1)
	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/product/id", nil)
	vars := map[string]string{
		"id": "2",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.GetProductById)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.NotNil(t, message)
	assert.Equal(t, "record not found", message.Message)
}

func TestAddProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}
	products := []entity.Product{entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}
	mockProductService.EXPECT().AddProducts(gomock.Any()).Return(products, nil).Times(1)
	//create request body
	req_body := []byte(`[
		{
			"name":"n1",
			"description":"d1",
			"price":"p1",
			"quantity":"q1"
		},
		{
			"name":"n2",
			"description":"d2",
			"price":"p2",
			"quantity":"q2"
		}
	]`)

	//Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(req_body))

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.AddProducts)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.NotNil(t, message)
	assert.Equal(t, "products added successfully", message.Message)

}

func TestAddProductsWhenReqBodyIsNotSuppliedInCorrectFormat(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}
	//create request body
	req_body := []byte(`garbage`)

	//Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(req_body))

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.AddProducts)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.NotNil(t, message)
	assert.Equal(t, "Error extracting products from request body", message.Message)

}

func TestAddProductsWhenServiceReturnsAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	mockProductService.EXPECT().AddProducts(gomock.Any()).Return(nil, errors.New("error")).Times(1)
	//create request body
	reqBody := []byte(`[
		{	"id":1,
			"name":"n1",
			"description":"d1",
			"price":"p1",
			"quantity":"q1"
		}
	]`)

	//Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(reqBody))

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.AddProducts)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.NotNil(t, message)
	assert.Equal(t, "Cannot add product. Something went wrong", message.Message)

}

func TestGetProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}
	expectedProducts := []entity.Product{entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}
	mockProductService.EXPECT().GetProducts().Return(expectedProducts, nil).Times(1)

	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/products", nil)

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.GetProducts)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	//Decode the HTTP response
	var products []entity.Product
	json.NewDecoder(response.Body).Decode(&products)

	//Assert HTTP response
	assert.NotNil(t, products[0])
	assert.Equal(t, expectedProducts[0].ID, products[0].ID)
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)
	assert.Equal(t, expectedProducts[0].Description, products[0].Description)
	assert.Equal(t, expectedProducts[0].Price, products[0].Price)
	assert.Equal(t, expectedProducts[0].Quantity, products[0].Quantity)
}

func TestGetProductsWhenServiceReturnsAnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	mockProductService.EXPECT().GetProducts().Return(nil, errors.New("error")).Times(1)

	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/products", nil)

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.GetProducts)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.Equal(t, "Cannot get Products.Something Went Wrong", message.Message)

}

func TestBuyProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	expectedProduct := &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	}

	mockProductService.EXPECT().BuyProduct("1", "1").Return(expectedProduct, nil).Times(1)

	//Create a new HTTP GET request
	req, _ := http.NewRequest("PUT", "/buyProduct/1/1", nil)
	vars := map[string]string{
		"id":       "1",
		"quantity": "1",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.BuyProduct)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.NotNil(t, message)
	assert.Equal(t, "Buy Successful", message.Message)
}

func TestBuyProductWhenQuantityNotEnough(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	mockProductService.EXPECT().BuyProduct("1", "10").Return(nil, errors.New("Max Quantity exceeded")).Times(1)

	//Create a new HTTP PUT request
	req, _ := http.NewRequest("PUT", "/buyProduct/1/20", nil)
	vars := map[string]string{
		"id":       "1",
		"quantity": "10",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.BuyProduct)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.NotNil(t, message)
	assert.Equal(t, "Max Quantity exceeded", message.Message)
}

func TestBuyProductWhenCalledWithWrongId(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	mockProductService.EXPECT().BuyProduct("1", "10").Return(nil, errors.New("record not found")).Times(1)

	//Create a new HTTP PUT request
	req, _ := http.NewRequest("PUT", "/buyProduct/10/1", nil)
	vars := map[string]string{
		"id":       "1",
		"quantity": "10",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.BuyProduct)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	//Decode the HTTP response
	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	//Assert HTTP response
	assert.NotNil(t, message)
	assert.Equal(t, "record not found", message.Message)
}

func TestGetTop5Products(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductService := mocks.NewMockProductService(mockCtrl)
	productController := controller.Controller{
		ProductService: mockProductService,
	}

	expectedProducts := []entity.Product{entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "Q",
	},
	}
	mockProductService.EXPECT().GetTop5Products().Return(expectedProducts, nil).Times(1)

	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/products", nil)

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productController.GetTop5Products)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	//Decode the HTTP response
	var products []entity.Product
	json.NewDecoder(response.Body).Decode(&products)

	//Assert HTTP response
	assert.NotNil(t, products)
	assert.Equal(t, expectedProducts[0].ID, products[0].ID)
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)
	assert.Equal(t, expectedProducts[0].Description, products[0].Description)
	assert.Equal(t, expectedProducts[0].Price, products[0].Price)
	assert.Equal(t, expectedProducts[0].Quantity, products[0].Quantity)
}
