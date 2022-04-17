package Controller1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/Model1"
	"github.com/kartikeya/product_catalog_DIY/Repository1"
	"gorm.io/gorm"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func AddProducts(w http.ResponseWriter, r *http.Request, DB *gorm.DB) {
	w.Header().Set("Content-type", "application/json")
	var products []entity.Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Repository1.AddProducts(DB, products)
	json.NewEncoder(w).Encode(`{ status : products added successfully}`)
}

func GetProductById(w http.ResponseWriter, r *http.Request, DB *gorm.DB) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//fmt.Println(params, w, r)
	json.NewEncoder(w).Encode(Repository1.GetProductById(DB, params["id"]))
}

func GetProducts(w http.ResponseWriter, r *http.Request, DB *gorm.DB) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Repository1.GetProducts(DB))
}

func BuyProduct(w http.ResponseWriter, r *http.Request, DB *gorm.DB) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	product := Repository1.GetProductById(DB, params["id"])
	numberOfProductsAvailable, _ := strconv.Atoi(product.Quantity)
	numberOfProductsRequired, _ := strconv.Atoi(params["quantity"])
	if numberOfProductsAvailable < numberOfProductsRequired {
		json.NewEncoder(w).Encode("Max Quantity available is " + strconv.Itoa(numberOfProductsAvailable))
		return
	}
	Repository1.UpdateProductQuantity(product, numberOfProductsAvailable-numberOfProductsRequired, DB)
	json.NewEncoder(w).Encode("{status : Buy Successful}")
}

func GetTop5Products(w http.ResponseWriter, r *http.Request, DB *gorm.DB) {
	w.Header().Set("Content-type", "application/json")
	products := Repository1.GetProducts(DB)
	sort.Slice(products, func(i, j int) bool {
		return products[i].UpdatedAt.After(products[j].UpdatedAt)
	})
	i := 0
	for _, p := range products {
		if p.UpdatedAt.After(time.Now().Add(-1*time.Hour)) && i < 5 {
			i++
		} else {
			break
		}
	}
	json.NewEncoder(w).Encode(products[0:i])
}
