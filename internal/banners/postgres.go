package banners

import (
	"context"

	"avito-banners/internal/models"
)

type Repository interface {
	Create(ctx context.Context, banner *models.BannerRequest) (int, error)
	GetContent(ctx context.Context, tagID, featureID int) (*models.BannerContent, error)
	GetByID(ctx context.Context, bannerID int) (*models.Banner, error)
	GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error)
	Update(ctx context.Context, bannerID int, banner *models.BannerRequest) error
	Delete(ctx context.Context, bannerID int) error
}
