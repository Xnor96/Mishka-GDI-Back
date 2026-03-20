package handler

import (
	"net/http"
	"strconv"

	"github.com/Mishka-GDI-Back/application"
	"github.com/Mishka-GDI-Back/infrastructure/http/dto"
	"github.com/gin-gonic/gin"
)

type CategoriaHandler struct {
	service application.CategoriaService
}

func NewCategoriaHandler(service application.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{service: service}
}

func (h *CategoriaHandler) GetAll(c *gin.Context) {
	categorias, err := h.service.GetAll()
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.CategoriasResponse{
		Success:    true,
		Message:    "Categorías obtenidas exitosamente",
		Data:       dto.CategoriasToResponse(categorias),
		TotalCount: len(categorias),
	})
}

func (h *CategoriaHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	categoria, err := h.service.GetByID(id)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Categoría encontrada",
		Data:    dto.CategoriaToResponse(categoria),
	})
}

func (h *CategoriaHandler) Create(c *gin.Context) {
	var req dto.CreateCategoriaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	categoria, err := h.service.Create(req.Nombre)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Categoría creada exitosamente",
		Data:    dto.CategoriaToResponse(categoria),
	})
}

func (h *CategoriaHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	var req dto.UpdateCategoriaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Datos inválidos", Error: err.Error()})
		return
	}
	categoria, err := h.service.Update(id, req.Nombre)
	if err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Categoría actualizada exitosamente",
		Data:    dto.CategoriaToResponse(categoria),
	})
}

func (h *CategoriaHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "ID inválido", Error: "El ID debe ser un número entero"})
		return
	}
	if err := h.service.Delete(id); err != nil {
		handleDomainError(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Categoría eliminada exitosamente"})
}
