package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/kartikeya/product_catalog_DIY/src/main/controller"
	"github.com/kartikeya/product_catalog_DIY/src/main/model"
	"github.com/kartikeya/product_catalog_DIY/src/main/view"
	"github.com/kartikeya/product_catalog_DIY/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	userController := controller.UserController{UserService: mockUserService}
	expectedUser := &model.User{
		Email: "Email",
		Name:  "Name",
		Type:  "Type",
	}

	mockUserService.EXPECT().CreateUser(gomock.Any()).Return(expectedUser, nil).Times(1)
	req_body := []byte(`{
			"name":"kartikeya2",
			"email":"kartikeya2@gmail.com",
			"type":"USER"
		}
	`)
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(req_body))

	handler := http.HandlerFunc(userController.CreateUser)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	status := response.Code
	assert.Equal(t, status, http.StatusOK)

	var user model.User
	json.NewDecoder(response.Body).Decode(&user)

	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Type, user.Type)

}

func TestCreateUserWhenServiceReturnsAnError(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	userController := controller.UserController{UserService: mockUserService}

	mockUserService.EXPECT().CreateUser(gomock.Any()).Return(nil, errors.New("error")).Times(1)
	req_body := []byte(`{
			"name":"kartikeya2",
			"email":"kartikeya2@gmail.com",
			"type":"USER"
		}
	`)
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(req_body))

	handler := http.HandlerFunc(userController.CreateUser)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	assert.Equal(t, message.Message, "error")

}

func TestCreateUserWhenReqBodyHasWrongFormat(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mocks.NewMockUserServiceInterface(mockCtrl)
	userController := controller.UserController{UserService: mockUserService}

	req_body := []byte(`garbage`)
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(req_body))

	handler := http.HandlerFunc(userController.CreateUser)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	status := response.Code
	assert.Equal(t, status, http.StatusInternalServerError)

	var message view.ResponseMessage
	json.NewDecoder(response.Body).Decode(&message)

	assert.Equal(t, message.Message, "Cannot Extract User Information from Request Body")

}
