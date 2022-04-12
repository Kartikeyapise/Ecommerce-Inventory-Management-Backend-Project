package API

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/Database"
	"github.com/kartikeya/product_catalog_DIY/Model"
	"net/http"
	"sort"
	"strconv"
	"time"
)

/*

payload =>>>>>>.

[
	{
		"name":"iphone12",
		"description":"apple Product",
		"price":"$1000",
		"quantity":"100"
	},
	{
		"name":"iphone13",
		"description":"apple Product",
		"price":"$1099",
		"quantity":"100"
	}
]

*/

func AddProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var products []Model.Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Database.DB.Create(&products)
	json.NewEncoder(w).Encode("products added successfully")
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var product Model.Product
	fmt.Println("printing product----")
	Database.DB.First(&product, params["id"])
	//fmt.Println(product)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var products []Model.Product
	Database.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var product Model.Product
	Database.DB.Find(&product, params["id"])
	numberOfProductsAvailable, _ := strconv.Atoi(product.Quantity)
	numberOfProductsRequired, _ := strconv.Atoi(params["quantity"])
	if numberOfProductsAvailable < numberOfProductsRequired {
		json.NewEncoder(w).Encode("Max Quantity available is " + strconv.Itoa(numberOfProductsAvailable))
		return
	}
	UpdateProductQuantity(&product, numberOfProductsAvailable-numberOfProductsRequired)
	json.NewEncoder(w).Encode("Buy Successful")
}

func UpdateProductQuantity(product *Model.Product, quantity int) {
	product.Quantity = strconv.Itoa(quantity)
	Database.DB.Save(&product)
}

func GetTop5Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var products []Model.Product
	Database.DB.Find(&products)
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
