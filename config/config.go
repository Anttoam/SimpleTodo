package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Turso TursoConfig
	Redis RedisConfig
	Http  HttpConfig
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file")
	}

	return &Config{
		Turso: LoadTursoConfig(),
		Redis: LoadRedisConfig(),
		Http:  LoadHttpConfig(),
	}
}
