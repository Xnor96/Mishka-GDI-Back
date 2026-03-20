package application

import (
	"github.com/Mishka-GDI-Back/domain"
)

type ResumenMensualService interface {
	GetByMesAnio(mes, anio int) (*domain.ResumenMensual, error)
	GetActual() (*domain.ResumenMensual, error)
	GetByProductoID(productoID, mes, anio int) (*domain.ResumenProducto, error)
	Generar(mes, anio int) (*domain.ResumenMensual, error)
	GuardarManual(resumen *domain.ResumenMensual) (*domain.ResumenMensual, error)
}

type resumenMensualService struct {
	resumenRepo domain.ResumenMensualRepository
}

func NewResumenMensualService(resumenRepo domain.ResumenMensualRepository) ResumenMensualService {
	return &resumenMensualService{resumenRepo: resumenRepo}
}

func (s *resumenMensualService) GetByMesAnio(mes, anio int) (*domain.ResumenMensual, error) {
	if mes < 1 || mes > 12 {
		return nil, &domain.ErrValidation{Field: "mes", Message: "debe estar entre 1 y 12"}
	}
	return s.resumenRepo.GetByMesAnio(mes, anio)
}

func (s *resumenMensualService) GetActual() (*domain.ResumenMensual, error) {
	return s.resumenRepo.GetActual()
}

func (s *resumenMensualService) GetByProductoID(productoID, mes, anio int) (*domain.ResumenProducto, error) {
	if productoID <= 0 {
		return nil, &domain.ErrValidation{Field: "id_producto", Message: "debe ser mayor a 0"}
	}
	return s.resumenRepo.GetByProductoID(productoID, mes, anio)
}

func (s *resumenMensualService) Generar(mes, anio int) (*domain.ResumenMensual, error) {
	if mes < 1 || mes > 12 {
		return nil, &domain.ErrValidation{Field: "mes", Message: "debe estar entre 1 y 12"}
	}
	return s.resumenRepo.Generar(mes, anio)
}

func (s *resumenMensualService) GuardarManual(resumen *domain.ResumenMensual) (*domain.ResumenMensual, error) {
	resumen.Balance = resumen.TotalIngresos - resumen.TotalGastosFijos - resumen.TotalGastosVariables
	if err := s.resumenRepo.Upsert(resumen); err != nil {
		return nil, err
	}
	return resumen, nil
}
