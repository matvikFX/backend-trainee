package usecase

import (
	"context"
	"testing"

	"avito-banners/internal/banners/mock"
	"avito-banners/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBannerUC_GetContent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPsqlRepo := mock.NewMockRepository(ctrl)
	mockRedisRepo := mock.NewRedisMockRepository(ctrl)

	bannerUC := NewBannersUseCase(nil, mockPsqlRepo, mockRedisRepo)

	ctx := context.Background()

	featureID, tagID := 1, 1
	banner := &models.BannerRequest{
		TagIDs:    []int{1, 2},
		FeatureID: 1,
		IsActive:  true,
		Content:   &models.BannerContent{},
	}

	t.Run("GetFromCache", func(t *testing.T) {
		last_rev := true

		mockRedisRepo.EXPECT().GetBanner(ctx, tagID, featureID).Return(banner, nil)

		content, err := bannerUC.GetContent(ctx, tagID, featureID, last_rev)
		require.NoError(t, err)
		require.Nil(t, err)
		require.NotNil(t, content)
	})

	t.Run("GetFromDB", func(t *testing.T) {
		last_rev := false

		mockPsqlRepo.EXPECT().GetContent(ctx, tagID, featureID).Return(banner.Content, nil)

		content, err := bannerUC.GetContent(ctx, tagID, featureID, last_rev)
		require.NoError(t, err)
		require.Nil(t, err)
		require.NotNil(t, content)
	})
}
