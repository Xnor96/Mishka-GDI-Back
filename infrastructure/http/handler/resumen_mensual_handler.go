package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type ResumenMensualHandler struct {
	service application.ResumenMensualService
}

func NewResumenMensualHandler(service application.ResumenMensualService) *ResumenMensualHandler {
	return &ResumenMensualHandler{service: service}
}

func (h *ResumenMensualHandler) GetByMesAnio(c *gin.Context) {
	mes, err := strconv.Atoi(c.Param("mes"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Error: "Mes inválido"})
		return
	}
	anio, err := strconv.Atoi(c.Param("anio"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Error: "Año inválido"})
		return
	}
	resumen, err := h.service.GetByMesAnio(mes, anio)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Resumen mensual obtenido",
		Data:    dto.ResumenMensualToResponse(resumen),
	})
}

func (h *ResumenMensualHandler) GetActual(c *gin.Context) {
	resumen, err := h.service.GetActual()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Resumen mensual actual obtenido",
		Data:    dto.ResumenMensualToResponse(resumen),
	})
}

func (h *ResumenMensualHandler) GetByProductoID(c *gin.Context) {
	productoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Error: "ID inválido"})
		return
	}
	mes, _ := strconv.Atoi(c.DefaultQuery("mes", "0"))
	anio, _ := strconv.Atoi(c.DefaultQuery("anio", "0"))

	resumenProducto, err := h.service.GetByProductoID(productoID, mes, anio)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Resumen del producto obtenido",
		Data:    dto.ResumenProductoToResponse(resumenProducto),
	})
}

func (h *ResumenMensualHandler) Generar(c *gin.Context) {
	var req dto.GenerarResumenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Error: err.Error()})
		return
	}

	var resumen *domain.ResumenMensual
	var err error

	// Si se proveen montos manuales, guardar directamente
	if req.TotalIngresos > 0 || req.TotalGastosFijos > 0 || req.TotalGastosVariables > 0 {
		manual := &domain.ResumenMensual{
			Mes:                  req.Mes,
			Anio:                 req.Anio,
			TotalIngresos:        req.TotalIngresos,
			TotalGastosFijos:     req.TotalGastosFijos,
			TotalGastosVariables: req.TotalGastosVariables,
			Observaciones:        req.Observaciones,
		}
		resumen, err = h.service.GuardarManual(manual)
	} else {
		resumen, err = h.service.Generar(req.Mes, req.Anio)
	}

	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Resumen mensual generado exitosamente",
		Data:    dto.ResumenMensualToResponse(resumen),
	})
}
