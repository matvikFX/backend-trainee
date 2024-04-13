package cache

import (
	"context"
	"log"
	"testing"

	"avito-banners/internal/banners"
	"avito-banners/internal/models"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func SetupRedis() banners.Cache {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	redisRepo := NewRedis(client)
	return redisRepo
}

func TestRedisRepo_SetBanner(t *testing.T) {
	t.Parallel()

	redisRepo := SetupRedis()

	t.Run("SetBanner", func(t *testing.T) {
		bannerID := 1
		banner := &models.BannerRequest{
			TagIDs:    []int{1, 2},
			FeatureID: 1,
			IsActive:  true,
			Content: &models.BannerContent{
				"title": "test",
			},
		}

		err := redisRepo.SetBanner(context.Background(), bannerID, banner)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

func TestRedisRepo_GetBanner(t *testing.T) {
	t.Parallel()

	redisRepo := SetupRedis()

	t.Run("GetBanner", func(t *testing.T) {
		bannerID := 1
		banner := &models.BannerRequest{
			FeatureID: 1,
			TagIDs:    []int{1, 2},
			Content: &models.BannerContent{
				"title": "test",
				"text":  "newText",
			},
			IsActive: true,
		}

		newBanner, err := redisRepo.GetBanner(context.Background(), banner.FeatureID, banner.TagIDs[0])
		require.Nil(t, newBanner)
		require.NotNil(t, err)

		err = redisRepo.SetBanner(context.Background(), bannerID, banner)
		require.NoError(t, err)
		require.Nil(t, err)

		newBanner, err = redisRepo.GetBanner(context.Background(), banner.FeatureID, banner.TagIDs[0])
		require.NoError(t, err)
		require.Nil(t, err)
		require.NotNil(t, newBanner)
	})
}

func TestRedisRepo_DeleteBanner(t *testing.T) {
	t.Parallel()

	redisRepo := SetupRedis()

	t.Run("DeleteBanner", func(t *testing.T) {
		bannerID := 1

		err := redisRepo.DeleteBanner(context.Background(), bannerID)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}

func TestRedisRepo_UpdateBanner(t *testing.T) {
	t.Parallel()

	redisRepo := SetupRedis()

	t.Run("UpdateBanner", func(t *testing.T) {
		bannerID := 1
		banner := &models.BannerRequest{
			TagIDs:    []int{1, 2},
			FeatureID: 1,
			IsActive:  true,
			Content: &models.BannerContent{
				"title": "test",
				"text":  "newText",
			},
		}

		err := redisRepo.UpdateBanner(context.Background(), bannerID, banner)
		require.NoError(t, err)
		require.Nil(t, err)

		newBanner, err := redisRepo.GetBanner(context.Background(), banner.FeatureID, banner.TagIDs[0])
		require.NoError(t, err)
		require.Nil(t, err)
		require.NotNil(t, newBanner)
		require.Equal(t, banner, newBanner)
	})
}
