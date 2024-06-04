package config

import "os"

type TursoConfig struct {
	Name  string
	Token string
}

func LoadTursoConfig() TursoConfig {
	return TursoConfig{
		Name:  os.Getenv("TURSO_NAME"),
		Token: os.Getenv("TURSO_DB_TOKEN"),
	}
}
