package storage

import (
	"context"
	"database/sql"
	"fmt"

	"avito-banners/internal/models"

	_ "github.com/lib/pq"
)

type postgresql struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &postgresql{db: db}
}

func (s *postgresql) GetUserBanner(ctx context.Context, tag_id, feature_id int) (*models.BannerContent, error) {
	return nil, nil
}

func (s *postgresql) GetByID(ctx context.Context, id int) (*models.Banner, error) {
	return nil, nil
}

func (s *postgresql) GetAll(ctx context.Context, opts *models.BannerOptions) ([]*models.Banner, error) {
	return nil, nil
}

func (s *postgresql) Create(ctx context.Context, banner *models.BannerRequest) (int, error) {
	return 0, nil
}

func (s *postgresql) Update(ctx context.Context, id int, banner *models.BannerRequest) error {
	return nil
}

func (s *postgresql) Delete(ctx context.Context, banner_id int) error {
	delString := `
		delete from banners
		where id = $1;
	`

	if _, err := s.db.Exec(delString, banner_id); err != nil {
		return err
	}

	return nil
}

func (s *postgresql) getTags(banner_id int) ([]int, error) {
	tagString := `
		select tag_id
		from banner_tag 
		where banner_id = $1;
	`

	rows, err := s.db.Query(tagString, banner_id)
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

func getQueryFromOpts(opts *models.BannerOptions) string {
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

	// , array_agg(distinct b_t.tag_id) as "tag_ids"
	queryString := `
	select
		b.id, b.content, b.is_active, b.created_at, b.updated_at,
		b_f.feature_id
	from banners b
	join banner_tag b_t on b.id = b_t.banner_id
	join banner_feature b_f on b.id = b_f.banner_id
	` + optsString + ";"

	return queryString
}
