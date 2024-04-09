package cache

import (
	"context"

	"avito-banners/internal/models"
)

type Cache interface {
	SetBanner(ctx context.Context, banner models.BannerRequest) error
	GetBanner(ctx context.Context, featureID int, tagIDs []int) (*models.BannerContent, error)
	DeleteBanner(ctx context.Context, featureID int, tagIDs []int) error
}
