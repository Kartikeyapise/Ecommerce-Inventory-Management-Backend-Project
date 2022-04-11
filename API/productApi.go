package API

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/Database"
	"github.com/kartikeya/product_catalog_DIY/Model"
	"net/http"
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

func GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inproduct API")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var product Model.Product
	fmt.Println("printing product----")
	Database.DB.First(&product, params["id"])
	//fmt.Println(product)
	json.NewEncoder(w).Encode(product)
}
