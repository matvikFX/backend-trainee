package usecase

import (
	"context"

	"avito-banners/config"
	"avito-banners/internal/banners"
	"avito-banners/internal/models"
)

type bannersUC struct {
	cfg       *config.Config
	psqlRepo  banners.Repository
	redisRepo banners.Cache
}

func NewBannersUseCase(cfg *config.Config, psqlRepo banners.Repository, redisRepo banners.Cache) banners.UseCase {
	return &bannersUC{
		cfg:       cfg,
		psqlRepo:  psqlRepo,
		redisRepo: redisRepo,
	}
}

func (uc *bannersUC) Create(ctx context.Context, banner *models.BannerRequest) (int, error) {
	bannerID, err := uc.psqlRepo.Create(ctx, banner)
	if err != nil {
		return 0, err
	}

	if banner.IsActive {
		if err := uc.redisRepo.SetBanner(ctx, bannerID, banner); err != nil {
			return 0, err
		}
	}

	return bannerID, nil
}

func (uc *bannersUC) GetContent(ctx context.Context, tagID, featureID int, last_rev bool) (*models.BannerContent, error) {
	if last_rev {
		banner := &models.BannerRequest{}
		banner, err := uc.redisRepo.GetBanner(ctx, featureID, tagID)
		if err == nil {
			return &banner.Content, nil
		} // else В redis нет записи данного баннера
	}

	content, err := uc.psqlRepo.GetContent(ctx, tagID, featureID)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (uc *bannersUC) GetByID(ctx context.Context, bannerID int) (*models.Banner, error) {
	banner, err := uc.psqlRepo.GetByID(ctx, bannerID)
	if err != nil {
		return nil, err
	}

	return banner, nil
}

func (uc *bannersUC) GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error) {
	banners, err := uc.psqlRepo.GetAll(ctx, opts)
	if err != nil {
		return nil, err
	}

	return banners, nil
}

func (uc *bannersUC) Update(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	if err := uc.redisRepo.SetBanner(ctx, bannerID, banner); err != nil {
		return err
	}

	if err := uc.psqlRepo.Update(ctx, bannerID, banner); err != nil {
		return err
	}

	return nil
}

func (uc *bannersUC) Delete(ctx context.Context, bannerID int) error {
	uc.redisRepo.DeleteBanner(ctx, bannerID)

	if err := uc.psqlRepo.Delete(ctx, bannerID); err != nil {
		return err
	}

	return nil
}
