package controller

import "net/http"

type ProductControllerInterface interface {
	AddProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	BuyProduct(w http.ResponseWriter, r *http.Request)
	GetRecommendedProducts(w http.ResponseWriter, r *http.Request)
}
