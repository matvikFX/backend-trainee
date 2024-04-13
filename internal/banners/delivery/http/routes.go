package http

import (
	"avito-banners/internal/banners"

	"github.com/labstack/echo/v4"
)

func MapBannersRoutes(bannersRouter *echo.Group, h banners.Handlers) {
	bannersRouter.POST("", h.Create(), h.AdminAuth)
	bannersRouter.GET("", h.GetAll(), h.AdminAuth)
	bannersRouter.PATCH("/:id", h.Update(), h.AdminAuth)
	bannersRouter.DELETE("/:id", h.Delete(), h.AdminAuth)
}
