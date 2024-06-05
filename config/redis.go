package config

import "os"

type RedisConfig struct {
	Host string
	Port string
	Url  string
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
		Url:  os.Getenv("REDIS_URL"),
	}
}
