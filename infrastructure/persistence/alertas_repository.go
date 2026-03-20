package persistence

import (
	"context"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
)

type alertasRepository struct {
	db *database.Database
}

func NewAlertasRepository(db *database.Database) domain.AlertasRepository {
	return &alertasRepository{db: db}
}

func (r *alertasRepository) GetStockBajo(limite int) ([]domain.AlertaStockBajo, error) {
	query := `SELECT p.id_producto, p.codigo, p.nombre, COALESCE(c.nombre, 'SIN CATEGORIA'), p.stock_actual, p.stock_inicial, p.precio_unitario FROM productos p LEFT JOIN categorias c ON p.id_categoria = c.id_categoria WHERE p.stock_actual <= $1 ORDER BY p.stock_actual ASC, c.nombre, p.nombre`
	rows, err := r.db.Pool.Query(context.Background(), query, limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.AlertaStockBajo
	for rows.Next() {
		var item domain.AlertaStockBajo
		if err := rows.Scan(&item.IDProducto, &item.Codigo, &item.Nombre, &item.Categoria, &item.StockActual, &item.StockInicial, &item.PrecioUnitario); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
