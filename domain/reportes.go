package domain

import "time"

type ReporteInventarioItem struct {
	IDProducto     int
	Codigo         string
	Nombre         string
	Categoria      string
	UnidadMedida   string
	PrecioUnitario float64
	StockActual    int
	ValorTotal     float64
}

type ReporteMovimiento struct {
	Fecha      time.Time
	Tipo       string
	Codigo     string
	Nombre     string
	Categoria  string
	Cantidad   int
	Precio     float64
	Total      float64
	LugarVenta string
	TipoPago   string
}

type ReporteProductoVendido struct {
	IDProducto    int
	Codigo        string
	Nombre        string
	Categoria     string
	TotalVendido  int
	TotalIngresos float64
}

type ReporteValoracion struct {
	IDCategoria     int
	NombreCategoria string
	TotalProductos  int
	TotalUnidades   int
	ValorTotal      float64
}

type AlertaStockBajo struct {
	IDProducto     int
	Codigo         string
	Nombre         string
	Categoria      string
	StockActual    int
	StockInicial   int
	PrecioUnitario float64
}
