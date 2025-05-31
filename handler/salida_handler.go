package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/dto"
	"github.com/Mishka-GDI-Back/service"
	"github.com/gin-gonic/gin"
)

type SalidaHandler struct {
	salidaService service.SalidaProductoService
}

func NewSalidaHandler(salidaService service.SalidaProductoService) *SalidaHandler {
	return &SalidaHandler{
		salidaService: salidaService,
	}
}

// GetAll godoc
// @Summary Lista todas las salidas
// @Description Obtiene una lista de todas las salidas de productos
// @Tags salidas
// @Accept json
// @Produce json
// @Success 200 {object} dto.SalidasResponse
// @Failure 500 {object} dto.Response
// @Router /api/salidas [get]
func (h *SalidaHandler) GetAll(c *gin.Context) {
	response, err := h.salidaService.GetAll()
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
// @Summary Obtiene una salida por ID
// @Description Obtiene los detalles de una salida específica por su ID
// @Tags salidas
// @Accept json
// @Produce json
// @Param id path int true "ID de la salida"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/salidas/{id} [get]
func (h *SalidaHandler) GetByID(c *gin.Context) {
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

	response, err := h.salidaService.GetByID(id)
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
// @Summary Obtiene salidas por ID de producto
// @Description Obtiene todas las salidas de un producto específico
// @Tags salidas
// @Accept json
// @Produce json
// @Param id path int true "ID del producto"
// @Success 200 {object} dto.SalidasResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/salidas/producto/{id} [get]
func (h *SalidaHandler) GetByProductoID(c *gin.Context) {
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

	response, err := h.salidaService.GetByProductoID(id)
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
// @Summary Obtiene salidas por fecha
// @Description Obtiene todas las salidas de una fecha específica
// @Tags salidas
// @Accept json
// @Produce json
// @Param fecha path string true "Fecha en formato YYYY-MM-DD"
// @Success 200 {object} dto.SalidasResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/salidas/fecha/{fecha} [get]
func (h *SalidaHandler) GetByFecha(c *gin.Context) {
	fecha := c.Param("fecha")
	if fecha == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Fecha requerida",
			Error:   "La fecha es requerida en formato YYYY-MM-DD",
		})
		return
	}

	response, err := h.salidaService.GetByFecha(fecha)
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
// @Summary Crea una nueva salida
// @Description Registra una nueva salida de producto
// @Tags salidas
// @Accept json
// @Produce json
// @Param salida body dto.CreateSalidaProductoRequest true "Datos de la salida"
// @Success 201 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/salidas [post]
func (h *SalidaHandler) Create(c *gin.Context) {
	var request dto.CreateSalidaProductoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Datos inválidos",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.salidaService.Create(&request)
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
