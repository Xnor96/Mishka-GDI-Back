package domain

import "time"

type ControlDiario struct {
	ID                 int
	Fecha              time.Time
	Descripcion        string
	MontoEntrada       float64
	MontoSalida        float64
	Observaciones      string
	EsVerbena          bool
	UsuarioRegistro    string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}
