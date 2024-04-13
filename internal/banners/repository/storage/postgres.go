package storage

import (
	"context"
	"database/sql"
	"fmt"

	"avito-banners/internal/banners"
	"avito-banners/internal/models"
)

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) banners.Repository {
	return &storage{db: db}
}

func (s *storage) GetContent(ctx context.Context, tagID, featureID int) (*models.BannerContent, error) {
	var content models.BannerContent

	row := s.db.QueryRowContext(ctx, getContent, tagID, featureID)
	if err := row.Scan(&content); err != nil {
		return nil, err
	}

	return &content, nil
}

func (s *storage) GetByID(ctx context.Context, id int) (*models.Banner, error) {
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

func (s *storage) GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error) {
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
		banner := &models.Banner{}
		if err := rows.Scan(&banner.ID, &banner.Content, &banner.IsActive,
			&banner.CreatedAt, &banner.UpdatedAt, &banner.FeatureID,
		); err != nil {
			return nil, err
		}

		tagIDs, err := s.getTags(tx, ctx, banner.ID)
		if err != nil {
			return nil, err
		}

		banner.TagIDs = tagIDs
		banners = append(banners, banner)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return banners, nil
}

func (s *storage) Create(ctx context.Context, banner *models.BannerRequest) (int, error) {
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

	var bannerID int
	row := tx.QueryRowContext(ctx, createBanner,
		banner.FeatureID, banner.Content, banner.IsActive,
	)
	if err := row.Scan(&bannerID); err != nil {
		return 0, err
	}

	// Добавление в banner_tag
	for _, tagID := range banner.TagIDs {
		if _, err := tx.ExecContext(ctx, createTag, tagID); err != nil {
			return 0, fmt.Errorf("ошибка при добавлении в таблицу banner_tag %s", err.Error())
		}
	}

	return bannerID, nil
}

func (s *storage) Update(ctx context.Context, bannerID int, banner *models.BannerRequest) error {
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
	updateString, args := s.makeUpdateString(bannerID, banner.FeatureID, banner.Content, banner.IsActive)
	if _, err := tx.ExecContext(ctx, updateString, args...); err != nil {
		return err
	}

	// Удаляем все теги баннера, что есть, и добавляем новые
	if _, err := tx.ExecContext(ctx, deleteTags, bannerID); err != nil {
		return err
	}

	for _, tagID := range banner.TagIDs {
		if _, err := tx.ExecContext(ctx, createTagByID, bannerID, tagID); err != nil {
			return err
		}
	}

	return nil
}

func (s *storage) Delete(ctx context.Context, banner_id int) error {
	if _, err := s.db.ExecContext(ctx, deleteBanner, banner_id); err != nil {
		return err
	}

	return nil
}

func (s *storage) getTags(tx *sql.Tx, ctx context.Context, bannerID int) ([]int, error) {
	// rows, err := tx.QueryContext(ctx, getTagsByID, bannerID)
	rows, err := s.db.QueryContext(ctx, getTagsByID, bannerID)
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

func (s *storage) getQueryFromOpts(opts *models.BannerOptions) string {
	optsString := ""
	if opts.FeatureID != 0 && opts.TagID != 0 {
		optsString += fmt.Sprintf(
			"where b_f.feature_id = %d and b_t.tag_id = %d ",
			opts.FeatureID, opts.TagID,
		)
	} else {
		if opts.FeatureID != 0 {
			optsString += fmt.Sprintf("where b.feature_id = %d ", opts.FeatureID)
		}

		if opts.TagID != 0 {
			optsString += fmt.Sprintf("where b_t.tag_id = %d ", opts.TagID)
		}
	}

	optsString += "\ngroup by b.id, b.feature_id "

	if opts.Limit != 0 {
		optsString += fmt.Sprintf("limit %d ", opts.Limit)
	}

	if opts.Offset != 0 {
		optsString += fmt.Sprintf("offset %d", opts.Offset)
	}

	queryString := getBannerWithoutTags + optsString + ";"

	return queryString
}

func (s *storage) makeUpdateString(bannerID, featureID int, content *models.BannerContent, isActive bool) (string, []any) {
	args := []any{bannerID}
	updateString := "update banners set"
	counter := 1

	if featureID != 0 {
		counter++
		updateString += fmt.Sprintf(" feature_id=$%d", counter)
		args = append(args, featureID)
	}

	if content != nil {
		if counter == 2 {
			updateString += ","
		}
		counter++
		updateString += fmt.Sprintf(" content=$%d", counter)
		args = append(args, content)
	}

	if counter >= 2 {
		updateString += ","
	}
	counter++
	updateString += fmt.Sprintf(" is_active=$%d", counter)
	args = append(args, isActive)

	updateString += " where id=$1"

	return updateString, args
}
