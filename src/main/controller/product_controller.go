package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/src/main/entity"
	"github.com/kartikeya/product_catalog_DIY/src/main/errors"
	"github.com/kartikeya/product_catalog_DIY/src/main/service"
	"github.com/kartikeya/product_catalog_DIY/src/main/view"
	"net/http"
)

var (
	productService service.ProductService
)

type controller struct{}

func NewProductController(service service.ProductService) ProductController {
	productService = service
	return &controller{}
}

func (c controller) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	product, err := productService.GetProductById(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ProductError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (c controller) AddProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var products []entity.Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ProductError{Message: "Error extracting products from request body"})
		return
	}
	_, err = productService.AddProducts(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ProductError{Message: "Cannot add product. Something went wrong"})
		return
	}
	json.NewEncoder(w).Encode(view.ResponseMessage{"products added successfully"})
}

func (c controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	products, err := productService.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ProductError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (c controller) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	_, err := productService.BuyProduct(params["id"], params["quantity"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ProductError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(view.ResponseMessage{"Buy Successful"})
}

func (c controller) GetTop5Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	products, err := productService.GetTop5Products()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ProductError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(products)
}
