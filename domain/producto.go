package domain

import "time"

type Producto struct {
	ID                 int       `db:"id_producto"`
	Codigo             string    `db:"codigo"`
	Nombre             string    `db:"nombre"`
	IDCategoria        *int      `db:"id_categoria"` // Puede ser NULL
	UnidadMedida       string    `db:"unidad_medida"`
	PrecioUnitario     float64   `db:"precio_unitario"`
	StockActual        int       `db:"stock_actual"`
	StockInicial       int       `db:"stock_inicial"`
	FechaCreacion      time.Time `db:"fecha_creacion"`
	FechaActualizacion time.Time `db:"fecha_actualizacion"`
}
