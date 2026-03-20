package application

import (
	"errors"
	"strings"

	"github.com/Mishka-GDI-Back/domain"
)

type CategoriaService interface {
	GetAll() ([]domain.Categoria, error)
	GetByID(id int) (*domain.Categoria, error)
	Create(nombre string) (*domain.Categoria, error)
	Update(id int, nombre string) (*domain.Categoria, error)
	Delete(id int) error
}

type categoriaService struct {
	repo domain.CategoriaRepository
}

func NewCategoriaService(repo domain.CategoriaRepository) CategoriaService {
	return &categoriaService{repo: repo}
}

func (s *categoriaService) GetAll() ([]domain.Categoria, error) {
	return s.repo.GetAll()
}

func (s *categoriaService) GetByID(id int) (*domain.Categoria, error) {
	if id <= 0 {
		return nil, &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	return s.repo.GetByID(id)
}

func (s *categoriaService) Create(nombre string) (*domain.Categoria, error) {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return nil, &domain.ErrValidation{Field: "nombre", Message: "es requerido"}
	}
	categoria := &domain.Categoria{Nombre: nombre}
	if err := s.repo.Create(categoria); err != nil {
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			return nil, &domain.ErrDuplicate{Entity: "categoría", Field: "nombre", Value: nombre}
		}
		return nil, err
	}
	return categoria, nil
}

func (s *categoriaService) Update(id int, nombre string) (*domain.Categoria, error) {
	if id <= 0 {
		return nil, &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return nil, &domain.ErrValidation{Field: "nombre", Message: "es requerido"}
	}
	categoria, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	categoria.Nombre = nombre
	if err := s.repo.Update(categoria); err != nil {
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			return nil, &domain.ErrDuplicate{Entity: "categoría", Field: "nombre", Value: nombre}
		}
		return nil, err
	}
	return categoria, nil
}

func (s *categoriaService) Delete(id int) error {
	if id <= 0 {
		return &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	if _, err := s.repo.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

var _ = errors.New // keep import
