package domain

import "time"

type Categoria struct {
	ID                 int       `db:"id_categoria"`
	Nombre             string    `db:"nombre"`
	FechaCreacion      time.Time `db:"fecha_creacion"`
	FechaActualizacion time.Time `db:"fecha_actualizacion"`
}
