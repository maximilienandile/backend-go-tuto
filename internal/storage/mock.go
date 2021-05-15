// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/interface.go

// Package storage is a generated GoMock package.
package storage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cart "github.com/maximilienandile/backend-go-tuto/internal/cart"
	category "github.com/maximilienandile/backend-go-tuto/internal/category"
	product "github.com/maximilienandile/backend-go-tuto/internal/product"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Categories mocks base method.
func (m *MockStorage) Categories() ([]category.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Categories")
	ret0, _ := ret[0].([]category.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Categories indicates an expected call of Categories.
func (mr *MockStorageMockRecorder) Categories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Categories", reflect.TypeOf((*MockStorage)(nil).Categories))
}

// CreateCart mocks base method.
func (m *MockStorage) CreateCart(cart cart.Cart, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", cart, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockStorageMockRecorder) CreateCart(cart, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockStorage)(nil).CreateCart), cart, userID)
}

// CreateCategory mocks base method.
func (m *MockStorage) CreateCategory(category category.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", category)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStorageMockRecorder) CreateCategory(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStorage)(nil).CreateCategory), category)
}

// CreateOrUpdateCart mocks base method.
func (m *MockStorage) CreateOrUpdateCart(userID, productID string, delta int) (cart.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateCart", userID, productID, delta)
	ret0, _ := ret[0].(cart.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateCart indicates an expected call of CreateOrUpdateCart.
func (mr *MockStorageMockRecorder) CreateOrUpdateCart(userID, productID, delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateCart", reflect.TypeOf((*MockStorage)(nil).CreateOrUpdateCart), userID, productID, delta)
}

// CreateProduct mocks base method.
func (m *MockStorage) CreateProduct(product product.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", product)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStorageMockRecorder) CreateProduct(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStorage)(nil).CreateProduct), product)
}

// GetCart mocks base method.
func (m *MockStorage) GetCart(userID string) (cart.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", userID)
	ret0, _ := ret[0].(cart.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockStorageMockRecorder) GetCart(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockStorage)(nil).GetCart), userID)
}

// Product mocks base method.
func (m *MockStorage) Product(ID string) (product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Product", ID)
	ret0, _ := ret[0].(product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Product indicates an expected call of Product.
func (mr *MockStorageMockRecorder) Product(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Product", reflect.TypeOf((*MockStorage)(nil).Product), ID)
}

// Products mocks base method.
func (m *MockStorage) Products() ([]product.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Products")
	ret0, _ := ret[0].([]product.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Products indicates an expected call of Products.
func (mr *MockStorageMockRecorder) Products() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Products", reflect.TypeOf((*MockStorage)(nil).Products))
}

// UpdateInventory mocks base method.
func (m *MockStorage) UpdateInventory(productID string, delta int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", productID, delta)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInventory indicates an expected call of UpdateInventory.
func (mr *MockStorageMockRecorder) UpdateInventory(productID, delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockStorage)(nil).UpdateInventory), productID, delta)
}

// UpdateProduct mocks base method.
func (m *MockStorage) UpdateProduct(input UpdateProductInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStorageMockRecorder) UpdateProduct(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStorage)(nil).UpdateProduct), input)
}
