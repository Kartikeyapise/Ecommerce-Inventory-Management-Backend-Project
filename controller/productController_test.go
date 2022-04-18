package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/Database"
	"github.com/kartikeya/product_catalog_DIY/entity"
	"github.com/kartikeya/product_catalog_DIY/repository"
	"github.com/kartikeya/product_catalog_DIY/service"
	"github.com/kartikeya/product_catalog_DIY/view"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	DB                    *gorm.DB                     = Database.ConnectTestDatabase()
	productRepositoryTest repository.ProductRepository = repository.NewProductRepository(DB)
	productServiceTest    service.ProductService       = service.NewProductService(productRepositoryTest)
	productControllerTest ProductController            = NewProductController(productServiceTest)
)

func cleanDatabase() {
	DB.Exec("DELETE FROM products")
	//DB.Where("id IS NOT NULL").Delete(&entity.Product{})
}

func addSampleProduct() {
	DB.Create(&entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "1",
	})
}

func getSampleProduct() *entity.Product {
	return &entity.Product{
		Model:       gorm.Model{ID: 1},
		Name:        "N",
		Description: "D",
		Price:       "P",
		Quantity:    "1",
	}
}

func TestGetProductById(t *testing.T) {
	cleanDatabase()
	addSampleProduct()

	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/product/id", nil)
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.GetProductById)

	//Record Http response (httptest)
	response := httptest.NewRecorder()

	//dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertion on the HTTP Status code and the response
	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	//Decode the HTTP response
	var product entity.Product
	json.NewDecoder(response.Body).Decode(&product)

	//Assert HTTP response
	assert.NotNil(t, product)
	assert.Equal(t, uint(0x1), product.ID)
	assert.Equal(t, "N", product.Name)
	assert.Equal(t, "D", product.Description)
	assert.Equal(t, "P", product.Price)
	assert.Equal(t, "1", product.Quantity)
}

func TestGetProductByIdWhenIdNotAvailable(t *testing.T) {
	cleanDatabase()
	addSampleProduct()

	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/product/id", nil)
	vars := map[string]string{
		"id": "2",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.GetProductById)

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
	cleanDatabase()
	//addSampleProduct()

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
	handler := http.HandlerFunc(productControllerTest.AddProducts)

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
	cleanDatabase()
	addSampleProduct()

	//create request body
	req_body := []byte(`garbage`)

	//Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(req_body))

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.AddProducts)

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

func TestAddProductsWhenServiceReturnsAnError(t *testing.T) {
	cleanDatabase()
	addSampleProduct()
	//addSampleProduct()

	//create request body
	req_body := []byte(`[
		{	"id":1,
			"name":"n1",
			"description":"d1",
			"price":"p1",
			"quantity":"q1"
		}
	]`)

	//Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(req_body))

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.AddProducts)

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
	cleanDatabase()
	addSampleProduct()

	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/products", nil)

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.GetProducts)

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
	assert.Equal(t, uint(0x1), products[0].ID)
	assert.Equal(t, "N", products[0].Name)
	assert.Equal(t, "D", products[0].Description)
	assert.Equal(t, "P", products[0].Price)
	assert.Equal(t, "1", products[0].Quantity)
}

func TestBuyProduct(t *testing.T) {
	cleanDatabase()
	addSampleProduct()

	//Create a new HTTP GET request
	req, _ := http.NewRequest("PUT", "/buyProduct/1/1", nil)
	vars := map[string]string{
		"id":       "1",
		"quantity": "1",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.BuyProduct)

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
	cleanDatabase()
	addSampleProduct()

	//Create a new HTTP PUT request
	req, _ := http.NewRequest("PUT", "/buyProduct/1/20", nil)
	vars := map[string]string{
		"id":       "1",
		"quantity": "10",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.BuyProduct)

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
	cleanDatabase()
	addSampleProduct()

	//Create a new HTTP PUT request
	req, _ := http.NewRequest("PUT", "/buyProduct/10/1", nil)
	vars := map[string]string{
		"id":       "10",
		"quantity": "1",
	}
	req = mux.SetURLVars(req, vars)
	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.BuyProduct)

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
	cleanDatabase()
	addSampleProduct()

	//Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/products", nil)

	//Assign Http handles function (Add post function)
	handler := http.HandlerFunc(productControllerTest.GetTop5Products)

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
	assert.Equal(t, uint(0x1), products[0].ID)
	assert.Equal(t, "N", products[0].Name)
	assert.Equal(t, "D", products[0].Description)
	assert.Equal(t, "P", products[0].Price)
	assert.Equal(t, "1", products[0].Quantity)
}
