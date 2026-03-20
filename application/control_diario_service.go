package application

import (
	"strings"
	"time"

	"github.com/Mishka-GDI-Back/domain"
)

type ControlDiarioService interface {
	GetAll() ([]domain.ControlDiario, error)
	GetByFecha(fecha string) ([]domain.ControlDiario, error)
	GetHoy() ([]domain.ControlDiario, error)
	GetVerbena() ([]domain.ControlDiario, error)
	Create(control *domain.ControlDiario) (*domain.ControlDiario, error)
	GenerarDesdeVentas(fecha string) (*domain.ControlDiario, error)
}

type controlDiarioService struct {
	controlRepo domain.ControlDiarioRepository
}

func NewControlDiarioService(controlRepo domain.ControlDiarioRepository) ControlDiarioService {
	return &controlDiarioService{controlRepo: controlRepo}
}

func (s *controlDiarioService) GetAll() ([]domain.ControlDiario, error) {
	return s.controlRepo.GetAll()
}

func (s *controlDiarioService) GetByFecha(fecha string) ([]domain.ControlDiario, error) {
	if _, err := time.Parse("2006-01-02", fecha); err != nil {
		return nil, &domain.ErrValidation{Field: "fecha", Message: "formato inválido, use YYYY-MM-DD"}
	}
	return s.controlRepo.GetByFecha(fecha)
}

func (s *controlDiarioService) GetHoy() ([]domain.ControlDiario, error) {
	return s.controlRepo.GetByFechaHoy()
}

func (s *controlDiarioService) GetVerbena() ([]domain.ControlDiario, error) {
	return s.controlRepo.GetVerbena()
}

func (s *controlDiarioService) Create(control *domain.ControlDiario) (*domain.ControlDiario, error) {
	if strings.TrimSpace(control.Descripcion) == "" {
		return nil, &domain.ErrValidation{Field: "descripcion", Message: "es requerida"}
	}
	control.Descripcion = strings.TrimSpace(control.Descripcion)
	control.Observaciones = strings.TrimSpace(control.Observaciones)
	control.UsuarioRegistro = strings.TrimSpace(control.UsuarioRegistro)
	if err := s.controlRepo.Create(control); err != nil {
		return nil, err
	}
	return control, nil
}

func (s *controlDiarioService) GenerarDesdeVentas(fecha string) (*domain.ControlDiario, error) {
	if _, err := time.Parse("2006-01-02", fecha); err != nil {
		return nil, &domain.ErrValidation{Field: "fecha", Message: "formato inválido, use YYYY-MM-DD"}
	}
	return s.controlRepo.GenerarDesdeVentas(fecha)
}
