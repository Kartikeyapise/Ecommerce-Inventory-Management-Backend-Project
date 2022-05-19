package controller

import "net/http"

type UserControllerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
