package storage

import (
	"context"
	"database/sql"
	"fmt"

	"avito-banners/internal/banners"
	"avito-banners/internal/models"
)

type postgresql struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) banners.Repository {
	return &postgresql{db: db}
}

func (s *postgresql) GetContent(ctx context.Context, tagID, featureID int) (*models.BannerContent, error) {
	var bannerID int
	var content models.BannerContent

	row := s.db.QueryRowContext(ctx, getContent, tagID, featureID)
	if err := row.Scan(&bannerID, &content); err != nil {
		return nil, err
	}

	return &content, nil
}

func (s *postgresql) GetByID(ctx context.Context, id int) (*models.Banner, error) {
	var banner models.Banner

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	row := tx.QueryRowContext(ctx, getByID, id)
	if err := row.Scan(&banner.ID, &banner.Content, &banner.IsActive,
		&banner.CreatedAt, &banner.UpdatedAt, &banner.FeatureID,
	); err != nil {
		return nil, err
	}

	tagIDs, err := s.getTags(tx, ctx, id)
	if err != nil {
		return nil, err
	}
	banner.TagIDs = tagIDs

	return &banner, nil
}

func (s *postgresql) GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error) {
	var banners []*models.Banner

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	queryString := s.getQueryFromOpts(opts)
	rows, err := tx.QueryContext(ctx, queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		banner := new(models.Banner)
		if err := rows.Scan(&banner.ID, &banner.Content, &banner.IsActive,
			&banner.CreatedAt, &banner.UpdatedAt, &banner.FeatureID,
		); err != nil {
			return banners, err
		}

		tagIDs, err := s.getTags(tx, ctx, banner.ID)
		if err != nil {
			return nil, err
		}

		banner.TagIDs = tagIDs
		banners = append(banners, banner)
	}

	if err = rows.Err(); err != nil {
		return banners, err
	}

	return banners, nil
}

func (s *postgresql) Create(ctx context.Context, banner *models.BannerRequest) (int, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// Добавление в banners
	if _, err := tx.ExecContext(ctx, createBanner, banner.Content, banner.IsActive); err != nil {
		return 0, err
	}

	// Добавление в banner_feature
	if _, err := tx.ExecContext(ctx, createFeature, banner.FeatureID); err != nil {
		return 0, err
	}

	// Добавление в banner_tag
	for _, tagID := range banner.TagIDs {
		if _, err := tx.ExecContext(ctx, createTag, tagID); err != nil {
			return 0, fmt.Errorf("ошибка при добавлении в таблицу banner_tag %s", err.Error())
		}
	}

	// Получение id нового баннера
	var bannerID int
	row := tx.QueryRow("select currval('banners_id_seq');")
	if err := row.Scan(&bannerID); err != nil {
		return 0, err
	}

	return 0, nil
}

func (s *postgresql) Update(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// Изменение баннера
	if _, err := tx.ExecContext(ctx, updateBanner,
		bannerID, banner.Content, banner.IsActive,
	); err != nil {
		return err
	}

	// Изменение фичи
	if _, err := tx.ExecContext(ctx, updateFeature,
		bannerID, banner.FeatureID,
	); err != nil {
		return err
	}

	// Удаляем все теги баннера, что есть, и добавляем новые
	if _, err := tx.ExecContext(ctx, deleteTags, bannerID); err != nil {
		return err
	}

	for _, tag := range banner.TagIDs {
		if _, err := tx.ExecContext(ctx, createTag, tag); err != nil {
			return err
		}
	}

	return nil
}

func (s *postgresql) Delete(ctx context.Context, banner_id int) error {
	if _, err := s.db.ExecContext(ctx, deleteBanner, banner_id); err != nil {
		return err
	}

	return nil
}

func (s *postgresql) getTags(tx *sql.Tx, ctx context.Context, banner_id int) ([]int, error) {
	rows, err := tx.QueryContext(ctx, getTagsByID, banner_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tagIDs []int
	for rows.Next() {
		var tagID int
		if err := rows.Scan(&tagID); err != nil {
			return nil, err
		}
		tagIDs = append(tagIDs, tagID)
	}

	if err = rows.Err(); err != nil {
		return tagIDs, err
	}

	return tagIDs, nil
}

func (s *postgresql) getQueryFromOpts(opts *models.BannerOptions) string {
	optsString := ""
	if opts.FeatureID != 0 && opts.TagID != 0 {
		optsString += fmt.Sprintf(
			"where b_f.feature_id = %d and b_t.tag_id = %d ",
			opts.FeatureID, opts.TagID,
		)
	} else {
		if opts.FeatureID != 0 {
			optsString += fmt.Sprintf("where b_f.feature_id = %d ", opts.FeatureID)
		}

		if opts.TagID != 0 {
			optsString += fmt.Sprintf("where b_t.tag_id = %d ", opts.TagID)
		}
	}

	optsString += "group by b.id, b_f.feature_id "

	if opts.Limit != 0 {
		optsString += fmt.Sprintf("limit %d ", opts.Limit)
	}

	if opts.Offset != 0 {
		optsString += fmt.Sprintf("offset %d", opts.Offset)
	}

	queryString := getBannerWithoutTags + optsString + ";"

	return queryString
}
