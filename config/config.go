package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type RedisConfig struct {
	Addr string
	Pass string
	DB   int
}

func LoadConfig() *Config {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverCfg := ServerConfig{
		Host: host,
		Port: port,
	}

	psqlCfg := PostgresConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
	}

	var redisDBInt int
	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		redisDBInt = 0
	} else {
		dbInt, err := strconv.Atoi(redisDB)
		if err != nil {
			log.Fatal(err)
		}
		redisDBInt = dbInt
	}

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = ":6379"
	}

	redisCfg := RedisConfig{
		Addr: redisHost,
		Pass: os.Getenv("REDIS_PASS"),
		DB:   redisDBInt,
	}

	return &Config{
		Server:   serverCfg,
		Postgres: psqlCfg,
		Redis:    redisCfg,
	}
}
