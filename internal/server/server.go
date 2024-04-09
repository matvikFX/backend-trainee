package server

import (
	"avito-banners/config"
	storage "avito-banners/internal/repository/postgresql"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	echo  *echo.Echo
	cfg   *config.Config
	db    *storage.Storage
	cache *redis.Client
}

func NewServer(cfg *config.Config, db *storage.Storage, rc *redis.Client) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, cache: rc}
}

func (s *Server) Run() error {
	return nil
}
