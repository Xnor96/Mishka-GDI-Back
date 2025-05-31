package service

import (
	"errors"
	"strings"

	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/dto"
)

type ProductoService interface {
	GetAll() (*dto.ProductosResponse, error)
	GetByID(id int) (*dto.Response, error)
	Create(request *dto.CreateProductoRequest) (*dto.Response, error)
	Update(id int, request *dto.UpdateProductoRequest) (*dto.Response, error)
	Delete(id int) (*dto.Response, error)
	GetStockBajo(limite int) (*dto.ProductosResponse, error)
	Search(termino string) (*dto.ProductosResponse, error)
}

type productoService struct {
	repo          domain.ProductoRepository
	categoriaRepo domain.CategoriaRepository
}

func NewProductoService(repo domain.ProductoRepository, categoriaRepo domain.CategoriaRepository) ProductoService {
	return &productoService{
		repo:          repo,
		categoriaRepo: categoriaRepo,
	}
}

func (s *productoService) GetAll() (*dto.ProductosResponse, error) {
	productos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return &dto.ProductosResponse{
		Success:    true,
		Message:    "Productos obtenidos exitosamente",
		Data:       dto.ProductosToResponse(productos),
		TotalCount: len(productos),
	}, nil
}

func (s *productoService) GetByID(id int) (*dto.Response, error) {
	if id <= 0 {
		return &dto.Response{
			Success: false,
			Error:   "ID de producto inválido",
		}, errors.New("id inválido")
	}

	producto, err := s.repo.GetByID(id)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &dto.Response{
		Success: true,
		Message: "Producto obtenido exitosamente",
		Data:    dto.ProductoToResponse(producto),
	}, nil
}

func (s *productoService) Create(request *dto.CreateProductoRequest) (*dto.Response, error) {
	// Validaciones adicionales
	if strings.TrimSpace(request.Codigo) == "" {
		return &dto.Response{
			Success: false,
			Error:   "El código del producto es requerido",
		}, errors.New("código requerido")
	}

	if strings.TrimSpace(request.Nombre) == "" {
		return &dto.Response{
			Success: false,
			Error:   "El nombre del producto es requerido",
		}, errors.New("nombre requerido")
	}

	// Validar que la categoría existe si se proporciona
	if request.IDCategoria != nil && *request.IDCategoria > 0 {
		_, err := s.categoriaRepo.GetByID(*request.IDCategoria)
		if err != nil {
			return &dto.Response{
				Success: false,
				Error:   "La categoría especificada no existe",
			}, err
		}
	}

	// Verificar que el código no existe
	existingProducto, _ := s.repo.GetByCodigo(strings.TrimSpace(request.Codigo))
	if existingProducto != nil {
		return &dto.Response{
			Success: false,
			Error:   "Ya existe un producto con ese código",
		}, errors.New("código duplicado")
	}

	unidadMedida := request.UnidadMedida
	if unidadMedida == "" {
		unidadMedida = "UNIDAD"
	}

	producto := &domain.Producto{
		Codigo:         strings.TrimSpace(request.Codigo),
		Nombre:         strings.TrimSpace(request.Nombre),
		IDCategoria:    request.IDCategoria,
		UnidadMedida:   unidadMedida,
		PrecioUnitario: request.PrecioUnitario,
		StockActual:    request.StockActual,
		StockInicial:   request.StockInicial,
	}

	err := s.repo.Create(producto)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Error al crear el producto",
		}, err
	}

	return &dto.Response{
		Success: true,
		Message: "Producto creado exitosamente",
		Data:    dto.ProductoToResponse(producto),
	}, nil
}

func (s *productoService) Update(id int, request *dto.UpdateProductoRequest) (*dto.Response, error) {
	if id <= 0 {
		return &dto.Response{
			Success: false,
			Error:   "ID de producto inválido",
		}, errors.New("id inválido")
	}

	if strings.TrimSpace(request.Codigo) == "" {
		return &dto.Response{
			Success: false,
			Error:   "El código del producto es requerido",
		}, errors.New("código requerido")
	}

	if strings.TrimSpace(request.Nombre) == "" {
		return &dto.Response{
			Success: false,
			Error:   "El nombre del producto es requerido",
		}, errors.New("nombre requerido")
	}

	// Verificar que el producto existe
	existingProducto, err := s.repo.GetByID(id)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Producto no encontrado",
		}, err
	}

	// Validar que la categoría existe si se proporciona
	if request.IDCategoria != nil && *request.IDCategoria > 0 {
		_, err := s.categoriaRepo.GetByID(*request.IDCategoria)
		if err != nil {
			return &dto.Response{
				Success: false,
				Error:   "La categoría especificada no existe",
			}, err
		}
	}

	// Verificar que el código no existe en otro producto
	productoByCodigo, _ := s.repo.GetByCodigo(strings.TrimSpace(request.Codigo))
	if productoByCodigo != nil && productoByCodigo.ID != id {
		return &dto.Response{
			Success: false,
			Error:   "Ya existe otro producto con ese código",
		}, errors.New("código duplicado")
	}

	unidadMedida := request.UnidadMedida
	if unidadMedida == "" {
		unidadMedida = "UNIDAD"
	}

	// Actualizar los campos
	existingProducto.Codigo = strings.TrimSpace(request.Codigo)
	existingProducto.Nombre = strings.TrimSpace(request.Nombre)
	existingProducto.IDCategoria = request.IDCategoria
	existingProducto.UnidadMedida = unidadMedida
	existingProducto.PrecioUnitario = request.PrecioUnitario
	existingProducto.StockActual = request.StockActual
	existingProducto.StockInicial = request.StockInicial

	err = s.repo.Update(existingProducto)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Error al actualizar el producto",
		}, err
	}

	return &dto.Response{
		Success: true,
		Message: "Producto actualizado exitosamente",
		Data:    dto.ProductoToResponse(existingProducto),
	}, nil
}

func (s *productoService) Delete(id int) (*dto.Response, error) {
	if id <= 0 {
		return &dto.Response{
			Success: false,
			Error:   "ID de producto inválido",
		}, errors.New("id inválido")
	}

	// Verificar que el producto existe antes de eliminar
	_, err := s.repo.GetByID(id)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Producto no encontrado",
		}, err
	}

	err = s.repo.Delete(id)
	if err != nil {
		return &dto.Response{
			Success: false,
			Error:   "Error al eliminar el producto",
		}, err
	}

	return &dto.Response{
		Success: true,
		Message: "Producto eliminado exitosamente",
	}, nil
}

func (s *productoService) GetStockBajo(limite int) (*dto.ProductosResponse, error) {
	if limite < 0 {
		limite = 5 // Límite por defecto
	}

	productos, err := s.repo.GetStockBajo(limite)
	if err != nil {
		return nil, err
	}

	return &dto.ProductosResponse{
		Success:    true,
		Message:    "Productos con stock bajo obtenidos exitosamente",
		Data:       dto.ProductosToResponse(productos),
		TotalCount: len(productos),
	}, nil
}

func (s *productoService) Search(termino string) (*dto.ProductosResponse, error) {
	productos, err := s.repo.Search(termino)
	if err != nil {
		return nil, err
	}

	return &dto.ProductosResponse{
		Success:    true,
		Message:    "Búsqueda de productos realizada exitosamente",
		Data:       dto.ProductosToResponse(productos),
		TotalCount: len(productos),
	}, nil
}
