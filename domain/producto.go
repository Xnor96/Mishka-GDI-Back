package domain

import "time"

type Producto struct {
	ID                 int
	Codigo             string
	Nombre             string
	IDCategoria        *int
	UnidadMedida       string
	PrecioUnitario     float64
	StockActual        int
	StockInicial       int
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}
