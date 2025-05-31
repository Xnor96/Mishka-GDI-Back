package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Mishka-GDI-Back/db"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/jackc/pgx/v5"
)

type productoRepository struct {
	db *db.Database
}

func NewProductoRepository(database *db.Database) domain.ProductoRepository {
	return &productoRepository{
		db: database,
	}
}

func (r *productoRepository) GetAll() ([]domain.Producto, error) {
	query := `SELECT id_producto, codigo, nombre, id_categoria, unidad_medida, 
			         precio_unitario, stock_actual, stock_inicial, 
			         fecha_creacion, fecha_actualizacion 
			  FROM productos ORDER BY nombre`

	rows, err := r.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []domain.Producto
	for rows.Next() {
		var producto domain.Producto
		err := rows.Scan(
			&producto.ID,
			&producto.Codigo,
			&producto.Nombre,
			&producto.IDCategoria,
			&producto.UnidadMedida,
			&producto.PrecioUnitario,
			&producto.StockActual,
			&producto.StockInicial,
			&producto.FechaCreacion,
			&producto.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		productos = append(productos, producto)
	}

	return productos, nil
}

func (r *productoRepository) GetByID(id int) (*domain.Producto, error) {
	query := `SELECT id_producto, codigo, nombre, id_categoria, unidad_medida, 
			         precio_unitario, stock_actual, stock_inicial, 
			         fecha_creacion, fecha_actualizacion 
			  FROM productos WHERE id_producto = $1`

	var producto domain.Producto
	err := r.db.Pool.QueryRow(context.Background(), query, id).Scan(
		&producto.ID,
		&producto.Codigo,
		&producto.Nombre,
		&producto.IDCategoria,
		&producto.UnidadMedida,
		&producto.PrecioUnitario,
		&producto.StockActual,
		&producto.StockInicial,
		&producto.FechaCreacion,
		&producto.FechaActualizacion,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("producto no encontrado")
		}
		return nil, err
	}

	return &producto, nil
}

func (r *productoRepository) GetByCodigo(codigo string) (*domain.Producto, error) {
	query := `SELECT id_producto, codigo, nombre, id_categoria, unidad_medida, 
			         precio_unitario, stock_actual, stock_inicial, 
			         fecha_creacion, fecha_actualizacion 
			  FROM productos WHERE codigo = $1`

	var producto domain.Producto
	err := r.db.Pool.QueryRow(context.Background(), query, codigo).Scan(
		&producto.ID,
		&producto.Codigo,
		&producto.Nombre,
		&producto.IDCategoria,
		&producto.UnidadMedida,
		&producto.PrecioUnitario,
		&producto.StockActual,
		&producto.StockInicial,
		&producto.FechaCreacion,
		&producto.FechaActualizacion,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("producto no encontrado")
		}
		return nil, err
	}

	return &producto, nil
}

func (r *productoRepository) Create(producto *domain.Producto) error {
	query := `INSERT INTO productos (codigo, nombre, id_categoria, unidad_medida, 
			                        precio_unitario, stock_actual, stock_inicial) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) 
			  RETURNING id_producto, fecha_creacion, fecha_actualizacion`

	err := r.db.Pool.QueryRow(context.Background(), query,
		producto.Codigo,
		producto.Nombre,
		producto.IDCategoria,
		producto.UnidadMedida,
		producto.PrecioUnitario,
		producto.StockActual,
		producto.StockInicial,
	).Scan(
		&producto.ID,
		&producto.FechaCreacion,
		&producto.FechaActualizacion,
	)

	return err
}

func (r *productoRepository) Update(producto *domain.Producto) error {
	query := `UPDATE productos 
			  SET codigo = $2, nombre = $3, id_categoria = $4, unidad_medida = $5,
			      precio_unitario = $6, stock_actual = $7, stock_inicial = $8
			  WHERE id_producto = $1
			  RETURNING fecha_actualizacion`

	err := r.db.Pool.QueryRow(context.Background(), query,
		producto.ID,
		producto.Codigo,
		producto.Nombre,
		producto.IDCategoria,
		producto.UnidadMedida,
		producto.PrecioUnitario,
		producto.StockActual,
		producto.StockInicial,
	).Scan(&producto.FechaActualizacion)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("producto no encontrado")
		}
		return err
	}

	return nil
}

func (r *productoRepository) Delete(id int) error {
	query := `DELETE FROM productos WHERE id_producto = $1`

	result, err := r.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("producto no encontrado")
	}

	return nil
}

func (r *productoRepository) GetStockBajo(limite int) ([]domain.Producto, error) {
	query := `SELECT id_producto, codigo, nombre, id_categoria, unidad_medida, 
			         precio_unitario, stock_actual, stock_inicial, 
			         fecha_creacion, fecha_actualizacion 
			  FROM productos 
			  WHERE stock_actual <= $1 
			  ORDER BY stock_actual ASC`

	rows, err := r.db.Pool.Query(context.Background(), query, limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []domain.Producto
	for rows.Next() {
		var producto domain.Producto
		err := rows.Scan(
			&producto.ID,
			&producto.Codigo,
			&producto.Nombre,
			&producto.IDCategoria,
			&producto.UnidadMedida,
			&producto.PrecioUnitario,
			&producto.StockActual,
			&producto.StockInicial,
			&producto.FechaCreacion,
			&producto.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		productos = append(productos, producto)
	}

	return productos, nil
}

func (r *productoRepository) Search(termino string) ([]domain.Producto, error) {
	termino = strings.ToLower(strings.TrimSpace(termino))
	if termino == "" {
		return r.GetAll()
	}

	query := `SELECT id_producto, codigo, nombre, id_categoria, unidad_medida, 
			         precio_unitario, stock_actual, stock_inicial, 
			         fecha_creacion, fecha_actualizacion 
			  FROM productos 
			  WHERE LOWER(codigo) LIKE $1 OR LOWER(nombre) LIKE $1
			  ORDER BY nombre`

	searchTerm := fmt.Sprintf("%%%s%%", termino)

	rows, err := r.db.Pool.Query(context.Background(), query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []domain.Producto
	for rows.Next() {
		var producto domain.Producto
		err := rows.Scan(
			&producto.ID,
			&producto.Codigo,
			&producto.Nombre,
			&producto.IDCategoria,
			&producto.UnidadMedida,
			&producto.PrecioUnitario,
			&producto.StockActual,
			&producto.StockInicial,
			&producto.FechaCreacion,
			&producto.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		productos = append(productos, producto)
	}

	return productos, nil
}
