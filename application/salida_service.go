package application

import (
	"strings"
	"time"

	"github.com/Mishka-GDI-Back/domain"
)

type SalidaProductoService interface {
	GetAll() ([]domain.SalidaConProducto, error)
	GetByID(id int) (*domain.SalidaConProducto, error)
	GetByProductoID(productoID int) ([]domain.SalidaConProducto, error)
	GetByFecha(fecha string) ([]domain.SalidaConProducto, error)
	GetByLugar(lugar string) ([]domain.SalidaConProducto, error)
	Create(salida *domain.SalidaProducto) (*domain.SalidaProducto, error)
}

type salidaProductoService struct {
	salidaRepo   domain.SalidaProductoRepository
	productoRepo domain.ProductoRepository
}

func NewSalidaProductoService(salidaRepo domain.SalidaProductoRepository, productoRepo domain.ProductoRepository) SalidaProductoService {
	return &salidaProductoService{salidaRepo: salidaRepo, productoRepo: productoRepo}
}

func (s *salidaProductoService) GetAll() ([]domain.SalidaConProducto, error) {
	return s.salidaRepo.GetAll()
}

func (s *salidaProductoService) GetByID(id int) (*domain.SalidaConProducto, error) {
	if id <= 0 {
		return nil, &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	return s.salidaRepo.GetByID(id)
}

func (s *salidaProductoService) GetByProductoID(productoID int) ([]domain.SalidaConProducto, error) {
	if productoID <= 0 {
		return nil, &domain.ErrValidation{Field: "id_producto", Message: "debe ser mayor a 0"}
	}
	if _, err := s.productoRepo.GetByID(productoID); err != nil {
		return nil, err
	}
	return s.salidaRepo.GetByProductoID(productoID)
}

func (s *salidaProductoService) GetByFecha(fecha string) ([]domain.SalidaConProducto, error) {
	if _, err := time.Parse("2006-01-02", fecha); err != nil {
		return nil, &domain.ErrValidation{Field: "fecha", Message: "formato inválido, use YYYY-MM-DD"}
	}
	return s.salidaRepo.GetByFecha(fecha)
}

func (s *salidaProductoService) GetByLugar(lugar string) ([]domain.SalidaConProducto, error) {
	if strings.TrimSpace(lugar) == "" {
		return nil, &domain.ErrValidation{Field: "lugar", Message: "es requerido"}
	}
	return s.salidaRepo.GetByLugar(lugar)
}

func (s *salidaProductoService) Create(salida *domain.SalidaProducto) (*domain.SalidaProducto, error) {
	if salida.IDProducto <= 0 {
		return nil, &domain.ErrValidation{Field: "id_producto", Message: "debe ser mayor a 0"}
	}
	if salida.Cantidad <= 0 {
		return nil, &domain.ErrValidation{Field: "cantidad", Message: "debe ser mayor a 0"}
	}
	if strings.TrimSpace(salida.UsuarioRegistro) == "" {
		return nil, &domain.ErrValidation{Field: "usuario_registro", Message: "es requerido"}
	}
	producto, err := s.productoRepo.GetByID(salida.IDProducto)
	if err != nil {
		return nil, err
	}
	if producto.StockActual < salida.Cantidad {
		return nil, &domain.ErrInsufficientStock{
			ProductoID:  salida.IDProducto,
			StockActual: producto.StockActual,
			CantidadReq: salida.Cantidad,
		}
	}
	// Calcular total
	salida.Total = salida.PrecioVenta*float64(salida.Cantidad) - salida.Descuento
	if salida.Total < 0 {
		salida.Total = 0
	}
	salida.LugarVenta = strings.TrimSpace(salida.LugarVenta)
	salida.TipoPago = strings.TrimSpace(salida.TipoPago)
	salida.Observaciones = strings.TrimSpace(salida.Observaciones)
	salida.UsuarioRegistro = strings.TrimSpace(salida.UsuarioRegistro)
	if err := s.salidaRepo.Create(salida); err != nil {
		return nil, err
	}
	// Actualizar stock
	producto.StockActual -= salida.Cantidad
	if err := s.productoRepo.Update(producto); err != nil {
		return nil, err
	}
	return salida, nil
}
