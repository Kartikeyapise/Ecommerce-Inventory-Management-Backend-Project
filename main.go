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

	r.HandleFunc("/products", Controller.AddProducts).Methods("POST")
	r.HandleFunc("/product/{id}", Controller.GetProductById).Methods("GET")
	r.HandleFunc("/products", Controller.GetProducts).Methods("GET")
	r.HandleFunc("/buyProduct/{id}/{quantity}", Controller.BuyProduct).Methods("PUT")
	r.HandleFunc("/getTop5Products", Controller.GetTop5Products).Methods("GET")

	fmt.Println("Listening to requests.......")
	log.Fatal(http.ListenAndServe(":9000", r))

}

func main() {
	Database.ConnectDatabase()
	fmt.Println("Database connected.......")
	intilizeRouter()
}
