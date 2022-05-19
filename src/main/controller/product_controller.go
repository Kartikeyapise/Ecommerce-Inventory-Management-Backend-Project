package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"github.com/kartikeya/product_catalog_DIY/src/main/service"
	"github.com/kartikeya/product_catalog_DIY/src/main/view"
	"log"
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
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (c ProductController) AddProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var userProducts view.UserProducts
	err := json.NewDecoder(r.Body).Decode(&userProducts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: "Error extracting products from request body"})
		return
	}
	userEmail := userProducts.Email
	products := userProducts.Products
	err = c.ProductService.AddProducts(userEmail, products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(view.ResponseMessage{Message: "products added successfully"})
}

func (c ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	products, err := c.ProductService.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: "Cannot get Products.Something Went Wrong"})
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (c ProductController) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var purchaseInfo model.Sales
	err := json.NewDecoder(r.Body).Decode(&purchaseInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: "Error extracting purchase request body"})
		return
	}
	_, err1 := c.ProductService.BuyProduct(purchaseInfo)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: err1.Error()})
		return
	}
	json.NewEncoder(w).Encode(view.ResponseMessage{"Buy Successful"})
}

func (c ProductController) GetRecommendedProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	topNProducts := params["n"]
	log.Println("topNProducts::", topNProducts)
	w.Header().Set("Content-type", "application/json")
	products, err := c.ProductService.GetRecommendedProducts(topNProducts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(products)
}
