package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Mishka-GDI-Back/db"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/jackc/pgx/v5"
)

type salidaProductoRepository struct {
	db *db.Database
}

func NewSalidaProductoRepository(database *db.Database) domain.SalidaProductoRepository {
	return &salidaProductoRepository{
		db: database,
	}
}

func (r *salidaProductoRepository) GetAll() ([]domain.SalidaProducto, error) {
	query := `SELECT id_salida, id_producto, fecha_salida, cantidad, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM salidas_productos ORDER BY fecha_salida DESC`

	rows, err := r.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salidas []domain.SalidaProducto
	for rows.Next() {
		var salida domain.SalidaProducto
		err := rows.Scan(
			&salida.ID,
			&salida.IDProducto,
			&salida.FechaSalida,
			&salida.Cantidad,
			&salida.Observaciones,
			&salida.UsuarioRegistro,
			&salida.FechaCreacion,
			&salida.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		salidas = append(salidas, salida)
	}

	return salidas, nil
}

func (r *salidaProductoRepository) GetByID(id int) (*domain.SalidaProducto, error) {
	query := `SELECT id_salida, id_producto, fecha_salida, cantidad, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM salidas_productos WHERE id_salida = $1`

	var salida domain.SalidaProducto
	err := r.db.Pool.QueryRow(context.Background(), query, id).Scan(
		&salida.ID,
		&salida.IDProducto,
		&salida.FechaSalida,
		&salida.Cantidad,
		&salida.Observaciones,
		&salida.UsuarioRegistro,
		&salida.FechaCreacion,
		&salida.FechaActualizacion,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("salida no encontrada")
		}
		return nil, err
	}

	return &salida, nil
}

func (r *salidaProductoRepository) GetByProductoID(productoID int) ([]domain.SalidaProducto, error) {
	query := `SELECT id_salida, id_producto, fecha_salida, cantidad, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM salidas_productos 
			  WHERE id_producto = $1 
			  ORDER BY fecha_salida DESC`

	rows, err := r.db.Pool.Query(context.Background(), query, productoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salidas []domain.SalidaProducto
	for rows.Next() {
		var salida domain.SalidaProducto
		err := rows.Scan(
			&salida.ID,
			&salida.IDProducto,
			&salida.FechaSalida,
			&salida.Cantidad,
			&salida.Observaciones,
			&salida.UsuarioRegistro,
			&salida.FechaCreacion,
			&salida.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		salidas = append(salidas, salida)
	}

	return salidas, nil
}

func (r *salidaProductoRepository) GetByFecha(fecha string) ([]domain.SalidaProducto, error) {
	query := `SELECT id_salida, id_producto, fecha_salida, cantidad, 
			         observaciones, usuario_registro, fecha_creacion, fecha_actualizacion 
			  FROM salidas_productos 
			  WHERE fecha_salida = $1 
			  ORDER BY fecha_creacion DESC`

	rows, err := r.db.Pool.Query(context.Background(), query, fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salidas []domain.SalidaProducto
	for rows.Next() {
		var salida domain.SalidaProducto
		err := rows.Scan(
			&salida.ID,
			&salida.IDProducto,
			&salida.FechaSalida,
			&salida.Cantidad,
			&salida.Observaciones,
			&salida.UsuarioRegistro,
			&salida.FechaCreacion,
			&salida.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		salidas = append(salidas, salida)
	}

	return salidas, nil
}

func (r *salidaProductoRepository) Create(salida *domain.SalidaProducto) error {
	// Parsear la fecha de salida
	fechaSalida, err := time.Parse("2006-01-02", salida.FechaSalida.Format("2006-01-02"))
	if err != nil {
		return err
	}

	query := `INSERT INTO salidas_productos (id_producto, fecha_salida, cantidad, 
			                                observaciones, usuario_registro) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING id_salida, fecha_creacion, fecha_actualizacion`

	err = r.db.Pool.QueryRow(context.Background(), query,
		salida.IDProducto,
		fechaSalida,
		salida.Cantidad,
		salida.Observaciones,
		salida.UsuarioRegistro,
	).Scan(
		&salida.ID,
		&salida.FechaCreacion,
		&salida.FechaActualizacion,
	)

	if err != nil {
		return err
	}

	salida.FechaSalida = fechaSalida
	return nil
}
