package persistence

import (
	"context"
	"errors"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

type categoriaRepository struct {
	db *database.Database
}

func NewCategoriaRepository(db *database.Database) domain.CategoriaRepository {
	return &categoriaRepository{db: db}
}

func (r *categoriaRepository) GetAll() ([]domain.Categoria, error) {
	query := `SELECT id_categoria, nombre, fecha_creacion, fecha_actualizacion FROM categorias ORDER BY nombre`
	rows, err := r.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categorias []domain.Categoria
	for rows.Next() {
		var c domain.Categoria
		if err := rows.Scan(&c.ID, &c.Nombre, &c.FechaCreacion, &c.FechaActualizacion); err != nil {
			return nil, err
		}
		categorias = append(categorias, c)
	}
	return categorias, nil
}

func (r *categoriaRepository) GetByID(id int) (*domain.Categoria, error) {
	query := `SELECT id_categoria, nombre, fecha_creacion, fecha_actualizacion FROM categorias WHERE id_categoria = $1`
	var c domain.Categoria
	err := r.db.Pool.QueryRow(context.Background(), query, id).Scan(&c.ID, &c.Nombre, &c.FechaCreacion, &c.FechaActualizacion)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.ErrNotFound{Entity: "categoria", ID: id}
		}
		return nil, err
	}
	return &c, nil
}

func (r *categoriaRepository) Create(categoria *domain.Categoria) error {
	query := `INSERT INTO categorias (nombre) VALUES ($1) RETURNING id_categoria, fecha_creacion, fecha_actualizacion`
	return r.db.Pool.QueryRow(context.Background(), query, categoria.Nombre).Scan(&categoria.ID, &categoria.FechaCreacion, &categoria.FechaActualizacion)
}

func (r *categoriaRepository) Update(categoria *domain.Categoria) error {
	query := `UPDATE categorias SET nombre = $2 WHERE id_categoria = $1 RETURNING fecha_actualizacion`
	err := r.db.Pool.QueryRow(context.Background(), query, categoria.ID, categoria.Nombre).Scan(&categoria.FechaActualizacion)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &domain.ErrNotFound{Entity: "categoria", ID: categoria.ID}
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
		return &domain.ErrNotFound{Entity: "categoria", ID: id}
	}
	return nil
}
