package storage

import (
	"context"

	"avito-banners/internal/models"
)

type Storage interface {
	GetUserBanner(ctx context.Context, tag_ids, feature_id int) (*models.BannerContent, error)
	GetByID(ctx context.Context, banner_id int) (*models.Banner, error)
	GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error)
	Create(ctx context.Context, banner *models.BannerRequest) (int, error)
	Update(ctx context.Context, id int, banner *models.BannerRequest) error
	Delete(ctx context.Context, banner_id int) error
}
