package config

import (
	"log"
	"os"
)

type Config struct {
	PostgresURI string
	Port        string
	GinMode     string
}

func NewConfig() *Config {
	cfg := &Config{
		PostgresURI: os.Getenv("POSTGRES_URI"),
		Port:        os.Getenv("PORT"),
		GinMode:     os.Getenv("GIN_MODE"),
	}

	if cfg.PostgresURI == "" {
		log.Fatal("POSTGRES_URI no está definida en las variables de entorno")
	}

	if cfg.Port == "" {
		cfg.Port = "8080" // Puerto por defecto
	}

	if cfg.GinMode == "" {
		cfg.GinMode = "debug" // Modo por defecto
	}

	return cfg
}
