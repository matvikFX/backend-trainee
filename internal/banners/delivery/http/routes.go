package http

import (
	"avito-banners/internal/banners"

	"github.com/labstack/echo/v4"
)

func MapBannersRoutes(bannersRouter *echo.Group, h banners.Handlers) {
	// Пока пусть здесь будет, потом может че придумаю
	// bannersRouter.GET("/user", h.GetContent())

	// Нормальные
	bannersRouter.POST("", h.Create(), h.UserAuth)
	bannersRouter.GET("", h.GetAll(), h.UserAuth)
	bannersRouter.PATCH("/:id", h.Update(), h.UserAuth)
	bannersRouter.DELETE("/:id", h.Delete(), h.UserAuth)
}
