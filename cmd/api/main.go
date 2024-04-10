package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"avito-banners/config"
	"avito-banners/internal/server"
	"avito-banners/pkg/db/postgresql"
	"avito-banners/pkg/db/redis"
)

func main() {
	log.Println("Starting server...")

	readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Fatal(err)
	}

	writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		log.Fatal(err)
	}

	serverCfg := config.ServerConfig{
		Port:         os.Getenv("PORT"),
		ReadTimeout:  time.Duration(readTimeout),
		WriteTimeout: time.Duration(writeTimeout),
	}

	psqlCfg := config.PostgresConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatal(err)
	}

	redisCfg := config.RedisConfig{
		Addr: os.Getenv("REDIS_HOST"),
		Pass: os.Getenv("REDIS_PASS"),
		DB:   redisDB,
	}

	cfg := &config.Config{
		Server:   serverCfg,
		Postgres: psqlCfg,
		Redis:    redisCfg,
	}

	psqlDB, err := postgresql.NewPsqlDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer psqlDB.Close()

	redisClient := redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg, psqlDB, redisClient)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
