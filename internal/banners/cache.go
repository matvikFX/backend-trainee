package banners

import (
	"context"

	"avito-banners/internal/models"
)

type Cache interface {
	SetBanner(ctx context.Context, bannerID int, banner *models.BannerRequest) error
	GetBanner(ctx context.Context, featureID, tagID int) (*models.BannerRequest, error)
	UpdateBanner(ctx context.Context, bannerID int, banner *models.BannerRequest) error
	DeleteBanner(ctx context.Context, bannerID int) error
}
