package mock

import (
	"context"
	"reflect"

	"avito-banners/internal/models"

	"github.com/golang/mock/gomock"
)

type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

func (m *MockRepository) Create(ctx context.Context, banner *models.BannerRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, banner)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции Create
func (mr *MockRepositoryMockRecorder) Create(ctx, banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create),
		ctx, banner,
	)
}

func (m *MockRepository) GetContent(ctx context.Context, tagID, featureID int) (*models.BannerContent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", ctx, tagID, featureID)
	ret0, _ := ret[0].(*models.BannerContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции GetContent
func (mr *MockRepositoryMockRecorder) GetContent(ctx, tagID, featureID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetContent", reflect.TypeOf((*MockRepository)(nil).GetContent),
		ctx, tagID, featureID,
	)
}

func (m *MockRepository) GetByID(ctx context.Context, bannerID int) (*models.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, bannerID)
	ret0, _ := ret[0].(*models.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции GetByID
func (mr *MockRepositoryMockRecorder) GetByID(ctx, bannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID),
		ctx, bannerID,
	)
}

func (m *MockRepository) GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, opts)
	ret0, _ := ret[0].([]*models.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции GetAll
func (mr *MockRepositoryMockRecorder) GetAll(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll),
		ctx, opts,
	)
}

func (m *MockRepository) Update(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, bannerID, banner)
	ret0, _ := ret[1].(error)
	return ret0
}

// Указывает на ожидаемый результат функции Update
func (mr *MockRepositoryMockRecorder) Update(ctx, bannerID, banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update),
		ctx, bannerID, banner,
	)
}

func (m *MockRepository) Delete(ctx context.Context, bannerID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bannerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Указывает на ожидаемый результат функции Delete
func (mr *MockRepositoryMockRecorder) Delete(ctx, bannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete),
		ctx, bannerID,
	)
}
