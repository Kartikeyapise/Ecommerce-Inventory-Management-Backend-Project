package controller

import "net/http"

type ProductController interface {
	AddProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	BuyProduct(w http.ResponseWriter, r *http.Request)
	GetTop5Products(w http.ResponseWriter, r *http.Request)
}
