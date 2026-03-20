package domain

import "time"

type EntradaProducto struct {
	ID                 int
	IDProducto         int
	FechaEntrada       time.Time
	Cantidad           int
	PrecioUnitario     *float64
	Observaciones      string
	UsuarioRegistro    string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

// EntradaConProducto es el modelo de lectura enriquecido con datos del producto
type EntradaConProducto struct {
	EntradaProducto
	NombreProducto  string
	CodigoProducto  string
	NombreCategoria string
}
