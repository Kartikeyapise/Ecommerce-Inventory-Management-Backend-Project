package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/API"
	"github.com/kartikeya/product_catalog_DIY/Database"
	"log"
	"net/http"
)

func intilizeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/products", API.AddProducts).Methods("POST")
	r.HandleFunc("/product/{id}", API.GetProductById).Methods("GET")
	r.HandleFunc("/products", API.GetProducts).Methods("GET")
	r.HandleFunc("/buyProduct/{id}/{quantity}", API.BuyProduct).Methods("PUT")
	r.HandleFunc("/getTop5Products", API.GetTop5Products).Methods("GET")

	fmt.Println("Listening to requests.......")
	log.Fatal(http.ListenAndServe(":9000", r))

}

func main() {
	Database.ConnectDatabase()
	fmt.Println("Database connected.......")
	intilizeRouter()
}
