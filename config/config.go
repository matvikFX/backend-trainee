package config

import "time"

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
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

func LoadConfig() {
}
