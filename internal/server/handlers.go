package server

import (
	bHttp "avito-banners/internal/banners/delivery/http"
	"avito-banners/internal/banners/repository/cache"
	"avito-banners/internal/banners/repository/storage"
	"avito-banners/internal/banners/usecase"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Init repositories
	bRepo := storage.NewStorage(s.db)
	bRedisRepo := cache.NewRedis(s.redis)

	// Init useCases
	bUC := usecase.NewBannersUseCase(s.cfg, bRepo, bRedisRepo)

	// Init handlers
	bHandlers := bHttp.NewBannersHandlers(s.cfg, bUC)

	v1 := e.Group("/api/v1")

	v1.GET("/user_banner", bHandlers.GetContent())
	bannerGroup := v1.Group("/banner")
	bHttp.MapBannersRoutes(bannerGroup, bHandlers)

	return nil
}
