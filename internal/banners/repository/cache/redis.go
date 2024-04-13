package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"avito-banners/internal/banners"
	"avito-banners/internal/models"

	"github.com/redis/go-redis/v9"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedis(redisClient *redis.Client) banners.Cache {
	return &redisRepo{client: redisClient}
}

func (r *redisRepo) SetBanner(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	newBytes, err := json.Marshal(banner)
	if err != nil {
		return err
	}

	// {"featureID:tagID": "bannerID"}
	for _, tagID := range banner.TagIDs {
		idKey := fmt.Sprintf("%d:%d", banner.FeatureID, tagID)
		if err := r.client.Set(ctx, idKey, bannerID, time.Minute*5).Err(); err != nil {
			return err
		}
	}

	// {"bannerID": "banner"}
	bannerKey := strconv.Itoa(bannerID)
	if err := r.client.Set(ctx, bannerKey, newBytes, time.Minute*5).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) GetBanner(ctx context.Context, featureID, tagID int) (*models.BannerRequest, error) {
	// Получаем ID баннера
	idKey := fmt.Sprintf("%d:%d", featureID, tagID)
	bannerID, err := r.client.Get(ctx, idKey).Result()
	if err != nil {
		return nil, err
	}

	// Получение баннера
	newBytes, err := r.client.Get(ctx, bannerID).Bytes()
	if err != nil {
		return nil, err
	}

	newBanner := new(models.BannerRequest)
	if err := json.Unmarshal(newBytes, newBanner); err != nil {
		return nil, err
	}

	return newBanner, nil
}

func (r *redisRepo) DeleteBanner(ctx context.Context, bannerID int) error {
	bannerKey := strconv.Itoa(bannerID)
	if err := r.client.Del(ctx, bannerKey).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) UpdateBanner(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	if err := r.DeleteBanner(ctx, bannerID); err != nil {
		return err
	}

	if err := r.SetBanner(ctx, bannerID, banner); err != nil {
		return err
	}

	return nil
}
