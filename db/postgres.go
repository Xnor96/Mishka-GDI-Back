package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewPostgresConnection(uri string) (*Database, error) {
	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Verificar la conexión
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("✅ Conexión exitosa a PostgreSQL")
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		log.Println("Conexión a PostgreSQL cerrada")
	}
}
