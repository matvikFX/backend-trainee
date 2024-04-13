package server

import (
	"net/http"

	bHttp "avito-banners/internal/banners/delivery/http"
	"avito-banners/internal/banners/repository/cache"
	"avito-banners/internal/banners/repository/storage"
	"avito-banners/internal/banners/usecase"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	bRepo := storage.NewStorage(s.db)
	bRedisRepo := cache.NewRedis(s.redis)

	// Init useCases
	bUC := usecase.NewBannersUseCase(s.cfg, bRepo, bRedisRepo)

	// Init handlers
	bHandlers := bHttp.NewBannersHandlers(s.cfg, bUC)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	v1 := e.Group("/api/v1")

	v1.GET("/user_banner", bHandlers.GetContent(), bHandlers.UserAuth)
	bannerGroup := v1.Group("/banner")
	bHttp.MapBannersRoutes(bannerGroup, bHandlers)

	return nil
}
