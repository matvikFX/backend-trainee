package main

import (
	"log"

	"avito-banners/config"
	"avito-banners/internal/server"
	"avito-banners/pkg/db/postgresql"
	"avito-banners/pkg/db/redis"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	log.Println("Starting server...")

	cfg := config.LoadConfig()

	psqlDB, err := postgresql.NewPsqlDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer psqlDB.Close()

	redisClient := redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	s := server.NewServer(cfg, psqlDB, redisClient)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
