package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type ReportesHandler struct {
	service application.ReportesService
}

func NewReportesHandler(service application.ReportesService) *ReportesHandler {
	return &ReportesHandler{service: service}
}

func (h *ReportesHandler) GetInventarioActual(c *gin.Context) {
	items, err := h.service.GetInventarioActual()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Inventario actual",
		Data:    dto.ReportesInventarioToResponse(items),
	})
}

func (h *ReportesHandler) GetMovimientos(c *gin.Context) {
	inicio := c.Param("inicio")
	fin := c.Param("fin")
	items, err := h.service.GetMovimientos(inicio, fin)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Movimientos de " + inicio + " a " + fin,
		Data:    dto.ReportesMovimientoToResponse(items),
	})
}

func (h *ReportesHandler) GetProductosMasVendidos(c *gin.Context) {
	limite, _ := strconv.Atoi(c.DefaultQuery("limite", "10"))
	items, err := h.service.GetProductosMasVendidos(limite)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Productos más vendidos",
		Data:    dto.ReportesProductoVendidoToResponse(items),
	})
}

func (h *ReportesHandler) GetProductosMasIngresados(c *gin.Context) {
	limite, _ := strconv.Atoi(c.DefaultQuery("limite", "10"))
	items, err := h.service.GetProductosMasIngresados(limite)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Productos más ingresados",
		Data:    dto.ReportesProductoVendidoToResponse(items),
	})
}

func (h *ReportesHandler) GetValoracionInventario(c *gin.Context) {
	items, err := h.service.GetValoracionInventario()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Valoración del inventario por categoría",
		Data:    dto.ReportesValoracionToResponse(items),
	})
}
