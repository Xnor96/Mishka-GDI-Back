package repository

import (
	"context"
	"errors"

	"github.com/Mishka-GDI-Back/db"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/jackc/pgx/v5"
)

type categoriaRepository struct {
	db *db.Database
}

func NewCategoriaRepository(database *db.Database) domain.CategoriaRepository {
	return &categoriaRepository{
		db: database,
	}
}

func (r *categoriaRepository) GetAll() ([]domain.Categoria, error) {
	query := `SELECT id_categoria, nombre, fecha_creacion, fecha_actualizacion 
			  FROM categorias ORDER BY nombre`

	rows, err := r.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []domain.Categoria
	for rows.Next() {
		var categoria domain.Categoria
		err := rows.Scan(
			&categoria.ID,
			&categoria.Nombre,
			&categoria.FechaCreacion,
			&categoria.FechaActualizacion,
		)
		if err != nil {
			return nil, err
		}
		categorias = append(categorias, categoria)
	}

	return categorias, nil
}

func (r *categoriaRepository) GetByID(id int) (*domain.Categoria, error) {
	query := `SELECT id_categoria, nombre, fecha_creacion, fecha_actualizacion 
			  FROM categorias WHERE id_categoria = $1`

	var categoria domain.Categoria
	err := r.db.Pool.QueryRow(context.Background(), query, id).Scan(
		&categoria.ID,
		&categoria.Nombre,
		&categoria.FechaCreacion,
		&categoria.FechaActualizacion,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("categoría no encontrada")
		}
		return nil, err
	}

	return &categoria, nil
}

func (r *categoriaRepository) Create(categoria *domain.Categoria) error {
	query := `INSERT INTO categorias (nombre) 
			  VALUES ($1) 
			  RETURNING id_categoria, fecha_creacion, fecha_actualizacion`

	err := r.db.Pool.QueryRow(context.Background(), query, categoria.Nombre).Scan(
		&categoria.ID,
		&categoria.FechaCreacion,
		&categoria.FechaActualizacion,
	)

	return err
}

func (r *categoriaRepository) Update(categoria *domain.Categoria) error {
	query := `UPDATE categorias 
			  SET nombre = $2
			  WHERE id_categoria = $1
			  RETURNING fecha_actualizacion`

	err := r.db.Pool.QueryRow(context.Background(), query, categoria.ID, categoria.Nombre).Scan(
		&categoria.FechaActualizacion,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("categoría no encontrada")
		}
		return err
	}

	return nil
}

func (r *categoriaRepository) Delete(id int) error {
	query := `DELETE FROM categorias WHERE id_categoria = $1`

	result, err := r.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("categoría no encontrada")
	}

	return nil
}
