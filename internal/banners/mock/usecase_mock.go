package mock

import (
	"context"
	"reflect"

	"avito-banners/internal/models"

	"github.com/golang/mock/gomock"
)

type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT возвращает объект, который позволяет вызывающей стороне указать ожидаемое использование.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

func (m *MockUseCase) Create(ctx context.Context, banner *models.BannerRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, banner)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции Create
func (mr *MockUseCaseMockRecorder) Create(ctx, banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create),
		ctx, banner,
	)
}

func (m *MockUseCase) GetContent(ctx context.Context, tagID, featureID int, last_rev bool) (*models.BannerContent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", ctx, tagID, featureID, last_rev)
	ret0, _ := ret[0].(*models.BannerContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции GetContent
func (mr *MockUseCaseMockRecorder) GetContent(ctx, tagID, featureID, last_rev interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetContent", reflect.TypeOf((*MockRepository)(nil).GetContent),
		ctx, tagID, featureID, last_rev,
	)
}

func (m *MockUseCase) GetByID(ctx context.Context, bannerID int) (*models.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, bannerID)
	ret0, _ := ret[0].(*models.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции GetByID
func (mr *MockUseCaseMockRecorder) GetByID(ctx, bannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID),
		ctx, bannerID,
	)
}

func (m *MockUseCase) GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, opts)
	ret0, _ := ret[0].([]*models.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции GetAll
func (mr *MockUseCaseMockRecorder) GetAll(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll),
		ctx, opts,
	)
}

func (m *MockUseCase) Update(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, bannerID, banner)
	ret0, _ := ret[0].(error)
	return ret0
}

// Указывает на ожидаемый результат функции Update
func (mr *MockUseCaseMockRecorder) Update(ctx, bannerID, banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Update),
		ctx, bannerID, banner,
	)
}

func (m *MockUseCase) Delete(ctx context.Context, bannerID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bannerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Указывает на ожидаемый результат функции Delete
func (mr *MockUseCaseMockRecorder) Delete(ctx, bannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete),
		ctx, bannerID,
	)
}
