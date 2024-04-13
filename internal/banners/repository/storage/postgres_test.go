package storage

import (
	"context"
	"regexp"
	"testing"

	"avito-banners/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestPsqlRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	psqlRepo := NewStorage(db)

	banner := &models.BannerRequest{
		TagIDs:    []int{1, 2},
		FeatureID: 1,
		IsActive:  true,
		Content: &models.BannerContent{
			"name": "matvey",
		},
	}
	expectID := 1

	rows := sqlmock.NewRows([]string{"id"}).AddRow(expectID)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(createBanner)).
		WithArgs(banner.FeatureID, banner.Content, banner.IsActive).
		WillReturnRows(rows)

	for _, tagID := range banner.TagIDs {
		mock.ExpectExec(regexp.QuoteMeta(createTag)).
			WithArgs(tagID).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}
	mock.ExpectCommit()

	if _, err := psqlRepo.Create(context.Background(), banner); err != nil {
		t.Errorf("ошибка при создании нового баннера: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("ожидаемый результат не был достигнут: %s", err)
	}
}

func TestPsqlRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	psqlRepo := NewStorage(db)

	bannerID := 1
	banner := models.BannerRequest{
		TagIDs:    []int{1, 2},
		FeatureID: 1,
		IsActive:  true,
		Content:   &models.BannerContent{},
	}

	updateString := `update banners set feature_id=$2, content=$3, is_active=$4 where id=$1`
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(updateString)).
		WithArgs(bannerID, banner.FeatureID, banner.Content, banner.IsActive).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(deleteTags)).
		WithArgs(bannerID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	for _, tagID := range banner.TagIDs {
		mock.ExpectExec(regexp.QuoteMeta(createTagByID)).
			WithArgs(bannerID, tagID).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}
	mock.ExpectCommit()

	if err := psqlRepo.Update(context.Background(), bannerID, &banner); err != nil {
		t.Errorf("ошибка при обновлении таблиц баннера: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("ожидаемый результат не был достигнут: %s", err)
	}
}

func TestPsqlRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	psqlRepo := NewStorage(db)

	bannerID := 3
	mock.ExpectExec(regexp.QuoteMeta(deleteBanner)).
		WithArgs(bannerID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := psqlRepo.Delete(context.Background(), bannerID); err != nil {
		t.Errorf("ошибка при удалении баннера: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("ожидаемый результат не был достигнут: %s", err)
	}
}
