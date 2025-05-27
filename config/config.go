package config

import (
	"log"
	"os"
)

type Config struct {
	PostgresURI string
	Port        string
}

func NewConfig() *Config {
	cfg := &Config{
		PostgresURI: os.Getenv("POSTGRES_URI"),
	}

	if cfg.PostgresURI == "" {
		log.Fatal("POSTGRES_URI no está definida en las variables de entorno")
	}

	return cfg
}
