package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kartikeya/product_catalog_DIY/Controller"
	"github.com/kartikeya/product_catalog_DIY/Database"
	"log"
	"net/http"
)

func intilizeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/products", func(writer http.ResponseWriter, request *http.Request) {
		Controller.AddProducts(writer, request, Database.DB)
	}).Methods("POST")

	r.HandleFunc("/product/{id}", func(writer http.ResponseWriter, request *http.Request) {
		Controller.GetProductById(writer, request, Database.DB)
	}).Methods("GET")

	r.HandleFunc("/products", func(writer http.ResponseWriter, request *http.Request) {
		Controller.GetProducts(writer, request, Database.DB)
	}).Methods("GET")

	r.HandleFunc("/buyProduct/{id}/{quantity}", func(writer http.ResponseWriter, request *http.Request) {
		Controller.BuyProduct(writer, request, Database.DB)
	}).Methods("PUT")
	
	r.HandleFunc("/getTop5Products", func(writer http.ResponseWriter, request *http.Request) {
		Controller.GetTop5Products(writer, request, Database.DB)
	}).Methods("GET")

	fmt.Println("Listening to requests.......")
	log.Fatal(http.ListenAndServe(":9000", r))

}

func main() {
	Database.ConnectDatabase()
	fmt.Println("Database connected.......")
	intilizeRouter()
}
