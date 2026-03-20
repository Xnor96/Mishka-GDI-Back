package domain

import "time"

type Categoria struct {
	ID                 int
	Nombre             string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}
