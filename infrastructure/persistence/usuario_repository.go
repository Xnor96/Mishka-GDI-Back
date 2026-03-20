package persistence

import (
	"context"
	"errors"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

type usuarioRepository struct {
	db *database.Database
}

func NewUsuarioRepository(db *database.Database) domain.UsuarioRepository {
	return &usuarioRepository{db: db}
}

const usuarioSelect = `SELECT id_usuario, username, email, password_hash, nombre, rol, activo, fecha_creacion, fecha_actualizacion FROM usuarios`

func (r *usuarioRepository) GetByUsername(username string) (*domain.Usuario, error) {
	var u domain.Usuario
	err := r.db.Pool.QueryRow(context.Background(), usuarioSelect+" WHERE username = $1 AND activo = TRUE", username).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Nombre, &u.Rol, &u.Activo, &u.FechaCreacion, &u.FechaActualizacion)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.ErrNotFound{Entity: "usuario", ID: username}
		}
		return nil, err
	}
	return &u, nil
}

func (r *usuarioRepository) GetByID(id int) (*domain.Usuario, error) {
	var u domain.Usuario
	err := r.db.Pool.QueryRow(context.Background(), usuarioSelect+" WHERE id_usuario = $1", id).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Nombre, &u.Rol, &u.Activo, &u.FechaCreacion, &u.FechaActualizacion)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &domain.ErrNotFound{Entity: "usuario", ID: id}
		}
		return nil, err
	}
	return &u, nil
}

func (r *usuarioRepository) Create(u *domain.Usuario) error {
	query := `INSERT INTO usuarios (username, email, password_hash, nombre, rol) VALUES ($1, $2, $3, $4, $5) RETURNING id_usuario, fecha_creacion, fecha_actualizacion`
	return r.db.Pool.QueryRow(context.Background(), query, u.Username, u.Email, u.PasswordHash, u.Nombre, u.Rol).Scan(&u.ID, &u.FechaCreacion, &u.FechaActualizacion)
}

func (r *usuarioRepository) Update(u *domain.Usuario) error {
	query := `UPDATE usuarios SET username = $2, email = $3, nombre = $4, rol = $5, activo = $6 WHERE id_usuario = $1 RETURNING fecha_actualizacion`
	return r.db.Pool.QueryRow(context.Background(), query, u.ID, u.Username, u.Email, u.Nombre, u.Rol, u.Activo).Scan(&u.FechaActualizacion)
}
