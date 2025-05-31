package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/dto"
	"github.com/Mishka-GDI-Back/service"
	"github.com/gin-gonic/gin"
)

type EntradaHandler struct {
	entradaService service.EntradaProductoService
}

func NewEntradaHandler(entradaService service.EntradaProductoService) *EntradaHandler {
	return &EntradaHandler{
		entradaService: entradaService,
	}
}

// GetAll godoc
// @Summary Lista todas las entradas
// @Description Obtiene una lista de todas las entradas de productos
// @Tags entradas
// @Accept json
// @Produce json
// @Success 200 {object} dto.EntradasResponse
// @Failure 500 {object} dto.Response
// @Router /api/entradas [get]
func (h *EntradaHandler) GetAll(c *gin.Context) {
	response, err := h.entradaService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error interno del servidor",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetByID godoc
// @Summary Obtiene una entrada por ID
// @Description Obtiene los detalles de una entrada específica por su ID
// @Tags entradas
// @Accept json
// @Produce json
// @Param id path int true "ID de la entrada"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/entradas/{id} [get]
func (h *EntradaHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "ID inválido",
			Error:   "El ID debe ser un número entero",
		})
		return
	}

	response, err := h.entradaService.GetByID(id)
	if err != nil {
		if response != nil && !response.Success {
			c.JSON(http.StatusNotFound, response)
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error interno del servidor",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetByProductoID godoc
// @Summary Obtiene entradas por ID de producto
// @Description Obtiene todas las entradas de un producto específico
// @Tags entradas
// @Accept json
// @Produce json
// @Param id path int true "ID del producto"
// @Success 200 {object} dto.EntradasResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/entradas/producto/{id} [get]
func (h *EntradaHandler) GetByProductoID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "ID inválido",
			Error:   "El ID debe ser un número entero",
		})
		return
	}

	response, err := h.entradaService.GetByProductoID(id)
	if err != nil {
		if response != nil && !response.Success {
			c.JSON(http.StatusNotFound, response)
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error interno del servidor",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetByFecha godoc
// @Summary Obtiene entradas por fecha
// @Description Obtiene todas las entradas de una fecha específica
// @Tags entradas
// @Accept json
// @Produce json
// @Param fecha path string true "Fecha en formato YYYY-MM-DD"
// @Success 200 {object} dto.EntradasResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/entradas/fecha/{fecha} [get]
func (h *EntradaHandler) GetByFecha(c *gin.Context) {
	fecha := c.Param("fecha")
	if fecha == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Fecha requerida",
			Error:   "La fecha es requerida en formato YYYY-MM-DD",
		})
		return
	}

	response, err := h.entradaService.GetByFecha(fecha)
	if err != nil {
		if response != nil && !response.Success {
			c.JSON(http.StatusBadRequest, response)
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error interno del servidor",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Crea una nueva entrada
// @Description Registra una nueva entrada de producto
// @Tags entradas
// @Accept json
// @Produce json
// @Param entrada body dto.CreateEntradaProductoRequest true "Datos de la entrada"
// @Success 201 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/entradas [post]
func (h *EntradaHandler) Create(c *gin.Context) {
	var request dto.CreateEntradaProductoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Datos inválidos",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.entradaService.Create(&request)
	if err != nil {
		if response != nil && !response.Success {
			c.JSON(http.StatusBadRequest, response)
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error interno del servidor",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}
