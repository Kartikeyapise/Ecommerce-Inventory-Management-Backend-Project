package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func (m muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (m muxRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (m muxRouter) PUT(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}

func (m muxRouter) DELETE(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (m muxRouter) SERVE(port string) {
	log.Println("Mux Http Server listening on port", port, "........")
	http.ListenAndServe(port, muxDispatcher)
}

func NewMuxRouter() Router {
	return &muxRouter{}
}
