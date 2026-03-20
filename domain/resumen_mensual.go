package domain

import "time"

type ResumenMensual struct {
	ID                   int
	Mes                  int
	Anio                 int
	TotalIngresos        float64
	TotalGastosFijos     float64
	TotalGastosVariables float64
	Balance              float64
	Observaciones        string
	FechaGeneracion      time.Time
	FechaActualizacion   time.Time
}

// ResumenProducto es el resumen de movimientos de un producto en un mes/año
type ResumenProducto struct {
	IDProducto    int
	Mes           int
	Anio          int
	TotalEntradas int
	MontoEntradas float64
	TotalSalidas  int
	MontoSalidas  float64
}
