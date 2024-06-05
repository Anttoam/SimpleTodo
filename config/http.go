package config

import "os"

type HttpConfig struct {
	Domain string
}

func LoadHttpConfig() HttpConfig {
	return HttpConfig{
		Domain: os.Getenv("API_DOMAIN"),
	}
}
