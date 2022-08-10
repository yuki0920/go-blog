// Code generated by MockGen. DO NOT EDIT.
// Source: article_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "yuki0920/go-notes/domain/model"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleRepository is a mock of ArticleRepository interface.
type MockArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepositoryMockRecorder
}

// MockArticleRepositoryMockRecorder is the mock recorder for MockArticleRepository.
type MockArticleRepositoryMockRecorder struct {
	mock *MockArticleRepository
}

// NewMockArticleRepository creates a new mock instance.
func NewMockArticleRepository(ctrl *gomock.Controller) *MockArticleRepository {
	mock := &MockArticleRepository{ctrl: ctrl}
	mock.recorder = &MockArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleRepository) EXPECT() *MockArticleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleRepository) Create(article *model.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", article)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArticleRepositoryMockRecorder) Create(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleRepository)(nil).Create), article)
}

// CreateCategories mocks base method.
func (m *MockArticleRepository) CreateCategories(id int, article *model.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategories", id, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCategories indicates an expected call of CreateCategories.
func (mr *MockArticleRepositoryMockRecorder) CreateCategories(id, article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategories", reflect.TypeOf((*MockArticleRepository)(nil).CreateCategories), id, article)
}

// Delete mocks base method.
func (m *MockArticleRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockArticleRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleRepository)(nil).Delete), id)
}

// DeleteCategories mocks base method.
func (m *MockArticleRepository) DeleteCategories(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategories", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategories indicates an expected call of DeleteCategories.
func (mr *MockArticleRepositoryMockRecorder) DeleteCategories(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategories", reflect.TypeOf((*MockArticleRepository)(nil).DeleteCategories), id)
}

// GetById mocks base method.
func (m *MockArticleRepository) GetById(id int) (*model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockArticleRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArticleRepository)(nil).GetById), id)
}

// ListByPage mocks base method.
func (m *MockArticleRepository) ListByPage(page int) ([]*model.Article, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPage", page)
	ret0, _ := ret[0].([]*model.Article)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByPage indicates an expected call of ListByPage.
func (mr *MockArticleRepositoryMockRecorder) ListByPage(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPage", reflect.TypeOf((*MockArticleRepository)(nil).ListByPage), page)
}

// Update mocks base method.
func (m *MockArticleRepository) Update(article *model.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", article)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleRepositoryMockRecorder) Update(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleRepository)(nil).Update), article)
}