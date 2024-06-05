package config

import "os"

type HttpConfig struct {
	Domain string
	Port   string
}

func LoadHttpConfig() HttpConfig {
	return HttpConfig{
		Domain: os.Getenv("API_DOMAIN"),
		Port:   os.Getenv("PORT"),
	}
}
