package application

import (
	"strings"
	"time"

	"github.com/Mishka-GDI-Back/domain"
)

type EntradaProductoService interface {
	GetAll() ([]domain.EntradaConProducto, error)
	GetByID(id int) (*domain.EntradaConProducto, error)
	GetByProductoID(productoID int) ([]domain.EntradaConProducto, error)
	GetByFecha(fecha string) ([]domain.EntradaConProducto, error)
	Create(entrada *domain.EntradaProducto) (*domain.EntradaProducto, error)
}

type entradaProductoService struct {
	entradaRepo  domain.EntradaProductoRepository
	productoRepo domain.ProductoRepository
}

func NewEntradaProductoService(entradaRepo domain.EntradaProductoRepository, productoRepo domain.ProductoRepository) EntradaProductoService {
	return &entradaProductoService{entradaRepo: entradaRepo, productoRepo: productoRepo}
}

func (s *entradaProductoService) GetAll() ([]domain.EntradaConProducto, error) {
	return s.entradaRepo.GetAll()
}

func (s *entradaProductoService) GetByID(id int) (*domain.EntradaConProducto, error) {
	if id <= 0 {
		return nil, &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	return s.entradaRepo.GetByID(id)
}

func (s *entradaProductoService) GetByProductoID(productoID int) ([]domain.EntradaConProducto, error) {
	if productoID <= 0 {
		return nil, &domain.ErrValidation{Field: "id_producto", Message: "debe ser mayor a 0"}
	}
	if _, err := s.productoRepo.GetByID(productoID); err != nil {
		return nil, err
	}
	return s.entradaRepo.GetByProductoID(productoID)
}

func (s *entradaProductoService) GetByFecha(fecha string) ([]domain.EntradaConProducto, error) {
	if _, err := time.Parse("2006-01-02", fecha); err != nil {
		return nil, &domain.ErrValidation{Field: "fecha", Message: "formato inválido, use YYYY-MM-DD"}
	}
	return s.entradaRepo.GetByFecha(fecha)
}

func (s *entradaProductoService) Create(entrada *domain.EntradaProducto) (*domain.EntradaProducto, error) {
	if entrada.IDProducto <= 0 {
		return nil, &domain.ErrValidation{Field: "id_producto", Message: "debe ser mayor a 0"}
	}
	if entrada.Cantidad <= 0 {
		return nil, &domain.ErrValidation{Field: "cantidad", Message: "debe ser mayor a 0"}
	}
	if strings.TrimSpace(entrada.UsuarioRegistro) == "" {
		return nil, &domain.ErrValidation{Field: "usuario_registro", Message: "es requerido"}
	}
	if entrada.PrecioUnitario != nil && *entrada.PrecioUnitario < 0 {
		return nil, &domain.ErrValidation{Field: "precio_unitario", Message: "no puede ser negativo"}
	}
	producto, err := s.productoRepo.GetByID(entrada.IDProducto)
	if err != nil {
		return nil, err
	}
	entrada.Observaciones = strings.TrimSpace(entrada.Observaciones)
	entrada.UsuarioRegistro = strings.TrimSpace(entrada.UsuarioRegistro)
	if err := s.entradaRepo.Create(entrada); err != nil {
		return nil, err
	}
	// Actualizar stock
	producto.StockActual += entrada.Cantidad
	if err := s.productoRepo.Update(producto); err != nil {
		return nil, err
	}
	return entrada, nil
}
