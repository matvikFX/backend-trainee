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

	t.Run("Create", func(t *testing.T) {
		banner := &models.BannerRequest{
			TagIDs:    []int{1, 2},
			FeatureID: 1,
			IsActive:  true,
			Content: models.BannerContent{
				"name": "matvey",
			},
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(createBanner)).
			WithArgs(banner.Content, banner.IsActive).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(createFeature)).
			WithArgs(banner.FeatureID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		for _, tagID := range banner.TagIDs {
			mock.ExpectExec(regexp.QuoteMeta(createTag)).
				WithArgs(tagID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectQuery(regexp.QuoteMeta(getCurrentBannerID)).
			WillReturnRows(sqlmock.NewRows([]string{"bannerID"}).
				AddRow(1))
		mock.ExpectCommit()

		if _, err := psqlRepo.Create(context.Background(), banner); err != nil {
			t.Errorf("ошибка при создании нового баннера: %s", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("ожидаемый результат не был достигнут: %s", err)
		}
	})
}

func TestPsqlRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	psqlRepo := NewStorage(db)

	t.Run("Update", func(t *testing.T) {
		bannerID := 1
		banner := models.BannerRequest{
			TagIDs:    []int{1, 2},
			FeatureID: 1,
			IsActive:  true,
			Content:   models.BannerContent{},
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(updateBanner)).
			WithArgs(bannerID, banner.Content, banner.IsActive).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(updateFeature)).
			WithArgs(bannerID, banner.FeatureID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(deleteTags)).
			WithArgs(bannerID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		for _, tagID := range banner.TagIDs {
			mock.ExpectExec(regexp.QuoteMeta(createTag)).
				WithArgs(tagID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectCommit()

		if err := psqlRepo.Update(context.Background(), bannerID, &banner); err != nil {
			t.Errorf("ошибка при обновлении таблиц баннера и фичи: %s", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("ожидаемый результат не был достигнут: %s", err)
		}
	})
}

func TestPsqlRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	psqlRepo := NewStorage(db)

	t.Run("Delete", func(t *testing.T) {
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
	})
}
