package mock

import (
	"context"
	"reflect"

	"avito-banners/internal/models"

	"github.com/golang/mock/gomock"
)

type MockRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisRepositoryMockRecorder
}

type MockRedisRepositoryMockRecorder struct {
	mock *MockRedisRepository
}

func NewRedisMockRepository(ctrl *gomock.Controller) *MockRedisRepository {
	mock := &MockRedisRepository{ctrl: ctrl}
	mock.recorder = &MockRedisRepositoryMockRecorder{mock}
	return mock
}

func (m *MockRedisRepository) EXPECT() *MockRedisRepositoryMockRecorder {
	return m.recorder
}

func (m *MockRedisRepository) SetBanner(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBanner", ctx, banner)
	ret0, _ := ret[0].(error)
	return ret0
}

// Указывает на ожидаемый результат функции SetBanner
func (mr *MockRedisRepositoryMockRecorder) SetBanner(ctx, bannerID, banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "SetBanner", reflect.TypeOf((*MockRedisRepository)(nil).SetBanner),
		ctx, bannerID, banner,
	)
}

func (m *MockRedisRepository) GetBanner(ctx context.Context, tagID, featureID int) (*models.BannerRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanner", ctx, tagID, featureID)
	ret0, _ := ret[0].(*models.BannerRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Указывает на ожидаемый результат функции GetBanner
func (mr *MockRedisRepositoryMockRecorder) GetBanner(ctx, tagID, featureID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetBanner", reflect.TypeOf((*MockRedisRepository)(nil).GetBanner),
		ctx, tagID, featureID,
	)
}

func (m *MockRedisRepository) UpdateBanner(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBanner", ctx, bannerID, banner)
	ret0, _ := ret[1].(error)
	return ret0
}

// Указывает на ожидаемый результат функции UpdateBanner
func (mr *MockRedisRepositoryMockRecorder) UpdateBanner(ctx, bannerID, banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "UpdateBanner", reflect.TypeOf((*MockRedisRepository)(nil).UpdateBanner),
		ctx, bannerID, banner,
	)
}

func (m *MockRedisRepository) DeleteBanner(ctx context.Context, bannerID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanner", ctx, bannerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Указывает на ожидаемый результат функции DeleteBanner
func (mr *MockRedisRepositoryMockRecorder) DeleteBanner(ctx, bannerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "DeleteBanner", reflect.TypeOf((*MockRedisRepository)(nil).DeleteBanner),
		ctx, bannerID,
	)
}
