package controller

import (
	"encoding/json"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"github.com/kartikeya/product_catalog_DIY/src/main/service"
	"github.com/kartikeya/product_catalog_DIY/src/main/view"
	"net/http"
)

type UserController struct {
	UserService service.UserServiceInterface
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{
			Message: "Cannot Extract User Information from Request Body",
		})
		return
	}
	returnedUser, err := c.UserService.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(view.ResponseMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(returnedUser)
}
