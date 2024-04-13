package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"avito-banners/internal/banners/mock"
	"avito-banners/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestBannerHandlers_GetContext(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBannerUC := mock.NewMockUseCase(ctrl)
	bannerHandler := NewBannersHandlers(nil, mockBannerUC)

	handlerFunc := bannerHandler.GetContent()

	t.Run("GetFromCache", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/user_banner?tag_id=1&feature_id=1&use_last_revision=true", nil)
		require.NoError(t, err)
		req.Header.Set("Token", "user_token")

		res := httptest.NewRecorder()

		ctxBG := context.Background()
		req = req.WithContext(ctxBG)

		e := echo.New()
		ctx := e.NewContext(req, res)

		ctxWithReqID := GetReqCtx(ctx)

		tagID, err := strconv.Atoi(ctx.QueryParam("tag_id"))
		require.NoError(t, err)
		require.Greater(t, tagID, 0)

		featureID, err := strconv.Atoi(ctx.QueryParam("feature_id"))
		require.NoError(t, err)
		require.Greater(t, featureID, 0)

		var last_rev bool
		if val := ctx.QueryParam("use_last_revision"); val == "true" {
			last_rev = true
		} else {
			last_rev = false
		}

		mockContent := &models.BannerContent{}
		mockBannerUC.EXPECT().GetContent(ctxWithReqID, tagID, featureID, last_rev).Return(mockContent, nil)

		err = handlerFunc(ctx)
		require.NoError(t, err)
	})

	t.Run("GetFromDB", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/user_banner?tag_id=1&feature_id=1", nil)
		require.NoError(t, err)
		req.Header.Set("Token", "user_token")

		res := httptest.NewRecorder()

		ctxBG := context.Background()
		req = req.WithContext(ctxBG)

		e := echo.New()
		ctx := e.NewContext(req, res)
		ctxWithReqID := GetReqCtx(ctx)

		tagID, err := strconv.Atoi(ctx.QueryParam("tag_id"))
		require.NoError(t, err)
		require.Greater(t, tagID, 0)

		featureID, err := strconv.Atoi(ctx.QueryParam("feature_id"))
		require.NoError(t, err)
		require.Greater(t, featureID, 0)

		var last_rev bool
		if val := ctx.QueryParam("use_last_revision"); val == "true" {
			last_rev = true
		} else {
			last_rev = false
		}

		mockContent := &models.BannerContent{}
		mockBannerUC.EXPECT().GetContent(ctxWithReqID, tagID, featureID, last_rev).Return(mockContent, nil)

		err = handlerFunc(ctx)
		require.NoError(t, err)
	})
}
