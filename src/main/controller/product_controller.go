package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/src/main/custum_errors"
	"github.com/kartikeya/product_catalog_DIY/src/main/entity"
	"github.com/kartikeya/product_catalog_DIY/src/main/service"
	"github.com/kartikeya/product_catalog_DIY/src/main/view"
	"net/http"
)

type ProductController struct {
	ProductService service.ProductServiceInterface
}

func (c ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	product, err := c.ProductService.GetProductById(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(custum_errors.ProductError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (c ProductController) AddProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var products []entity.Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(custum_errors.ProductError{Message: "Error extracting products from request body"})
		return
	}
	_, err = c.ProductService.AddProducts(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(custum_errors.ProductError{Message: "Cannot add product. Something went wrong"})
		return
	}
	json.NewEncoder(w).Encode(view.ResponseMessage{"products added successfully"})
}

func (c ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	products, err := c.ProductService.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(custum_errors.ProductError{Message: "Cannot get Products.Something Went Wrong"})
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (c ProductController) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	_, err := c.ProductService.BuyProduct(params["id"], params["quantity"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(custum_errors.ProductError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(view.ResponseMessage{"Buy Successful"})
}

func (c ProductController) GetTop5Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	products, err := c.ProductService.GetTop5Products()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(custum_errors.ProductError{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(products)
}
