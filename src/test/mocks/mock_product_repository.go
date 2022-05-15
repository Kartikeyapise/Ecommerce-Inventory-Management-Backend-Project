// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kartikeya/product_catalog_DIY/src/main/repository (interfaces: ProductRepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/kartikeya/product_catalog_DIY/src/main/entity"
)

// MockProductRepositoryInterface is a mock of ProductRepositoryInterface interface.
type MockProductRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryInterfaceMockRecorder
}

// MockProductRepositoryInterfaceMockRecorder is the mock recorder for MockProductRepositoryInterface.
type MockProductRepositoryInterfaceMockRecorder struct {
	mock *MockProductRepositoryInterface
}

// NewMockProductRepositoryInterface creates a new mock instance.
func NewMockProductRepositoryInterface(ctrl *gomock.Controller) *MockProductRepositoryInterface {
	mock := &MockProductRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepositoryInterface) EXPECT() *MockProductRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductRepositoryInterface) Create(arg0 []entity.Product) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductRepositoryInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductRepositoryInterface)(nil).Create), arg0)
}

// FindAll mocks base method.
func (m *MockProductRepositoryInterface) FindAll() ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockProductRepositoryInterfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockProductRepositoryInterface)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockProductRepositoryInterface) FindById(arg0 string) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockProductRepositoryInterfaceMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockProductRepositoryInterface)(nil).FindById), arg0)
}

// Update mocks base method.
func (m *MockProductRepositoryInterface) Update(arg0 *entity.Product) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepositoryInterface)(nil).Update), arg0)
}