package http

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"avito-banners/config"
	"avito-banners/internal/banners"
	"avito-banners/internal/models"

	"github.com/labstack/echo/v4"
)

type bannersHandlers struct {
	cfg       *config.Config
	bannersUC banners.UseCase
}

func NewBannersHandlers(cfg *config.Config, uc banners.UseCase) banners.Handlers {
	return &bannersHandlers{cfg: cfg, bannersUC: uc}
}

// Middleware
func (h bannersHandlers) AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if tokenString := c.Request().Header["Token"]; tokenString != nil {
			if strings.Split(tokenString[0], "_")[0] == "admin" {
				return next(c)
			} else {
				return c.JSON(http.StatusForbidden, "Пользователь не имеет доступа")
			}
		} else {
			return c.JSON(http.StatusUnauthorized, "Пользователь не авторизован")
		}
	}
}

func (h bannersHandlers) UserAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if tokenString := c.Request().Header["Token"]; tokenString != nil {
			if rights := strings.Split(tokenString[0], "_")[0]; rights == "admin" || rights == "user" {
				return next(c)
			} else {
				return c.JSON(http.StatusForbidden, "Пользователь не имеет доступа")
			}
		} else {
			return c.JSON(http.StatusUnauthorized, "Пользователь не авторизован")
		}
	}
}

// Handlers
func (h bannersHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := GetReqCtx(c)

		bannerReq := &models.BannerRequest{}
		if err := c.Bind(bannerReq); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		bannerID, err := h.bannersUC.Create(ctx, bannerReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, bannerID)
	}
}

func (h bannersHandlers) GetContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := GetReqCtx(c)

		tagID, err := strconv.Atoi(c.QueryParam("tag_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		featureID, err := strconv.Atoi(c.QueryParam("feature_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		var last_rev bool
		if val := c.QueryParam("use_last_revision"); val == "true" {
			last_rev = true
		} else {
			last_rev = false
		}

		bannerContent := &models.BannerContent{}
		bannerContent, err = h.bannersUC.GetContent(ctx, tagID, featureID, last_rev)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, bannerContent)
	}
}

func (h bannersHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := GetReqCtx(c)

		bannerID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		banner, err := h.bannersUC.GetByID(ctx, bannerID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, banner)
	}
}

func (h bannersHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := GetReqCtx(c)

		tagID, err := strconv.Atoi(c.QueryParam("tag_id"))
		if err != nil && c.QueryParam("tag_id") != "" {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		featureID, err := strconv.Atoi(c.QueryParam("feature_id"))
		if err != nil && c.QueryParam("feature_id") != "" {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil && c.QueryParam("limit") != "" {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil && c.QueryParam("offset") != "" {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		bannerOpts := models.NewBannerOpts(
			featureID, tagID, limit, offset,
		)

		banners, err := h.bannersUC.GetAll(ctx, bannerOpts)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, banners)
	}
}

func (h bannersHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := GetReqCtx(c)

		bannerID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		bannerReq := &models.BannerRequest{}
		if err := c.Bind(bannerReq); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := h.bannersUC.Update(ctx, bannerID, bannerReq); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}

func (h bannersHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := GetReqCtx(c)

		bannerID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Некорректные данные")
		}

		if err := h.bannersUC.Delete(ctx, bannerID); err != nil {
			return c.JSON(http.StatusInternalServerError, "Некорректные данные")
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func GetReqCtx(c echo.Context) context.Context {
	type ReqIDCtxKey struct{}

	parent := c.Request().Context()
	key := ReqIDCtxKey{}
	val := c.Response().Header().Get(echo.HeaderXRequestID)
	return context.WithValue(parent, key, val)
}
