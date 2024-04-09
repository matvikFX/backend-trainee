package cache

import (
	"context"
	"fmt"

	"avito-banners/internal/models"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewRedis(redisClient *redis.Client) Cache {
	return &redisCache{client: redisClient}
}

func (r *redisCache) SetBanner(ctx context.Context, banner models.BannerRequest) error {
	return nil
}

func (r *redisCache) GetBanner(ctx context.Context, featureID int, tagIDs []int) (*models.BannerContent, error) {
	return nil, nil
}

func (r *redisCache) DeleteBanner(ctx context.Context, featureID int, tagIDs []int) error {
	return nil
}

func (r *redisCache) makeBannerKey(featureID int, tagIDs []int) string {
	bannerKey := fmt.Sprintf("featureID=%d;tagIDs=", featureID)
	for _, tagID := range tagIDs {
		bannerKey += fmt.Sprintf("%d,", tagID)
	}

	return bannerKey
}
