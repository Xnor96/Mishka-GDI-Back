package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/domain"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type SalidaHandler struct {
	service application.SalidaProductoService
}

func NewSalidaHandler(service application.SalidaProductoService) *SalidaHandler {
	return &SalidaHandler{service: service}
}

func (h *SalidaHandler) GetAll(c *gin.Context) {
	salidas, err := h.service.GetAll()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.SalidasResponse{
		Success:    true,
		Message:    "Salidas obtenidas exitosamente",
		Data:       dto.SalidasConProductoToResponse(salidas),
		TotalCount: len(salidas),
	})
}

func (h *SalidaHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	salida, err := h.service.GetByID(id)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Salida encontrada",
		Data:    dto.SalidaConProductoToResponse(salida),
	})
}

func (h *SalidaHandler) GetByProductoID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	salidas, err := h.service.GetByProductoID(id)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.SalidasResponse{
		Success:    true,
		Message:    "Salidas del producto obtenidas",
		Data:       dto.SalidasConProductoToResponse(salidas),
		TotalCount: len(salidas),
	})
}

func (h *SalidaHandler) GetByFecha(c *gin.Context) {
	fecha := c.Param("fecha")
	salidas, err := h.service.GetByFecha(fecha)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.SalidasResponse{
		Success:    true,
		Message:    "Salidas de la fecha obtenidas",
		Data:       dto.SalidasConProductoToResponse(salidas),
		TotalCount: len(salidas),
	})
}

func (h *SalidaHandler) GetByLugar(c *gin.Context) {
	lugar := c.Param("lugar")
	salidas, err := h.service.GetByLugar(lugar)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.SalidasResponse{
		Success:    true,
		Message:    "Salidas del lugar obtenidas",
		Data:       dto.SalidasConProductoToResponse(salidas),
		TotalCount: len(salidas),
	})
}

func (h *SalidaHandler) Create(c *gin.Context) {
	var req dto.CreateSalidaProductoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	fechaSalida, err := time.Parse("2006-01-02", req.FechaSalida)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Fecha inválida", Error: "Use el formato YYYY-MM-DD"})
		return
	}
	salida := &domain.SalidaProducto{
		IDProducto:      req.IDProducto,
		FechaSalida:     fechaSalida,
		Cantidad:        req.Cantidad,
		PrecioVenta:     req.PrecioVenta,
		Descuento:       req.Descuento,
		LugarVenta:      req.LugarVenta,
		TipoPago:        req.TipoPago,
		Observaciones:   req.Observaciones,
		UsuarioRegistro: req.UsuarioRegistro,
	}
	result, err := h.service.Create(salida)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Salida registrada exitosamente",
		Data:    result,
	})
}
