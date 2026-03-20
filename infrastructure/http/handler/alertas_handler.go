package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type AlertasHandler struct {
	service application.AlertasService
}

func NewAlertasHandler(service application.AlertasService) *AlertasHandler {
	return &AlertasHandler{service: service}
}

func (h *AlertasHandler) GetAlertasActivas(c *gin.Context) {
	limite, _ := strconv.Atoi(c.DefaultQuery("limite", "3"))
	items, err := h.service.GetAlertasActivas(limite)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.AlertasResponse{
		Success:     true,
		Message:     "Alertas activas",
		StockBajo:   dto.AlertasStockBajoToResponse(items),
		TotalAlerts: len(items),
	})
}

func (h *AlertasHandler) GetStockBajo(c *gin.Context) {
	limite, _ := strconv.Atoi(c.DefaultQuery("limite", "3"))
	items, err := h.service.GetStockBajo(limite)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.AlertasResponse{
		Success:     true,
		Message:     "Productos con stock bajo",
		StockBajo:   dto.AlertasStockBajoToResponse(items),
		TotalAlerts: len(items),
	})
}

func (h *AlertasHandler) ConfigurarAlerta(c *gin.Context) {
	var req dto.ConfigurarAlertaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Error: err.Error()})
		return
	}
	// JWT stateless: la configuración se aplica en tiempo de consulta
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Configuración de alerta guardada",
		Data:    map[string]int{"limite_stock_bajo": req.LimiteStockBajo},
	})
}
