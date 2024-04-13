package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"avito-banners/config"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

const (
	ctxTimeout = 10
)

type Server struct {
	echo  *echo.Echo
	cfg   *config.Config
	db    *sql.DB
	redis *redis.Client
}

func NewServer(cfg *config.Config, db *sql.DB, rc *redis.Client) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, redis: rc}
}

func (s *Server) Run() error {
	addr := s.cfg.Server.Host + ":" + s.cfg.Server.Port
	server := &http.Server{
		Addr: addr,
	}

	go func() {
		if err := s.echo.StartServer(server); err != nil {
			log.Fatal(err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*ctxTimeout)
	defer shutdown()

	return s.echo.Server.Shutdown(ctx)
}
