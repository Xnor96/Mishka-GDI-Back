package persistence

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

type productoRepository struct {
	db *database.Database
}

func NewProductoRepository(db *database.Database) domain.ProductoRepository {
	return &productoRepository{db: db}
}

const productoSelect = `SELECT id_producto, codigo, nombre, id_categoria, unidad_medida, precio_unitario, stock_actual, stock_inicial, fecha_creacion, fecha_actualizacion FROM productos`

func scanProducto(row interface{ Scan(dest ...any) error }) (domain.Producto, error) {
	var p domain.Producto
	err := row.Scan(&p.ID, &p.Codigo, &p.Nombre, &p.IDCategoria, &p.UnidadMedida, &p.PrecioUnitario, &p.StockActual, &p.StockInicial, &p.FechaCreacion, &p.FechaActualizacion)
	return p, err
}

func (r *productoRepository) GetAll() ([]domain.Producto, error) {
	rows, err := r.db.Pool.Query(context.Background(), productoSelect+" ORDER BY nombre")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productos []domain.Producto
	for rows.Next() {
		p, err := scanProducto(rows)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}

func (r *productoRepository) GetByID(id int) (*domain.Producto, error) {
	row := r.db.Pool.QueryRow(context.Background(), productoSelect+" WHERE id_producto = $1", id)
	p, err := scanProducto(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.ErrNotFound{Entity: "producto", ID: id}
		}
		return nil, err
	}
	return &p, nil
}

func (r *productoRepository) GetByCodigo(codigo string) (*domain.Producto, error) {
	row := r.db.Pool.QueryRow(context.Background(), productoSelect+" WHERE codigo = $1", codigo)
	p, err := scanProducto(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.ErrNotFound{Entity: "producto", ID: codigo}
		}
		return nil, err
	}
	return &p, nil
}

func (r *productoRepository) Create(producto *domain.Producto) error {
	query := `INSERT INTO productos (codigo, nombre, id_categoria, unidad_medida, precio_unitario, stock_actual, stock_inicial) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_producto, fecha_creacion, fecha_actualizacion`
	return r.db.Pool.QueryRow(context.Background(), query, producto.Codigo, producto.Nombre, producto.IDCategoria, producto.UnidadMedida, producto.PrecioUnitario, producto.StockActual, producto.StockInicial).Scan(&producto.ID, &producto.FechaCreacion, &producto.FechaActualizacion)
}

func (r *productoRepository) Update(producto *domain.Producto) error {
	query := `UPDATE productos SET codigo = $2, nombre = $3, id_categoria = $4, unidad_medida = $5, precio_unitario = $6, stock_actual = $7, stock_inicial = $8 WHERE id_producto = $1 RETURNING fecha_actualizacion`
	err := r.db.Pool.QueryRow(context.Background(), query, producto.ID, producto.Codigo, producto.Nombre, producto.IDCategoria, producto.UnidadMedida, producto.PrecioUnitario, producto.StockActual, producto.StockInicial).Scan(&producto.FechaActualizacion)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &domain.ErrNotFound{Entity: "producto", ID: producto.ID}
		}
		return err
	}
	return nil
}

func (r *productoRepository) Delete(id int) error {
	result, err := r.db.Pool.Exec(context.Background(), `DELETE FROM productos WHERE id_producto = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return &domain.ErrNotFound{Entity: "producto", ID: id}
	}
	return nil
}

func (r *productoRepository) GetStockBajo(limite int) ([]domain.Producto, error) {
	rows, err := r.db.Pool.Query(context.Background(), productoSelect+" WHERE stock_actual <= $1 ORDER BY stock_actual ASC", limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productos []domain.Producto
	for rows.Next() {
		p, err := scanProducto(rows)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}

func (r *productoRepository) Search(termino string) ([]domain.Producto, error) {
	termino = strings.ToLower(strings.TrimSpace(termino))
	if termino == "" {
		return r.GetAll()
	}
	searchTerm := fmt.Sprintf("%%%s%%", termino)
	rows, err := r.db.Pool.Query(context.Background(), productoSelect+" WHERE LOWER(codigo) LIKE $1 OR LOWER(nombre) LIKE $1 ORDER BY nombre", searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productos []domain.Producto
	for rows.Next() {
		p, err := scanProducto(rows)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}
