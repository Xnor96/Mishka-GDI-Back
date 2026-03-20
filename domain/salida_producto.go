package domain

import "time"

type SalidaProducto struct {
	ID                 int
	IDProducto         int
	FechaSalida        time.Time
	Cantidad           int
	PrecioVenta        float64
	Descuento          float64
	Total              float64
	LugarVenta         string
	TipoPago           string
	Observaciones      string
	UsuarioRegistro    string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

// SalidaConProducto es el modelo de lectura enriquecido con datos del producto
type SalidaConProducto struct {
	SalidaProducto
	NombreProducto  string
	CodigoProducto  string
	NombreCategoria string
}
