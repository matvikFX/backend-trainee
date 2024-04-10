package usecase

import (
	"context"
	"testing"

	"avito-banners/internal/banners/mock"
	"avito-banners/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBannerUC_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPsqlRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewRedisMockRepository(ctrl)

	bannerUC := NewBannersUseCase(nil, mockPsqlRepo, mockRedisRepo)

	bannerID := 1
	banner := &models.BannerRequest{
		TagIDs:    []int{1, 2},
		FeatureID: 1,
		IsActive:  true,
		Content: models.BannerContent{
			"name": "matvey",
		},
	}

	ctx := context.Background()

	mockPsqlRepo.EXPECT().Create(ctx, banner).Return(bannerID, nil)
	if banner.IsActive {
		mockRedisRepo.EXPECT().SetBanner(ctx, bannerID, banner).Return(nil)
	}

	createdBanner, err := bannerUC.Create(ctx, banner)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, createdBanner)
}

func TestBannerUC_GetContent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPsqlRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewRedisMockRepository(ctrl)

	bannerUC := NewBannersUseCase(nil, mockPsqlRepo, nil)

	featureID := 1
	tagID := 1
	last_rev := true

	bannerContent := &models.BannerContent{}
	bannerReq := &models.BannerRequest{}

	ctx := context.Background()

	if last_rev {
		mockRedisRepo.EXPECT().GetBanner(ctx, tagID, featureID).Return(bannerReq, nil)
	}
	mockPsqlRepo.EXPECT().GetContent(ctx, tagID, featureID).Return(bannerContent, nil)

	content, err := bannerUC.GetContent(ctx, tagID, featureID, last_rev)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, content)
}

func TestBannerUC_GetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPsqlRepo := mock.NewMockRepository(ctrl)
	bannerUC := NewBannersUseCase(nil, mockPsqlRepo, nil)

	banners := []*models.Banner{}
	opts := models.BannerOptions{
		TagID:     1,
		FeatureID: 1,
		Limit:     10,
		Offset:    0,
	}

	ctx := context.Background()
	mockPsqlRepo.EXPECT().GetAll(ctx, &opts).Return(banners, nil)

	newBanners, err := bannerUC.GetAll(ctx, &opts)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, newBanners)
}

func TestBannerUC_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPsqlRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewRedisMockRepository(ctrl)

	bannerUC := NewBannersUseCase(nil, mockPsqlRepo, mockRedisRepo)

	bannerID := 1
	banner := &models.BannerRequest{
		TagIDs:    []int{1, 2},
		FeatureID: 1,
		IsActive:  true,
		Content:   models.BannerContent{},
	}

	ctx := context.Background()

	mockRedisRepo.EXPECT().SetBanner(ctx, bannerID, banner).Return(nil)
	mockPsqlRepo.EXPECT().Update(ctx, bannerID, banner).Return(nil)

	err := bannerUC.Update(ctx, bannerID, banner)
	require.NoError(t, err)
	require.Nil(t, err)
}

func TestBannerUC_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPsqlRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewRedisMockRepository(ctrl)

	bannerUC := NewBannersUseCase(nil, mockPsqlRepo, mockRedisRepo)

	ctx := context.Background()

	bannerID := 1

	mockPsqlRepo.EXPECT().Delete(ctx, bannerID).Return(nil)
	mockRedisRepo.EXPECT().DeleteBanner(ctx, bannerID).Return(nil)

	err := bannerUC.Delete(ctx, bannerID)
	require.NoError(t, err)
	require.Nil(t, err)
}
