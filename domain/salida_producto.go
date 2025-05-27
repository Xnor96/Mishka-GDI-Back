package domain

import "time"

type SalidaProducto struct {
	ID                 int       `db:"id_salida"`
	IDProducto         int       `db:"id_producto"`
	FechaSalida        time.Time `db:"fecha_salida"`
	Cantidad           int       `db:"cantidad"`
	Observaciones      string    `db:"observaciones"`
	UsuarioRegistro    string    `db:"usuario_registro"`
	FechaCreacion      time.Time `db:"fecha_creacion"`
	FechaActualizacion time.Time `db:"fecha_actualizacion"`
}
