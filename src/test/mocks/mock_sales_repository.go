// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kartikeya/product_catalog_DIY/src/main/repository (interfaces: SalesRepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/kartikeya/product_catalog_DIY/src/main/model"
)

// MockSalesRepositoryInterface is a mock of SalesRepositoryInterface interface.
type MockSalesRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSalesRepositoryInterfaceMockRecorder
}

// MockSalesRepositoryInterfaceMockRecorder is the mock recorder for MockSalesRepositoryInterface.
type MockSalesRepositoryInterfaceMockRecorder struct {
	mock *MockSalesRepositoryInterface
}

// NewMockSalesRepositoryInterface creates a new mock instance.
func NewMockSalesRepositoryInterface(ctrl *gomock.Controller) *MockSalesRepositoryInterface {
	mock := &MockSalesRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockSalesRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSalesRepositoryInterface) EXPECT() *MockSalesRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSalesRepositoryInterface) Create(arg0 model.Sales) (*model.Sales, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*model.Sales)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSalesRepositoryInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSalesRepositoryInterface)(nil).Create), arg0)
}
