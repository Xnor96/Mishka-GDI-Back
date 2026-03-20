package handler

import (
	"net/http"
	"time"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type ControlDiarioHandler struct {
	service application.ControlDiarioService
}

func NewControlDiarioHandler(service application.ControlDiarioService) *ControlDiarioHandler {
	return &ControlDiarioHandler{service: service}
}

func (h *ControlDiarioHandler) GetAll(c *gin.Context) {
	controles, err := h.service.GetAll()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	var totalEntrada, totalSalida float64
	for _, ctrl := range controles {
		totalEntrada += ctrl.MontoEntrada
		totalSalida += ctrl.MontoSalida
	}
	c.JSON(http.StatusOK, dto.ControlDiariosResponse{
		Success:      true,
		Message:      "Controles diarios obtenidos exitosamente",
		Data:         dto.ControlDiariosToResponse(controles),
		TotalCount:   len(controles),
		TotalEntrada: totalEntrada,
		TotalSalida:  totalSalida,
		Balance:      totalEntrada - totalSalida,
	})
}

func (h *ControlDiarioHandler) GetByFecha(c *gin.Context) {
	fecha := c.Param("fecha")
	controles, err := h.service.GetByFecha(fecha)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	var totalEntrada, totalSalida float64
	for _, ctrl := range controles {
		totalEntrada += ctrl.MontoEntrada
		totalSalida += ctrl.MontoSalida
	}
	c.JSON(http.StatusOK, dto.ControlDiariosResponse{
		Success:      true,
		Message:      "Controles del día obtenidos",
		Data:         dto.ControlDiariosToResponse(controles),
		TotalCount:   len(controles),
		TotalEntrada: totalEntrada,
		TotalSalida:  totalSalida,
		Balance:      totalEntrada - totalSalida,
	})
}

func (h *ControlDiarioHandler) GetHoy(c *gin.Context) {
	controles, err := h.service.GetHoy()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	var totalEntrada, totalSalida float64
	for _, ctrl := range controles {
		totalEntrada += ctrl.MontoEntrada
		totalSalida += ctrl.MontoSalida
	}
	c.JSON(http.StatusOK, dto.ControlDiariosResponse{
		Success:      true,
		Message:      "Control del día de hoy",
		Data:         dto.ControlDiariosToResponse(controles),
		TotalCount:   len(controles),
		TotalEntrada: totalEntrada,
		TotalSalida:  totalSalida,
		Balance:      totalEntrada - totalSalida,
	})
}

func (h *ControlDiarioHandler) GetVerbena(c *gin.Context) {
	controles, err := h.service.GetVerbena()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	var totalEntrada, totalSalida float64
	for _, ctrl := range controles {
		totalEntrada += ctrl.MontoEntrada
		totalSalida += ctrl.MontoSalida
	}
	c.JSON(http.StatusOK, dto.ControlDiariosResponse{
		Success:      true,
		Message:      "Controles de verbena obtenidos",
		Data:         dto.ControlDiariosToResponse(controles),
		TotalCount:   len(controles),
		TotalEntrada: totalEntrada,
		TotalSalida:  totalSalida,
		Balance:      totalEntrada - totalSalida,
	})
}

func (h *ControlDiarioHandler) Create(c *gin.Context) {
	var req dto.CreateControlDiarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	fecha, err := time.Parse("2006-01-02", req.Fecha)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Fecha inválida", Error: "Use el formato YYYY-MM-DD"})
		return
	}
	control := &domain.ControlDiario{
		Fecha:           fecha,
		Descripcion:     req.Descripcion,
		MontoEntrada:    req.MontoEntrada,
		MontoSalida:     req.MontoSalida,
		Observaciones:   req.Observaciones,
		EsVerbena:       req.EsVerbena,
		UsuarioRegistro: req.UsuarioRegistro,
	}
	result, err := h.service.Create(control)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Control diario registrado exitosamente",
		Data:    dto.ControlDiarioToResponse(result),
	})
}

func (h *ControlDiarioHandler) GenerarDesdeVentas(c *gin.Context) {
	fecha := c.Param("fecha")
	control, err := h.service.GenerarDesdeVentas(fecha)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Control diario generado desde ventas",
		Data:    dto.ControlDiarioToResponse(control),
	})
}
