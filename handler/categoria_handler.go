package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/dto"
	"github.com/Mishka-GDI-Back/service"
	"github.com/gin-gonic/gin"
)

type CategoriaHandler struct {
	categoriaService service.CategoriaService
}

func NewCategoriaHandler(categoriaService service.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{
		categoriaService: categoriaService,
	}
}

// GetAll godoc
// @Summary Lista todas las categorías
// @Description Obtiene una lista de todas las categorías disponibles
// @Tags categorias
// @Accept json
// @Produce json
// @Success 200 {object} dto.CategoriasResponse
// @Failure 500 {object} dto.Response
// @Router /api/categorias [get]
func (h *CategoriaHandler) GetAll(c *gin.Context) {
	response, err := h.categoriaService.GetAll()
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
// @Summary Obtiene una categoría por ID
// @Description Obtiene los detalles de una categoría específica por su ID
// @Tags categorias
// @Accept json
// @Produce json
// @Param id path int true "ID de la categoría"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/categorias/{id} [get]
func (h *CategoriaHandler) GetByID(c *gin.Context) {
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

	response, err := h.categoriaService.GetByID(id)
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
// @Summary Crea una nueva categoría
// @Description Crea una nueva categoría con los datos proporcionados
// @Tags categorias
// @Accept json
// @Produce json
// @Param categoria body dto.CreateCategoriaRequest true "Datos de la categoría"
// @Success 201 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/categorias [post]
func (h *CategoriaHandler) Create(c *gin.Context) {
	var request dto.CreateCategoriaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Datos inválidos",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.categoriaService.Create(&request)
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
// @Summary Actualiza una categoría
// @Description Actualiza los datos de una categoría existente
// @Tags categorias
// @Accept json
// @Produce json
// @Param id path int true "ID de la categoría"
// @Param categoria body dto.UpdateCategoriaRequest true "Datos actualizados de la categoría"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/categorias/{id} [put]
func (h *CategoriaHandler) Update(c *gin.Context) {
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

	var request dto.UpdateCategoriaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Datos inválidos",
			Error:   err.Error(),
		})
		return
	}

	response, err := h.categoriaService.Update(id, &request)
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
// @Summary Elimina una categoría
// @Description Elimina una categoría existente por su ID
// @Tags categorias
// @Accept json
// @Produce json
// @Param id path int true "ID de la categoría"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/categorias/{id} [delete]
func (h *CategoriaHandler) Delete(c *gin.Context) {
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

	response, err := h.categoriaService.Delete(id)
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
