package application

import (
	"github.com/Mishka-GDI-Back/domain"
)

type ReportesService interface {
	GetInventarioActual() ([]domain.ReporteInventarioItem, error)
	GetMovimientos(inicio, fin string) ([]domain.ReporteMovimiento, error)
	GetProductosMasVendidos(limite int) ([]domain.ReporteProductoVendido, error)
	GetProductosMasIngresados(limite int) ([]domain.ReporteProductoVendido, error)
	GetValoracionInventario() ([]domain.ReporteValoracion, error)
}

type reportesService struct {
	repo domain.ReportesRepository
}

func NewReportesService(repo domain.ReportesRepository) ReportesService {
	return &reportesService{repo: repo}
}

func (s *reportesService) GetInventarioActual() ([]domain.ReporteInventarioItem, error) {
	return s.repo.GetInventarioActual()
}

func (s *reportesService) GetMovimientos(inicio, fin string) ([]domain.ReporteMovimiento, error) {
	return s.repo.GetMovimientos(inicio, fin)
}

func (s *reportesService) GetProductosMasVendidos(limite int) ([]domain.ReporteProductoVendido, error) {
	if limite <= 0 {
		limite = 10
	}
	return s.repo.GetProductosMasVendidos(limite)
}

func (s *reportesService) GetProductosMasIngresados(limite int) ([]domain.ReporteProductoVendido, error) {
	if limite <= 0 {
		limite = 10
	}
	return s.repo.GetProductosMasIngresados(limite)
}

func (s *reportesService) GetValoracionInventario() ([]domain.ReporteValoracion, error) {
	return s.repo.GetValoracionInventario()
}
