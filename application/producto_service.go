package application

import (
	"strings"

	"github.com/Mishka-GDI-Back/domain"
)

type ProductoService interface {
	GetAll() ([]domain.Producto, error)
	GetByID(id int) (*domain.Producto, error)
	Create(producto *domain.Producto) (*domain.Producto, error)
	Update(id int, producto *domain.Producto) (*domain.Producto, error)
	Delete(id int) error
	GetStockBajo(limite int) ([]domain.Producto, error)
	Search(termino string) ([]domain.Producto, error)
}

type productoService struct {
	repo          domain.ProductoRepository
	categoriaRepo domain.CategoriaRepository
}

func NewProductoService(repo domain.ProductoRepository, categoriaRepo domain.CategoriaRepository) ProductoService {
	return &productoService{repo: repo, categoriaRepo: categoriaRepo}
}

func (s *productoService) GetAll() ([]domain.Producto, error) {
	return s.repo.GetAll()
}

func (s *productoService) GetByID(id int) (*domain.Producto, error) {
	if id <= 0 {
		return nil, &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	return s.repo.GetByID(id)
}

func (s *productoService) Create(producto *domain.Producto) (*domain.Producto, error) {
	if strings.TrimSpace(producto.Codigo) == "" {
		return nil, &domain.ErrValidation{Field: "codigo", Message: "es requerido"}
	}
	if strings.TrimSpace(producto.Nombre) == "" {
		return nil, &domain.ErrValidation{Field: "nombre", Message: "es requerido"}
	}
	if producto.IDCategoria != nil && *producto.IDCategoria > 0 {
		if _, err := s.categoriaRepo.GetByID(*producto.IDCategoria); err != nil {
			return nil, &domain.ErrValidation{Field: "id_categoria", Message: "la categoría especificada no existe"}
		}
	}
	existing, _ := s.repo.GetByCodigo(strings.TrimSpace(producto.Codigo))
	if existing != nil {
		return nil, &domain.ErrDuplicate{Entity: "producto", Field: "codigo", Value: producto.Codigo}
	}
	producto.Codigo = strings.TrimSpace(producto.Codigo)
	producto.Nombre = strings.TrimSpace(producto.Nombre)
	if producto.UnidadMedida == "" {
		producto.UnidadMedida = "UNIDAD"
	}
	if err := s.repo.Create(producto); err != nil {
		return nil, err
	}
	return producto, nil
}

func (s *productoService) Update(id int, producto *domain.Producto) (*domain.Producto, error) {
	if id <= 0 {
		return nil, &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	if strings.TrimSpace(producto.Codigo) == "" {
		return nil, &domain.ErrValidation{Field: "codigo", Message: "es requerido"}
	}
	if strings.TrimSpace(producto.Nombre) == "" {
		return nil, &domain.ErrValidation{Field: "nombre", Message: "es requerido"}
	}
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if producto.IDCategoria != nil && *producto.IDCategoria > 0 {
		if _, err := s.categoriaRepo.GetByID(*producto.IDCategoria); err != nil {
			return nil, &domain.ErrValidation{Field: "id_categoria", Message: "la categoría especificada no existe"}
		}
	}
	byCode, _ := s.repo.GetByCodigo(strings.TrimSpace(producto.Codigo))
	if byCode != nil && byCode.ID != id {
		return nil, &domain.ErrDuplicate{Entity: "producto", Field: "codigo", Value: producto.Codigo}
	}
	existing.Codigo = strings.TrimSpace(producto.Codigo)
	existing.Nombre = strings.TrimSpace(producto.Nombre)
	existing.IDCategoria = producto.IDCategoria
	existing.UnidadMedida = producto.UnidadMedida
	if existing.UnidadMedida == "" {
		existing.UnidadMedida = "UNIDAD"
	}
	existing.PrecioUnitario = producto.PrecioUnitario
	existing.StockActual = producto.StockActual
	existing.StockInicial = producto.StockInicial
	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *productoService) Delete(id int) error {
	if id <= 0 {
		return &domain.ErrValidation{Field: "id", Message: "debe ser mayor a 0"}
	}
	if _, err := s.repo.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *productoService) GetStockBajo(limite int) ([]domain.Producto, error) {
	if limite < 0 {
		limite = 5
	}
	return s.repo.GetStockBajo(limite)
}

func (s *productoService) Search(termino string) ([]domain.Producto, error) {
	return s.repo.Search(termino)
}
