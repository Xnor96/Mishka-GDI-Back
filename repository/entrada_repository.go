package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Mishka-GDI-Back/db"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/jackc/pgx/v5"
)

type entradaProductoRepository struct {
	db *db.Database
}

func NewEntradaProductoRepository(database *db.Database) domain.EntradaProductoRepository {
	return &entradaProductoRepository{
		db: database,
	}
}

func (r *entradaProductoRepository) GetAll() ([]domain.EntradaProducto, error) {
	query := `SELECT id_entrada, id_producto, fecha_entrada, cantidad, precio_unitario, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM entradas_productos ORDER BY fecha_entrada DESC`

	rows, err := r.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entradas []domain.EntradaProducto
	for rows.Next() {
		var entrada domain.EntradaProducto
		err := rows.Scan(
			&entrada.ID,
			&entrada.IDProducto,
			&entrada.FechaEntrada,
			&entrada.Cantidad,
			&entrada.PrecioUnitario,
			&entrada.Observaciones,
			&entrada.UsuarioRegistro,
			&entrada.FechaCreacion,
			&entrada.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		entradas = append(entradas, entrada)
	}

	return entradas, nil
}

func (r *entradaProductoRepository) GetByID(id int) (*domain.EntradaProducto, error) {
	query := `SELECT id_entrada, id_producto, fecha_entrada, cantidad, precio_unitario, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM entradas_productos WHERE id_entrada = $1`

	var entrada domain.EntradaProducto
	err := r.db.Pool.QueryRow(context.Background(), query, id).Scan(
		&entrada.ID,
		&entrada.IDProducto,
		&entrada.FechaEntrada,
		&entrada.Cantidad,
		&entrada.PrecioUnitario,
		&entrada.Observaciones,
		&entrada.UsuarioRegistro,
		&entrada.FechaCreacion,
		&entrada.FechaActualizacion,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("entrada no encontrada")
		}
		return nil, err
	}

	return &entrada, nil
}

func (r *entradaProductoRepository) GetByProductoID(productoID int) ([]domain.EntradaProducto, error) {
	query := `SELECT id_entrada, id_producto, fecha_entrada, cantidad, precio_unitario, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM entradas_productos 
			  WHERE id_producto = $1 
			  ORDER BY fecha_entrada DESC`

	rows, err := r.db.Pool.Query(context.Background(), query, productoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entradas []domain.EntradaProducto
	for rows.Next() {
		var entrada domain.EntradaProducto
		err := rows.Scan(
			&entrada.ID,
			&entrada.IDProducto,
			&entrada.FechaEntrada,
			&entrada.Cantidad,
			&entrada.PrecioUnitario,
			&entrada.Observaciones,
			&entrada.UsuarioRegistro,
			&entrada.FechaCreacion,
			&entrada.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		entradas = append(entradas, entrada)
	}

	return entradas, nil
}

func (r *entradaProductoRepository) GetByFecha(fecha string) ([]domain.EntradaProducto, error) {
	query := `SELECT id_entrada, id_producto, fecha_entrada, cantidad, precio_unitario, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM entradas_productos 
			  WHERE fecha_entrada = $1 
			  ORDER BY fecha_creacion DESC`

	rows, err := r.db.Pool.Query(context.Background(), query, fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entradas []domain.EntradaProducto
	for rows.Next() {
		var entrada domain.EntradaProducto
		err := rows.Scan(
			&entrada.ID,
			&entrada.IDProducto,
			&entrada.FechaEntrada,
			&entrada.Cantidad,
			&entrada.PrecioUnitario,
			&entrada.Observaciones,
			&entrada.UsuarioRegistro,
			&entrada.FechaCreacion,
			&entrada.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		entradas = append(entradas, entrada)
	}

	return entradas, nil
}

func (r *entradaProductoRepository) Create(entrada *domain.EntradaProducto) error {
	// Parsear la fecha de entrada
	fechaEntrada, err := time.Parse("2006-01-02", entrada.FechaEntrada.Format("2006-01-02"))
	if err != nil {
		return err
	}

	query := `INSERT INTO entradas_productos (id_producto, fecha_entrada, cantidad, precio_unitario, 
			                                 observaciones, usuario_registro) 
			  VALUES ($1, $2, $3, $4, $5, $6) 
			  RETURNING id_entrada, fecha_creacion, fecha_actualizacion`

	err = r.db.Pool.QueryRow(context.Background(), query,
		entrada.IDProducto,
		fechaEntrada,
		entrada.Cantidad,
		entrada.PrecioUnitario,
		entrada.Observaciones,
		entrada.UsuarioRegistro,
	).Scan(
		&entrada.ID,
		&entrada.FechaCreacion,
		&entrada.FechaActualizacion,
	)

	if err != nil {
		return err
	}

	entrada.FechaEntrada = fechaEntrada
	return nil
}
