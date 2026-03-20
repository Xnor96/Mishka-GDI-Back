package persistence

import (
	"context"
	"time"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

type entradaProductoRepository struct {
	db *database.Database
}

func NewEntradaProductoRepository(db *database.Database) domain.EntradaProductoRepository {
	return &entradaProductoRepository{db: db}
}

const entradaSelectJoin = `
	SELECT ep.id_entrada, ep.id_producto, ep.fecha_entrada, ep.cantidad,
	       ep.precio_unitario, ep.observaciones, ep.usuario_registro,
	       ep.fecha_creacion, ep.fecha_actualizacion,
	       p.nombre, p.codigo, COALESCE(c.nombre, '') AS nombre_categoria
	FROM entradas_productos ep
	JOIN productos p ON ep.id_producto = p.id_producto
	LEFT JOIN categorias c ON p.id_categoria = c.id_categoria`

func scanEntradaConProducto(rows pgx.Rows) (domain.EntradaConProducto, error) {
	var e domain.EntradaConProducto
	err := rows.Scan(
		&e.ID, &e.IDProducto, &e.FechaEntrada, &e.Cantidad,
		&e.PrecioUnitario, &e.Observaciones, &e.UsuarioRegistro,
		&e.FechaCreacion, &e.FechaActualizacion,
		&e.NombreProducto, &e.CodigoProducto, &e.NombreCategoria,
	)
	return e, err
}

func (r *entradaProductoRepository) GetAll() ([]domain.EntradaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), entradaSelectJoin+" ORDER BY ep.fecha_entrada DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entradas []domain.EntradaConProducto
	for rows.Next() {
		e, err := scanEntradaConProducto(rows)
		if err != nil {
			return nil, err
		}
		entradas = append(entradas, e)
	}
	return entradas, nil
}

func (r *entradaProductoRepository) GetByID(id int) (*domain.EntradaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), entradaSelectJoin+" WHERE ep.id_entrada = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, &domain.ErrNotFound{Entity: "entrada", ID: id}
	}
	e, err := scanEntradaConProducto(rows)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *entradaProductoRepository) GetByProductoID(productoID int) ([]domain.EntradaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), entradaSelectJoin+" WHERE ep.id_producto = $1 ORDER BY ep.fecha_entrada DESC", productoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entradas []domain.EntradaConProducto
	for rows.Next() {
		e, err := scanEntradaConProducto(rows)
		if err != nil {
			return nil, err
		}
		entradas = append(entradas, e)
	}
	return entradas, nil
}

func (r *entradaProductoRepository) GetByFecha(fecha string) ([]domain.EntradaConProducto, error) {
	rows, err := r.db.Pool.Query(context.Background(), entradaSelectJoin+" WHERE ep.fecha_entrada = $1 ORDER BY ep.fecha_creacion DESC", fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entradas []domain.EntradaConProducto
	for rows.Next() {
		e, err := scanEntradaConProducto(rows)
		if err != nil {
			return nil, err
		}
		entradas = append(entradas, e)
	}
	return entradas, nil
}

func (r *entradaProductoRepository) Create(entrada *domain.EntradaProducto) error {
	fechaEntrada, err := time.Parse("2006-01-02", entrada.FechaEntrada.Format("2006-01-02"))
	if err != nil {
		return err
	}
	query := `INSERT INTO entradas_productos (id_producto, fecha_entrada, cantidad, precio_unitario, observaciones, usuario_registro) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_entrada, fecha_creacion, fecha_actualizacion`
	err = r.db.Pool.QueryRow(context.Background(), query, entrada.IDProducto, fechaEntrada, entrada.Cantidad, entrada.PrecioUnitario, entrada.Observaciones, entrada.UsuarioRegistro).Scan(&entrada.ID, &entrada.FechaCreacion, &entrada.FechaActualizacion)
	if err != nil {
		return err
	}
	entrada.FechaEntrada = fechaEntrada
	return nil
}
