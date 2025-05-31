package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/dto"
)

type SalidaProductoService interface {
	GetAll() (*dto.SalidasResponse, error)
	GetByID(id int) (*dto.Response, error)
	GetByProductoID(productoID int) (*dto.SalidasResponse, error)
	GetByFecha(fecha string) (*dto.SalidasResponse, error)
	Create(request *dto.CreateSalidaProductoRequest) (*dto.Response, error)
}

type salidaProductoService struct {
	salidaRepo   domain.SalidaProductoRepository
	productoRepo domain.ProductoRepository
}

func NewSalidaProductoService(salidaRepo domain.SalidaProductoRepository, productoRepo domain.ProductoRepository) SalidaProductoService {
	return &salidaProductoService{
		salidaRepo:   salidaRepo,
		productoRepo: productoRepo,
	}
}

func (s *salidaProductoService) GetAll() (*dto.SalidasResponse, error) {
	salidas, err := s.salidaRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &dto.SalidasResponse{
		Success:    true,
		Message:    "Salidas obtenidas exitosamente",
		Data:       dto.SalidasToResponse(salidas),
		TotalCount: len(salidas),
	}, nil
}

func (s *salidaProductoService) GetByID(id int) (*dto.Response, error) {
	if id <= 0 {
		return &dto.Response{
			Success: false,
			Error:   "ID de salida inválido",
		}, errors.New("id inválido")
	}

	salida, err := s.salidaRepo.GetByID(id)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &dto.Response{
		Success: true,
		Message: "Salida obtenida exitosamente",
		Data:    dto.SalidaProductoToResponse(salida),
	}, nil
}

func (s *salidaProductoService) GetByProductoID(productoID int) (*dto.SalidasResponse, error) {
	if productoID <= 0 {
		return &dto.SalidasResponse{
			Success: false,
			Message: "ID de producto inválido",
			Data:    []dto.SalidaProductoResponse{},
		}, errors.New("id de producto inválido")
	}

	// Verificar que el producto existe
	_, err := s.productoRepo.GetByID(productoID)
	if err != nil {
		return &dto.SalidasResponse{
			Success: false,
			Message: "Producto no encontrado",
			Data:    []dto.SalidaProductoResponse{},
		}, err
	}

	salidas, err := s.salidaRepo.GetByProductoID(productoID)
	if err != nil {
		return nil, err
	}

	return &dto.SalidasResponse{
		Success:    true,
		Message:    "Salidas del producto obtenidas exitosamente",
		Data:       dto.SalidasToResponse(salidas),
		TotalCount: len(salidas),
	}, nil
}

func (s *salidaProductoService) GetByFecha(fecha string) (*dto.SalidasResponse, error) {
	// Validar formato de fecha
	_, err := time.Parse("2006-01-02", fecha)
	if err != nil {
		return &dto.SalidasResponse{
			Success: false,
			Message: "Formato de fecha inválido. Use YYYY-MM-DD",
			Data:    []dto.SalidaProductoResponse{},
		}, errors.New("formato de fecha inválido")
	}

	salidas, err := s.salidaRepo.GetByFecha(fecha)
	if err != nil {
		return nil, err
	}

	return &dto.SalidasResponse{
		Success:    true,
		Message:    "Salidas de la fecha obtenidas exitosamente",
		Data:       dto.SalidasToResponse(salidas),
		TotalCount: len(salidas),
	}, nil
}

func (s *salidaProductoService) Create(request *dto.CreateSalidaProductoRequest) (*dto.Response, error) {
	// Validaciones
	if request.IDProducto <= 0 {
		return &dto.Response{
			Success: false,
			Error:   "ID de producto inválido",
		}, errors.New("id de producto inválido")
	}

	if request.Cantidad <= 0 {
		return &dto.Response{
			Success: false,
			Error:   "La cantidad debe ser mayor a 0",
		}, errors.New("cantidad inválida")
	}

	if strings.TrimSpace(request.UsuarioRegistro) == "" {
		return &dto.Response{
			Success: false,
			Error:   "Usuario de registro es requerido",
		}, errors.New("usuario requerido")
	}

	// Validar que el producto existe y tiene suficiente stock
	producto, err := s.productoRepo.GetByID(request.IDProducto)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Producto no encontrado",
		}, err
	}

	if producto.StockActual < request.Cantidad {
		return &dto.Response{
			Success: false,
			Error:   "Stock insuficiente para realizar la salida",
		}, errors.New("stock insuficiente")
	}

	// Validar formato de fecha
	fechaSalida, err := time.Parse("2006-01-02", request.FechaSalida)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Formato de fecha inválido. Use YYYY-MM-DD",
		}, err
	}

	salida := &domain.SalidaProducto{
		IDProducto:      request.IDProducto,
		FechaSalida:     fechaSalida,
		Cantidad:        request.Cantidad,
		Observaciones:   strings.TrimSpace(request.Observaciones),
		UsuarioRegistro: strings.TrimSpace(request.UsuarioRegistro),
	}

	err = s.salidaRepo.Create(salida)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Error al crear la salida",
		}, err
	}

	// Actualizar el stock del producto
	producto.StockActual -= request.Cantidad
	err = s.productoRepo.Update(producto)
	if err != nil {
		// TODO: Implementar rollback o transacciones
		return &dto.Response{
			Success: false,
			Error:   "Error al actualizar el stock del producto",
		}, err
	}

	return &dto.Response{
		Success: true,
		Message: "Salida creada exitosamente",
		Data:    dto.SalidaProductoToResponse(salida),
	}, nil
}
