// Code generated by MockGen. DO NOT EDIT.
// Source: article_usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "yuki0920/go-notes/domain/model"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleUsecase is a mock of ArticleUsecase interface.
type MockArticleUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockArticleUsecaseMockRecorder
}

// MockArticleUsecaseMockRecorder is the mock recorder for MockArticleUsecase.
type MockArticleUsecaseMockRecorder struct {
	mock *MockArticleUsecase
}

// NewMockArticleUsecase creates a new mock instance.
func NewMockArticleUsecase(ctrl *gomock.Controller) *MockArticleUsecase {
	mock := &MockArticleUsecase{ctrl: ctrl}
	mock.recorder = &MockArticleUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleUsecase) EXPECT() *MockArticleUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleUsecase) Create(article *model.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", article)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArticleUsecaseMockRecorder) Create(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleUsecase)(nil).Create), article)
}

// Delete mocks base method.
func (m *MockArticleUsecase) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockArticleUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleUsecase)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockArticleUsecase) GetById(id int) (*model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockArticleUsecaseMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArticleUsecase)(nil).GetById), id)
}

// ListByPage mocks base method.
func (m *MockArticleUsecase) ListByPage(page int) ([]*model.Article, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPage", page)
	ret0, _ := ret[0].([]*model.Article)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByPage indicates an expected call of ListByPage.
func (mr *MockArticleUsecaseMockRecorder) ListByPage(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPage", reflect.TypeOf((*MockArticleUsecase)(nil).ListByPage), page)
}

// Update mocks base method.
func (m *MockArticleUsecase) Update(article *model.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", article)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleUsecaseMockRecorder) Update(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleUsecase)(nil).Update), article)
}