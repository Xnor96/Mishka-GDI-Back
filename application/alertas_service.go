package application

import (
	"github.com/Mishka-GDI-Back/domain"
)

type AlertasService interface {
	GetAlertasActivas(limiteStock int) ([]domain.AlertaStockBajo, error)
	GetStockBajo(limite int) ([]domain.AlertaStockBajo, error)
}

type alertasService struct {
	repo domain.AlertasRepository
}

func NewAlertasService(repo domain.AlertasRepository) AlertasService {
	return &alertasService{repo: repo}
}

func (s *alertasService) GetStockBajo(limite int) ([]domain.AlertaStockBajo, error) {
	if limite <= 0 {
		limite = 3
	}
	return s.repo.GetStockBajo(limite)
}

func (s *alertasService) GetAlertasActivas(limiteStock int) ([]domain.AlertaStockBajo, error) {
	return s.GetStockBajo(limiteStock)
}
