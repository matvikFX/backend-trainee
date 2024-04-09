package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type BannerContent map[string]interface{}

func (content BannerContent) Value() (driver.Value, error) {
	return json.Marshal(content)
}

func (content *BannerContent) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &content)
}

type BannerRequest struct {
	TagIDs    []int         `json:"tag_ids,omitempty" query:"tag_ids,omitempty"`
	FeatureID int           `json:"feature_id,omitempty" query:"feature_id,omitempty"`
	IsActive  bool          `json:"is_active,omitempty" query:"is_active,omitempty"`
	Content   BannerContent `json:"content"`
}

type Banner struct {
	ID        int           `json:"banner_id"`
	TagIDs    []int         `json:"tag_ids"`
	FeatureID int           `json:"feature_id"`
	IsActive  bool          `json:"is_active"`
	Content   BannerContent `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type BannerOptions struct {
	TagID     int
	FeatureID int
	Limit     int
	Offset    int
}

func NewBannerOpts(featureID, tagID, limit, offset int) *BannerOptions {
	banner := BannerOptions{
		FeatureID: featureID,
		TagID:     tagID,
		Limit:     limit,
		Offset:    offset,
	}
	banner.init()

	return &banner
}

func (opts *BannerOptions) init() {
	if opts.FeatureID < 0 {
		opts.FeatureID = 0
	}

	if opts.TagID < 0 {
		opts.TagID = 0
	}

	if opts.Limit < 0 {
		opts.Limit = 0
	}

	if opts.Offset < 0 {
		opts.Offset = 0
	}
}
