package persistence

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

type resumenMensualRepository struct {
	db *database.Database
}

func NewResumenMensualRepository(db *database.Database) domain.ResumenMensualRepository {
	return &resumenMensualRepository{db: db}
}

const resumenSelect = `SELECT id_resumen, mes, anio, total_ingresos, total_gastos_fijos, total_gastos_variables, balance, observaciones, fecha_generacion, fecha_actualizacion FROM resumen_mensual`

func (r *resumenMensualRepository) GetByMesAnio(mes, anio int) (*domain.ResumenMensual, error) {
	var rm domain.ResumenMensual
	err := r.db.Pool.QueryRow(context.Background(), resumenSelect+" WHERE mes = $1 AND anio = $2", mes, anio).Scan(&rm.ID, &rm.Mes, &rm.Anio, &rm.TotalIngresos, &rm.TotalGastosFijos, &rm.TotalGastosVariables, &rm.Balance, &rm.Observaciones, &rm.FechaGeneracion, &rm.FechaActualizacion)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.ErrNotFound{Entity: "resumen_mensual", ID: fmt.Sprintf("%d/%d", mes, anio)}
		}
		return nil, err
	}
	return &rm, nil
}

func (r *resumenMensualRepository) GetActual() (*domain.ResumenMensual, error) {
	now := time.Now()
	return r.GetByMesAnio(int(now.Month()), now.Year())
}

func (r *resumenMensualRepository) GetByProductoID(productoID, mes, anio int) (*domain.ResumenProducto, error) {
	rp := &domain.ResumenProducto{IDProducto: productoID, Mes: mes, Anio: anio}
	err := r.db.Pool.QueryRow(context.Background(),
		`SELECT COALESCE(SUM(cantidad), 0), COALESCE(SUM(cantidad * COALESCE(precio_unitario, 0)), 0) FROM entradas_productos WHERE id_producto = $1 AND EXTRACT(MONTH FROM fecha_entrada) = $2 AND EXTRACT(YEAR FROM fecha_entrada) = $3`,
		productoID, mes, anio).Scan(&rp.TotalEntradas, &rp.MontoEntradas)
	if err != nil {
		return nil, err
	}
	err = r.db.Pool.QueryRow(context.Background(),
		`SELECT COALESCE(SUM(cantidad), 0), COALESCE(SUM(total), 0) FROM salidas_productos WHERE id_producto = $1 AND EXTRACT(MONTH FROM fecha_salida) = $2 AND EXTRACT(YEAR FROM fecha_salida) = $3`,
		productoID, mes, anio).Scan(&rp.TotalSalidas, &rp.MontoSalidas)
	if err != nil {
		return nil, err
	}
	return rp, nil
}

func (r *resumenMensualRepository) Generar(mes, anio int) (*domain.ResumenMensual, error) {
	var totalVentas float64
	err := r.db.Pool.QueryRow(context.Background(),
		`SELECT COALESCE(SUM(total), 0) FROM salidas_productos WHERE EXTRACT(MONTH FROM fecha_salida) = $1 AND EXTRACT(YEAR FROM fecha_salida) = $2`, mes, anio).Scan(&totalVentas)
	if err != nil {
		return nil, err
	}
	var totalGastos float64
	err = r.db.Pool.QueryRow(context.Background(),
		`SELECT COALESCE(SUM(monto_salida), 0) FROM control_diario WHERE EXTRACT(MONTH FROM fecha) = $1 AND EXTRACT(YEAR FROM fecha) = $2`, mes, anio).Scan(&totalGastos)
	if err != nil {
		return nil, err
	}
	rm := &domain.ResumenMensual{
		Mes: mes, Anio: anio,
		TotalIngresos: totalVentas, TotalGastosFijos: 0, TotalGastosVariables: totalGastos,
		Observaciones: "Generado automaticamente", FechaGeneracion: time.Now(),
	}
	rm.Balance = rm.TotalIngresos - rm.TotalGastosFijos - rm.TotalGastosVariables
	err = r.Upsert(rm)
	return rm, err
}

func (r *resumenMensualRepository) Upsert(rm *domain.ResumenMensual) error {
	query := `INSERT INTO resumen_mensual (mes, anio, total_ingresos, total_gastos_fijos, total_gastos_variables, observaciones) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (mes, anio) DO UPDATE SET total_ingresos = EXCLUDED.total_ingresos, total_gastos_fijos = EXCLUDED.total_gastos_fijos, total_gastos_variables = EXCLUDED.total_gastos_variables, observaciones = EXCLUDED.observaciones RETURNING id_resumen, balance, fecha_generacion, fecha_actualizacion`
	return r.db.Pool.QueryRow(context.Background(), query, rm.Mes, rm.Anio, rm.TotalIngresos, rm.TotalGastosFijos, rm.TotalGastosVariables, rm.Observaciones).Scan(&rm.ID, &rm.Balance, &rm.FechaGeneracion, &rm.FechaActualizacion)
}
