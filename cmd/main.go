package main

import (
	"log"

	"github.com/Mishka-GDI-Back/config"
)

func main() {
	cfg := config.NewConfig()
	log.Printf("Conectando a base de datos en: %s", cfg.PostgresURI)
}
