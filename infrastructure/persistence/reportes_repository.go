package persistence

import (
	"context"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
)

type reportesRepository struct {
	db *database.Database
}

func NewReportesRepository(db *database.Database) domain.ReportesRepository {
	return &reportesRepository{db: db}
}

func (r *reportesRepository) GetInventarioActual() ([]domain.ReporteInventarioItem, error) {
	query := `SELECT p.id_producto, p.codigo, p.nombre, COALESCE(c.nombre, 'SIN CATEGORIA'), p.unidad_medida, p.precio_unitario, p.stock_actual, p.stock_actual * p.precio_unitario AS valor_total FROM productos p LEFT JOIN categorias c ON p.id_categoria = c.id_categoria WHERE p.stock_actual > 0 ORDER BY c.nombre, p.nombre`
	rows, err := r.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ReporteInventarioItem
	for rows.Next() {
		var item domain.ReporteInventarioItem
		if err := rows.Scan(&item.IDProducto, &item.Codigo, &item.Nombre, &item.Categoria, &item.UnidadMedida, &item.PrecioUnitario, &item.StockActual, &item.ValorTotal); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *reportesRepository) GetMovimientos(inicio, fin string) ([]domain.ReporteMovimiento, error) {
	query := `SELECT ep.fecha_entrada, 'ENTRADA', p.codigo, p.nombre, COALESCE(c.nombre,''), ep.cantidad, COALESCE(ep.precio_unitario,0), ep.cantidad * COALESCE(ep.precio_unitario,0), '', '' FROM entradas_productos ep JOIN productos p ON ep.id_producto = p.id_producto LEFT JOIN categorias c ON p.id_categoria = c.id_categoria WHERE ep.fecha_entrada BETWEEN $1 AND $2 UNION ALL SELECT sp.fecha_salida, 'SALIDA', p.codigo, p.nombre, COALESCE(c.nombre,''), sp.cantidad, sp.precio_venta, sp.total, sp.lugar_venta, sp.tipo_pago FROM salidas_productos sp JOIN productos p ON sp.id_producto = p.id_producto LEFT JOIN categorias c ON p.id_categoria = c.id_categoria WHERE sp.fecha_salida BETWEEN $1 AND $2 ORDER BY 1 DESC`
	rows, err := r.db.Pool.Query(context.Background(), query, inicio, fin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ReporteMovimiento
	for rows.Next() {
		var item domain.ReporteMovimiento
		if err := rows.Scan(&item.Fecha, &item.Tipo, &item.Codigo, &item.Nombre, &item.Categoria, &item.Cantidad, &item.Precio, &item.Total, &item.LugarVenta, &item.TipoPago); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *reportesRepository) GetProductosMasVendidos(limite int) ([]domain.ReporteProductoVendido, error) {
	query := `SELECT p.id_producto, p.codigo, p.nombre, COALESCE(c.nombre,''), SUM(sp.cantidad) AS total_vendido, SUM(sp.total) AS total_ingresos FROM salidas_productos sp JOIN productos p ON sp.id_producto = p.id_producto LEFT JOIN categorias c ON p.id_categoria = c.id_categoria GROUP BY p.id_producto, p.codigo, p.nombre, c.nombre ORDER BY total_vendido DESC LIMIT $1`
	rows, err := r.db.Pool.Query(context.Background(), query, limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ReporteProductoVendido
	for rows.Next() {
		var item domain.ReporteProductoVendido
		if err := rows.Scan(&item.IDProducto, &item.Codigo, &item.Nombre, &item.Categoria, &item.TotalVendido, &item.TotalIngresos); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *reportesRepository) GetProductosMasIngresados(limite int) ([]domain.ReporteProductoVendido, error) {
	query := `SELECT p.id_producto, p.codigo, p.nombre, COALESCE(c.nombre,''), SUM(ep.cantidad) AS total_ingresado, SUM(ep.cantidad * COALESCE(ep.precio_unitario,0)) AS total_costo FROM entradas_productos ep JOIN productos p ON ep.id_producto = p.id_producto LEFT JOIN categorias c ON p.id_categoria = c.id_categoria GROUP BY p.id_producto, p.codigo, p.nombre, c.nombre ORDER BY total_ingresado DESC LIMIT $1`
	rows, err := r.db.Pool.Query(context.Background(), query, limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ReporteProductoVendido
	for rows.Next() {
		var item domain.ReporteProductoVendido
		if err := rows.Scan(&item.IDProducto, &item.Codigo, &item.Nombre, &item.Categoria, &item.TotalVendido, &item.TotalIngresos); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *reportesRepository) GetValoracionInventario() ([]domain.ReporteValoracion, error) {
	query := `SELECT c.id_categoria, COALESCE(c.nombre, 'SIN CATEGORIA'), COUNT(p.id_producto), COALESCE(SUM(p.stock_actual), 0), COALESCE(SUM(p.stock_actual * p.precio_unitario), 0) FROM categorias c LEFT JOIN productos p ON p.id_categoria = c.id_categoria GROUP BY c.id_categoria, c.nombre ORDER BY SUM(p.stock_actual * p.precio_unitario) DESC NULLS LAST`
	rows, err := r.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.ReporteValoracion
	for rows.Next() {
		var item domain.ReporteValoracion
		if err := rows.Scan(&item.IDCategoria, &item.NombreCategoria, &item.TotalProductos, &item.TotalUnidades, &item.ValorTotal); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
