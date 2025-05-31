package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/dto"
)

type EntradaProductoService interface {
	GetAll() (*dto.EntradasResponse, error)
	GetByID(id int) (*dto.Response, error)
	GetByProductoID(productoID int) (*dto.EntradasResponse, error)
	GetByFecha(fecha string) (*dto.EntradasResponse, error)
	Create(request *dto.CreateEntradaProductoRequest) (*dto.Response, error)
}

type entradaProductoService struct {
	entradaRepo  domain.EntradaProductoRepository
	productoRepo domain.ProductoRepository
}

func NewEntradaProductoService(entradaRepo domain.EntradaProductoRepository, productoRepo domain.ProductoRepository) EntradaProductoService {
	return &entradaProductoService{
		entradaRepo:  entradaRepo,
		productoRepo: productoRepo,
	}
}

func (s *entradaProductoService) GetAll() (*dto.EntradasResponse, error) {
	entradas, err := s.entradaRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &dto.EntradasResponse{
		Success:    true,
		Message:    "Entradas obtenidas exitosamente",
		Data:       dto.EntradasToResponse(entradas),
		TotalCount: len(entradas),
	}, nil
}

func (s *entradaProductoService) GetByID(id int) (*dto.Response, error) {
	if id <= 0 {
		return &dto.Response{
			Success: false,
			Error:   "ID de entrada inválido",
		}, errors.New("id inválido")
	}

	entrada, err := s.entradaRepo.GetByID(id)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &dto.Response{
		Success: true,
		Message: "Entrada obtenida exitosamente",
		Data:    dto.EntradaProductoToResponse(entrada),
	}, nil
}

func (s *entradaProductoService) GetByProductoID(productoID int) (*dto.EntradasResponse, error) {
	if productoID <= 0 {
		return &dto.EntradasResponse{
			Success: false,
			Message: "ID de producto inválido",
			Data:    []dto.EntradaProductoResponse{},
		}, errors.New("id de producto inválido")
	}

	// Verificar que el producto existe
	_, err := s.productoRepo.GetByID(productoID)
	if err != nil {
		return &dto.EntradasResponse{
			Success: false,
			Message: "Producto no encontrado",
			Data:    []dto.EntradaProductoResponse{},
		}, err
	}

	entradas, err := s.entradaRepo.GetByProductoID(productoID)
	if err != nil {
		return nil, err
	}

	return &dto.EntradasResponse{
		Success:    true,
		Message:    "Entradas del producto obtenidas exitosamente",
		Data:       dto.EntradasToResponse(entradas),
		TotalCount: len(entradas),
	}, nil
}

func (s *entradaProductoService) GetByFecha(fecha string) (*dto.EntradasResponse, error) {
	// Validar formato de fecha
	_, err := time.Parse("2006-01-02", fecha)
	if err != nil {
		return &dto.EntradasResponse{
			Success: false,
			Message: "Formato de fecha inválido. Use YYYY-MM-DD",
			Data:    []dto.EntradaProductoResponse{},
		}, errors.New("formato de fecha inválido")
	}

	entradas, err := s.entradaRepo.GetByFecha(fecha)
	if err != nil {
		return nil, err
	}

	return &dto.EntradasResponse{
		Success:    true,
		Message:    "Entradas de la fecha obtenidas exitosamente",
		Data:       dto.EntradasToResponse(entradas),
		TotalCount: len(entradas),
	}, nil
}

func (s *entradaProductoService) Create(request *dto.CreateEntradaProductoRequest) (*dto.Response, error) {
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

	// Validar que el producto existe
	producto, err := s.productoRepo.GetByID(request.IDProducto)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Producto no encontrado",
		}, err
	}

	// Validar formato de fecha
	fechaEntrada, err := time.Parse("2006-01-02", request.FechaEntrada)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Formato de fecha inválido. Use YYYY-MM-DD",
		}, err
	}

	// Validar precio unitario si se proporciona
	if request.PrecioUnitario != nil && *request.PrecioUnitario < 0 {
		return &dto.Response{
			Success: false,
			Error:   "El precio unitario no puede ser negativo",
		}, errors.New("precio inválido")
	}

	entrada := &domain.EntradaProducto{
		IDProducto:      request.IDProducto,
		FechaEntrada:    fechaEntrada,
		Cantidad:        request.Cantidad,
		PrecioUnitario:  request.PrecioUnitario,
		Observaciones:   strings.TrimSpace(request.Observaciones),
		UsuarioRegistro: strings.TrimSpace(request.UsuarioRegistro),
	}

	err = s.entradaRepo.Create(entrada)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Error al crear la entrada",
		}, err
	}

	// Actualizar el stock del producto
	producto.StockActual += request.Cantidad
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
		Message: "Entrada creada exitosamente",
		Data:    dto.EntradaProductoToResponse(entrada),
	}, nil
}
