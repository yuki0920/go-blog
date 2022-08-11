// Code generated by MockGen. DO NOT EDIT.
// Source: category_usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "yuki0920/go-notes/domain/model"

	gomock "github.com/golang/mock/gomock"
)

// MockCategoryUsecase is a mock of CategoryUsecase interface.
type MockCategoryUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryUsecaseMockRecorder
}

// MockCategoryUsecaseMockRecorder is the mock recorder for MockCategoryUsecase.
type MockCategoryUsecaseMockRecorder struct {
	mock *MockCategoryUsecase
}

// NewMockCategoryUsecase creates a new mock instance.
func NewMockCategoryUsecase(ctrl *gomock.Controller) *MockCategoryUsecase {
	mock := &MockCategoryUsecase{ctrl: ctrl}
	mock.recorder = &MockCategoryUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryUsecase) EXPECT() *MockCategoryUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryUsecase) Create(category *model.Category) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", category)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCategoryUsecaseMockRecorder) Create(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryUsecase)(nil).Create), category)
}

// List mocks base method.
func (m *MockCategoryUsecase) List() ([]*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCategoryUsecaseMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCategoryUsecase)(nil).List))
}
