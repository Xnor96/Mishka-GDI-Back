package domain

import "time"

type EntradaProducto struct {
	ID                 int       `db:"id_entrada"`
	IDProducto         int       `db:"id_producto"`
	FechaEntrada       time.Time `db:"fecha_entrada"`
	Cantidad           int       `db:"cantidad"`
	PrecioUnitario     *float64  `db:"precio_unitario"` // Puede ser NULL
	Observaciones      string    `db:"observaciones"`
	UsuarioRegistro    string    `db:"usuario_registro"`
	FechaCreacion      time.Time `db:"fecha_creacion"`
	FechaActualizacion time.Time `db:"fecha_actualizacion"`
}
