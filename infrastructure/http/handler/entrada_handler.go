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

type EntradaHandler struct {
	service application.EntradaProductoService
}

func NewEntradaHandler(service application.EntradaProductoService) *EntradaHandler {
	return &EntradaHandler{service: service}
}

func (h *EntradaHandler) GetAll(c *gin.Context) {
	entradas, err := h.service.GetAll()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.EntradasResponse{
		Success:    true,
		Message:    "Entradas obtenidas exitosamente",
		Data:       dto.EntradasConProductoToResponse(entradas),
		TotalCount: len(entradas),
	})
}

func (h *EntradaHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	entrada, err := h.service.GetByID(id)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Entrada encontrada",
		Data:    dto.EntradaConProductoToResponse(entrada),
	})
}

func (h *EntradaHandler) GetByProductoID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	entradas, err := h.service.GetByProductoID(id)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.EntradasResponse{
		Success:    true,
		Message:    "Entradas del producto obtenidas",
		Data:       dto.EntradasConProductoToResponse(entradas),
		TotalCount: len(entradas),
	})
}

func (h *EntradaHandler) GetByFecha(c *gin.Context) {
	fecha := c.Param("fecha")
	entradas, err := h.service.GetByFecha(fecha)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.EntradasResponse{
		Success:    true,
		Message:    "Entradas de la fecha obtenidas",
		Data:       dto.EntradasConProductoToResponse(entradas),
		TotalCount: len(entradas),
	})
}

func (h *EntradaHandler) Create(c *gin.Context) {
	var req dto.CreateEntradaProductoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	fechaEntrada, err := time.Parse("2006-01-02", req.FechaEntrada)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Fecha inválida", Error: "Use el formato YYYY-MM-DD"})
		return
	}
	entrada := &domain.EntradaProducto{
		IDProducto:      req.IDProducto,
		FechaEntrada:    fechaEntrada,
		Cantidad:        req.Cantidad,
		PrecioUnitario:  req.PrecioUnitario,
		Observaciones:   req.Observaciones,
		UsuarioRegistro: req.UsuarioRegistro,
	}
	result, err := h.service.Create(entrada)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Entrada registrada exitosamente",
		Data:    result,
	})
}
