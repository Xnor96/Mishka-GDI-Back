package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/dto"
	"github.com/Mishka-GDI-Back/service"
	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	productoService service.ProductoService
}

func NewProductoHandler(productoService service.ProductoService) *ProductoHandler {
	return &ProductoHandler{
		productoService: productoService,
	}
}

// GetAll godoc
// @Summary Lista todos los productos
// @Description Obtiene una lista de todos los productos disponibles
// @Tags productos
// @Accept json
// @Produce json
// @Success 200 {object} dto.ProductosResponse
// @Failure 500 {object} dto.Response
// @Router /api/productos [get]
func (h *ProductoHandler) GetAll(c *gin.Context) {
	response, err := h.productoService.GetAll()
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
// @Summary Obtiene un producto por ID
// @Description Obtiene los detalles de un producto específico por su ID
// @Tags productos
// @Accept json
// @Produce json
// @Param id path int true "ID del producto"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/productos/{id} [get]
func (h *ProductoHandler) GetByID(c *gin.Context) {
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

	response, err := h.productoService.GetByID(id)
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

// Create godoc
// @Summary Crea un nuevo producto
// @Description Crea un nuevo producto con los datos proporcionados
// @Tags productos
// @Accept json
// @Produce json
// @Param producto body dto.CreateProductoRequest true "Datos del producto"
// @Success 201 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/productos [post]
func (h *ProductoHandler) Create(c *gin.Context) {
	var request dto.CreateProductoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Datos inválidos",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.productoService.Create(&request)
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

// Update godoc
// @Summary Actualiza un producto
// @Description Actualiza los datos de un producto existente
// @Tags productos
// @Accept json
// @Produce json
// @Param id path int true "ID del producto"
// @Param producto body dto.UpdateProductoRequest true "Datos actualizados del producto"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/productos/{id} [put]
func (h *ProductoHandler) Update(c *gin.Context) {
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

	var request dto.UpdateProductoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Datos inválidos",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.productoService.Update(id, &request)
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

// Delete godoc
// @Summary Elimina un producto
// @Description Elimina un producto existente por su ID
// @Tags productos
// @Accept json
// @Produce json
// @Param id path int true "ID del producto"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/productos/{id} [delete]
func (h *ProductoHandler) Delete(c *gin.Context) {
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

	response, err := h.productoService.Delete(id)
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

// GetStockBajo godoc
// @Summary Obtiene productos con stock bajo
// @Description Obtiene una lista de productos con stock por debajo del límite especificado
// @Tags productos
// @Accept json
// @Produce json
// @Param limite query int false "Límite de stock (default: 5)"
// @Success 200 {object} dto.ProductosResponse
// @Failure 500 {object} dto.Response
// @Router /api/productos/stock-bajo [get]
func (h *ProductoHandler) GetStockBajo(c *gin.Context) {
	limiteParam := c.Query("limite")
	limite := 5 // valor por defecto

	if limiteParam != "" {
		if l, err := strconv.Atoi(limiteParam); err == nil && l >= 0 {
			limite = l
		}
	}

	response, err := h.productoService.GetStockBajo(limite)
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

// Search godoc
// @Summary Busca productos
// @Description Busca productos por código o nombre
// @Tags productos
// @Accept json
// @Produce json
// @Param q query string true "Término de búsqueda"
// @Success 200 {object} dto.ProductosResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/productos/buscar [get]
func (h *ProductoHandler) Search(c *gin.Context) {
	termino := c.Query("q")
	if termino == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Parámetro de búsqueda requerido",
			Error:   "El parámetro 'q' es requerido",
		})
		return
	}

	response, err := h.productoService.Search(termino)
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
